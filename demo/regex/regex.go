package main

import (
	"fmt"
	"regexp"
)

const text = `
my name is ccmouse@gmain.com
my name is 89dssad@qq.com
my name is 89dssad@qq.com.cn
`

//正则表达式学习
func main() {
	//re := regexp.MustCompile("ccmouse@gmain.com")

	// .匹配任意字符，+至少一个字符，*0个或多个字符
	// ""中的字符会转义，\.表示转义字符. \\.表示字符.   ,在``中的字符不会转义
	//re := regexp.MustCompile(".+@.+\\..+")

	const text = `
my name is ccmouse@gmain.com
my name is 89dssad@qq.com
my name is 89dssad@qq.com.cn
`
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)
	res := re.FindAllString(text, -1)
	fmt.Println(res)
}
