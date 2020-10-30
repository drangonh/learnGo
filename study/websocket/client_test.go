/*
@Time : 2020/10/29 3:02 下午
@Author : mac
@File : client_test.go
@Software: GoLand
*/
package main

import (
	"net/http"
	"testing"
)

func TestWsHandler(t *testing.T) {
	//http标准库
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:5555", nil)
}
