/*
@Time : 2020/10/30 4:17 下午
@Author : mac
@File : data.go
@Software: GoLand
*/
package main

type Data struct {
	Ip       string   `json:"ip"`
	User     string   `json:"user"`
	From     string   `json:"from"`
	Type     string   `json:"type"`
	Content  string   `json:"content"`
	UserList []string `json:"user_list"`
}
