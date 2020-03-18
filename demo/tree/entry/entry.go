package main

import (
	"fmt"
	"go.uber.org/zap"
	"gomodtest/demo/tree"
)

//定义别名扩展已有类
type myTreeTree struct {
	node *tree.Node
}

func (tree *myTreeTree) postOver() {
	if tree == nil || tree.node == nil {
		return
	}

	right := myTreeTree{tree.node.Right}
	left := myTreeTree{tree.node.Left}

	left.postOver()
	right.postOver()
	tree.node.Printf()
}

//使用内嵌方式扩展,可以继承类的方法和属性
type myTree struct {
	*tree.Node
}

func (tree *myTree) post() {
	if tree == nil || tree.Node == nil {
		return
	}

	left := myTree{tree.Left}
	right := myTree{tree.Right}

	left.post()
	right.post()
	tree.Printf()
}

//重载
func (tree *myTree) Traverse() {
	fmt.Println("this is shadowed.")
}

func main() {
	root := myTree{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Traverse()
	root.Node.Traverse()
	fmt.Println()
	//my := myTreeTree{&root}
	//my.postOver()
	//fmt.Println()

	//newMy := myTree{&root}
	//root.post()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})

	fmt.Println("node count:", nodeCount)

	maxCount := 0
	c := root.TraverseWithChannel()
	for node := range c {
		if node.Value > maxCount {
			maxCount = node.Value
		}
	}
	fmt.Println("maxCount:", maxCount)
}

func modTest() {
	logger, _ := zap.NewProduction()
	logger.Warn("warn test")
}
