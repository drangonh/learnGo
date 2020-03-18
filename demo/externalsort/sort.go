package main

import (
	"bufio"
	"fmt"
	"gomodtest/demo/pipeline"
	"os"
)

func main() {
	p := createPipeline("large.in", 800000000, 4)
	writerFile(p)
	printFile()
}

func printFile() {
	file, err := os.Open("large.out")
	if err != nil {
		fmt.Println("打开文件错误")
		panic(err)
	}
	defer file.Close()
	read := pipeline.ReadSourceFile(file)
	fmt.Println("文件存在")

	count := 0
	for v := range read {
		fmt.Println(v)
		count++
		if count >= 50 {
			break
		}
	}
}

func writerFile(ints <-chan int) {
	fileName := "large.out"
	file, err := os.Create(fileName)
	if err != nil {
		//出错了，不知道怎么办了
		panic(err)
	}
	fmt.Println("创建成功")
	//在程序结束之前关闭file
	defer file.Close()

	pipeline.WriterFile(file, ints)
}

func createPipeline(fileName string, fileSize int, chunkCount int) <-chan int {
	size := fileSize / chunkCount

	arr := []<-chan int{}

	pipeline.InitStartTime()
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}

		//file.seek：file从头开始，偏移为size*i；就是跳转到文本的size*i处
		file.Seek(int64(size*i), 0)

		fmt.Println(i, "_______", size)
		//从file中指定为位置开始读取数据，并且读取size的量
		read := pipeline.ReadSource(bufio.NewReader(file), size)

		sort := pipeline.SortArr(read)

		//pipeline.PrintChan(sort)
		//把file中读取出来的分成了chunkCount块，然后放进arr（slice）中
		arr = append(arr, sort)
	}

	//MergeN接受的是chan，所以arr后面加...
	return pipeline.MergeN(arr...)
}
