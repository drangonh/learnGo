package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	num := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 > num {
			num = i - start + 1
		}
		lastOccurred[ch] = i
	}

	fmt.Println(num, start)
	return num
}

func main() {
	lengthOfNonRepeatingSubStr("123456")
	lengthOfNonRepeatingSubStr("strfdas")
	lengthOfNonRepeatingSubStr("我我是一只只")
}
