package queue

import "fmt"

/*
示例代码：
1：文件名:queue_test.go
2:方法名Example开头
3：fmt.Println输出结果，并且用Output:和期望的结果值来判断是否正确
*/
func ExampleQueue_IsEmpty() {
	q := Queue{1}

	fmt.Println(q.IsEmpty())
	//Output:
	//false
}

func ExampleQueue_Pop() {
	q := Queue{1, 3, 10}
	q.Pop()

	fmt.Println(q)
	//Output:
	//[3 10]
}
