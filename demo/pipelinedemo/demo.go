package main

import (
	"bufio"
	"fmt"
	"gomodtest/demo/pipeline"
	"os"
)

//defer是先进后出
func main() {
	fileName := "large.in"
	file, err := os.Create(fileName)
	if err != nil {
		//出错了，不知道怎么办了
		panic(err)
	}
	fmt.Println("创建成功")
	//在程序结束之前关闭file
	defer file.Close()
	p := pipeline.RandomNum(100000000)
	//目前有个bug在bufio.NewWriter(file)时如果生成的数据不够大，会直接结束程序
	writer := bufio.NewWriter(file)
	//writer := file
	pipeline.WriterFile(writer, p)

	file, err = os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p = pipeline.ReadSource(bufio.NewReader(file), -1)

	num := 0
	for v := range p {
		fmt.Println(v)
		num++
		if num > 49 {
			break
		}
	}
}

func merge() {
	p := pipeline.SourceArr(2, 5, 1, 3, 10, 6)
	arr := pipeline.SortArr(p)

	mergeArr := pipeline.MergeArr(arr, pipeline.SortArr(pipeline.SourceArr(13, 3, 4, 9, 5, 6)))
	for v := range mergeArr {
		fmt.Println(v)
	}

	//for {
	//	if num, ok := <-p; ok {
	//		fmt.Println(num)
	//	} else {
	//		break
	//	}
	//}
}
