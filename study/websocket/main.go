/*
@Time : 2020/10/27 10:01 上午
@Author : mac
@File : main.go
@Software: GoLand
*/
package main

import (
	"fmt"
	"gomodtest/study/websocket/impl"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var (
	upgrade = websocket.Upgrader{
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		//websocket 长连接
		wsConn *websocket.Conn
		err    error
		conn   *impl.Connection
		data   []byte
	)
	//header中添加Upgrade:websocket
	if wsConn, err = upgrade.Upgrade(w, r, nil); err != nil {
		return
	}

	if conn, err = impl.InitConnection(wsConn); err != nil {
		goto ERR
	}

	go func() {
		var (
			err error
		)
		for {
			if err = conn.WriteMessage([]byte("heartbeat")); err != nil {
				return
			}
			time.Sleep(time.Second * 1)
		}
	}()

	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()
}

type Says interface {
	say()
}

type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println(c.name)
}

func main() {

	// 类型断言配合interface使用。因为interface空接口可以代表任何类型

	// recover必须在Panic之前，而且recover必须和defer一起使用
	// 多个Panic，只能捕获最后一个Panic。
	// defer不能直接跳用recover()
	defer func() {
		recover()
	}()
	panic("panic error!")

	////http标准库
	//http.HandleFunc("/ws", wsHandler)
	//http.ListenAndServe("0.0.0.0:5555", nil)
}

// 返回2个函数类型的返回值
func test01(base int) (func(int) int, func(int) int) {
	// 定义2个函数，并返回
	// 相加
	add := func(i int) int {
		base += i
		return base
	}
	// 相减
	sub := func(i int) int {
		base -= i
		return base
	}
	// 返回
	return add, sub
}
