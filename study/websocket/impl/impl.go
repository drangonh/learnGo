/*
@Time : 2020/10/27 10:01 上午
@Author : mac
@File : impl.go
@Software: GoLand
*/
package impl

import (
	"errors"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Connection struct {
	wsConn *websocket.Conn
	//读取websocket的channel
	inChan chan []byte
	//给websocket写消息的channel
	outChan   chan []byte
	closeChan chan byte
	mutex     sync.Mutex
	//closeChan 状态
	isClosed bool
}

//初始化长连接
func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConn:    wsConn,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte, 1),
	}
	//启动读协程
	go conn.readLoop()
	//启动写协程
	go conn.writeLoop()
	return
}

//读取websocket消息
func (conn *Connection) ReadMessage() (data []byte, err error) {
	str := <-conn.inChan
	fmt.Println(str)
	fmt.Println(string(str))
	select {
	case data = <-conn.inChan:
		fmt.Println(string(<-conn.inChan))
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

//发送消息到websocket
func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

//关闭连接
func (conn *Connection) Close() {
	//线程安全的Close,可重入
	conn.wsConn.Close()

	//只执行一次
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			goto ERR
		}
		//如果数据量过大阻塞在这里,等待inChan有空闲的位置！
		select {
		case conn.inChan <- data:
		case <-conn.closeChan:
			//closeChan关闭的时候
			goto ERR

		}
	}
ERR:
	conn.Close()
}

func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)
	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			goto ERR

		}
		if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}
