package tree

/*
go语言的封装：
驼峰命名
首字母大写：public
首字母小写：private
每个目录一个包：package，但是包名可以不用和目录名字一样，main包包含可执行入口
为结构定义的方法必须放在同一个包内，可以是不同的文件

*/
import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

/*
为结构定义方法，需要一个接受者,(node treeNode)接受是谁调用print()方法。其中node就是接受者
*/
func (node *Node) Printf() {
	fmt.Println(node.Value)
}

/*
go中使用值传递，所以需要修改node的value需要传入node的地址进来；值传递需要拷贝一份给接受者。地址传递不需要拷贝
*/
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("如果node为nil也可以调用方法，但是node不存在的话无法取value属性")
		return
	}
	node.Value = value
}

func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		n.Printf()
	})
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

//使用工厂来构造我们想要的实例，返回的&treeNode{value: value}的地址是局部函数内部的变量地址。
//在go中局部变量也可以返回地址给别人使用
//工厂函数返回一个局部变量的地址，我们不需要管它在哪里分配。
/*
如果局部变量不需要返回给别人使用那就是在栈上分配的；否在在堆上分配，然后这个变量就会参与垃圾回收
等带别人使用完这个局部变量之后就会回收掉
*/
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(nodeNew *Node) {
			out <- nodeNew
		})
		//传出数据之后关闭chan
		close(out)
	}()
	return out
}
