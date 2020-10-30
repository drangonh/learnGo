/*
@Time : 2020/10/30 10:56 上午
@Author : mac
@File : server.go
@Software: GoLand
*/
// 处理函数
package main

import (
	"bufio"
	"fmt"
	"gomodtest/study/tcp/proto"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		msg, err := proto.Decode(reader) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		//recvStr := string(buf[:n])
		recvStr := "我已经收到了您的消息"
		fmt.Println("client端：", msg)
		conn.Write([]byte(recvStr)) // 发送数据
	}
}

func main() {
	// Listen函数创建的服务端
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}
