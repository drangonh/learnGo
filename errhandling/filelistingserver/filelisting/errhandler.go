package filelisting

import (
	"log"
	"net/http"
	"os"
)

type AppHandler func(writer http.ResponseWriter, request *http.Request) error

/*函数式编程，接受一个函数并且返回一个函数，这里返回的函数统一对错误进行了处理*/
func RrrWrapper(handler AppHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			log.Printf("handlering err: %s", err.Error())
			code := http.StatusOK
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
