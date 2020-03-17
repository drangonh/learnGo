package main

import (
	"gomodtest/errhandling/filelistingserver/filelisting"
	"gomodtest/errhandling/filelistingserver/handler"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/", filelisting.ErrWrapper(handler.HandlerFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
