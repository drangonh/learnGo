package main

import (
	"fmt"
	"regexp"
	"strings"
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
	// ([a-zA-Z0-9]+)加上括号之后就可以把相关的部分提取给我们，配合方法FindAllStringSubmatch可以获取一个二维数组
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	//res := re.FindAllString(text, -1)
	res := re.FindAllStringSubmatch(text, -1)
	fmt.Println(res)

	s := `### 编辑文章内容

	使用markdown的语法
	`
	replace(s)

}

func replace(s string) {
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, " ", "", -1)

	fmt.Println("替换后的s：：", s)

	str := "这里是 www\n.runoob\n.com"
	fmt.Println("-------- 原字符串 ----------")
	fmt.Println(str)
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	fmt.Println("-------- 去除空格与换行后 ----------")
	fmt.Println(str)
}
