package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node *Node) Printf() {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("如果node为nil也可以调用方法，但是node不存在的话无法取value属性")
		return
	}
	node.Value = value
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Printf()
	node.Right.Traverse()
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}
