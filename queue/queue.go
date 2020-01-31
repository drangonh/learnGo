package queue

//interface{}代表任意类型，表示这个slice接受任何类型的值
type Queue []interface{}

func (q *Queue) Push(n interface{}) {
	//限定Push的值必须为int
	*q = append(*q, n.(int))
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]

	//interface{}转int
	return head.(int)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
