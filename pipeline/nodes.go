package pipeline

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

var startTime time.Time

func InitStartTime() {
	startTime = time.Now()
}

func SourceArr(a ...int) chan int {
	out := make(chan int)
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func SortArr(in <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		arr := []int{}
		for num := range in {
			arr = append(arr, num)
		}
		fmt.Println("read done:", time.Now().Sub(startTime))
		sort.Ints(arr)
		fmt.Println("sort done:", time.Now().Sub(startTime))
		for _, v := range arr {
			out <- v
		}
		close(out)
	}()

	return out
}

func MergeArr(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024)

	//chan的操作需要在goruntine中处理
	go func() {
		a1, ok1 := <-in1
		a2, ok2 := <-in2

		for ok1 || ok2 {
			//这个if只管满足a1满足条件即可（a1单独存在，或者都存在时候a1小于a2）
			if !ok2 || (ok1 && ok2 && a1 < a2) {
				out <- a1
				//更新a1值
				a1, ok1 = <-in1
			} else {
				out <- a2
				a2, ok2 = <-in2
			}
		}
		fmt.Println("merge done:", time.Now().Sub(startTime))
		close(out)
	}()

	return out
}

//读文件，返回一个channel
func ReadSourceFile(reader io.Reader) <-chan int {
	out := make(chan int)
	go func() {
		buffer := make([]byte, 8)

		for {
			n, err := reader.Read(buffer)
			if n > 0 {
				v := binary.BigEndian.Uint64(buffer)
				out <- int(v)
			}

			if err != nil {
				break
			}
		}

		close(out)
	}()
	return out
}

//写文件,把通道中的数据写进文件中
func WriterFile(writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

//生成随机数
func RandomNum(num int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < num; i++ {
			out <- rand.Int()
			//out <- i
		}

		close(out)
	}()
	return out
}

//归并多个,利用递归进行归并
func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}
	m := len(inputs) / 2
	//fmt.Println("slice：：", m)
	//这里用...是因为inputs[:m]是slice，加上...转成channel
	merge := MergeArr(
		MergeN(inputs[:m]...), MergeN(inputs[m:]...))
	return merge
}

//分块读文件，返回一个channel;chunkSize是-1的话就全部读，不是的话就分块读
func ReadSource(reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			n, err := reader.Read(buffer)
			bytesRead += n
			//fmt.Println("读了多少", chunkSize, bytesRead)
			if n > 0 {
				v := binary.BigEndian.Uint64(buffer)
				out <- int(v)
			}

			if err != nil || (chunkSize != -1 && chunkSize <= bytesRead) {
				break
			}
		}

		close(out)
	}()
	return out
}

func PrintChan(in <-chan int) {

	fmt.Println("输出channel")
	for v := range in {
		fmt.Println(v)
	}
}
