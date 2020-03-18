package filelisting

import (
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

/*函数式编程，接受一个函数并且返回一个函数，这里返回的函数统一对错误进行了处理*/
func ErrWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic :%v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			log.Printf("handlering err: %s", err.Error())
			if userError, ok := err.(userError); ok {
				//userError显示给用户
				http.Error(writer, userError.Error(), http.StatusBadRequest)
				return
			}
			code := http.StatusOK

			//systemError不显示给用户
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)

		}
	}
}

type userError interface {
	Error() string
	Message() string
}
