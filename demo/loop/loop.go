package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scaner := bufio.NewScanner(reader)

	for scaner.Scan() {
		fmt.Println(scaner.Text())
	}
}

func main() {
	s := `abs
fdafdas
czcz
`
	printFileContents(strings.NewReader(s))

	printFile("demo/loop/abc.txt")
}
