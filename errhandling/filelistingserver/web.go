package main

import (
	"gomodtest/errhandling/filelistingserver/filelisting"
	"net/http"
)

func main() {
	http.HandleFunc("/list/", filelisting.RrrWrapper(filelisting.HandlerFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
