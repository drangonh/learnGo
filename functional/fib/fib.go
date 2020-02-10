package fib

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*斐波拉契数列*/
func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//函数类型也可以实现接口
type intGen func() int

//实现了Read接口
func (i intGen) Read(p []byte) (n int, err error) {
	fmt.Println("进入intGen的Read")
	next := i()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	fmt.Printf("%T %v\n", scanner, scanner)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()
	printFileContents(f)
}
