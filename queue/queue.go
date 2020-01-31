package queue

type Queue []int

func (q *Queue) Push(n int) {
	*q = append(*q, n)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
