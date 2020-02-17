package main

import (
	"bufio"
	"fmt"
	"gomodtest/functional/fib"
	os2 "os"
)

func writeFile(fileName string) {
	os, err := os2.Create(fileName)
	if err != nil {
		fmt.Println("error", err.Error())
		if pathError, ok := err.(*os2.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer os.Close()
	f := fib.Fibonacci()

	write := bufio.NewWriter(os)
	defer write.Flush()
	for i := 0; i < 20; i++ {
		n, _ := fmt.Fprintln(write, f())
		fmt.Println(n)
	}
}

func main() {
	writeFile("fib.txt")
}
