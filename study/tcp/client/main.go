/*
@Time : 2020/10/30 10:54 上午
@Author : mac
@File : main.go
@Software: GoLand
*/
package main

import (
	"bufio"
	"fmt"
	"gomodtest/study/tcp/proto"
	"net"
	"os"
	"strings"
)

func main() {
	// Dial函数和服务端建立连接：
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	defer conn.Close() // 关闭连接

	// Stdin是指向标准输入描述符。
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("发送")
		input, _ := inputReader.ReadString('\n') // 读取数据直到用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}

		data, err := proto.Encode(inputInfo)
		_, err = conn.Write(data) // 发送数据
		if err != nil {
			fmt.Println("暂无数据发送")
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println("服务端：", string(buf[:n]))
	}
}
