/*
@Time : 2020/10/21 2:47 下午
@Author : mac
@File : mapstudy.go
@Software: GoLand
*/
package _map

import "fmt"

func MapDemo1() {
	b := make(map[string]int)
	fmt.Println(b)
	b["fda"] = 12
	fmt.Println(&b)

	ma := new(map[string]int)
	fmt.Println(ma)
	//第一种初始化方法
	*ma = map[string]int{}
	(*ma)["a"] = 44
	fmt.Println(*ma)
}
