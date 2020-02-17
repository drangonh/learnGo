package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

/*把文件内容写进writer，并且返回错误给errhandler去错误*/
const prefix = "/list/"

func HandlerFileList(writer http.ResponseWriter, request *http.Request) error {
	if index := strings.Index(request.URL.Path, prefix); index != 0 {
		fmt.Println(index)
		return userError("path must start with:" + prefix)
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}

type userError string

func (err userError) Error() string {
	return "Error:" + err.Message()
}

func (err userError) Message() string {
	return string(err)
}
