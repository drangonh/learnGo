package main

import "fmt"

var lastOccurred = make([]int, 0xffff) //初始了65535字节的大小

func lengthOfNonRepeatingSubStr(s string) int {

	//这里用slice来优化map，利用空间换时间，达到优化运行速度的效果
	//lastOccurred := make(map[rune]int)
	//lastOccurred := make([]int, 0xffff)

	for i := range lastOccurred {
		lastOccurred[i] = -1
	}

	start := 0
	num := 0

	for i, ch := range []rune(s) {
		fmt.Println(i, ch)
		if lastI := lastOccurred[ch]; lastOccurred[ch] != -1 && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 > num {
			num = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return num
}

func main() {
	lengthOfNonRepeatingSubStr("123456")
	lengthOfNonRepeatingSubStr("strfdas")
	lengthOfNonRepeatingSubStr("我我是一只只")
}
