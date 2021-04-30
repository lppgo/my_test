package tree

import (
	"container/list"
	"fmt"
	"sync"
)

func ExampleTree() {

	tree := NewTree()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)

	if tree.Length() > 0 {
		tree.PreOrderTraversal(tree.Nodes[0])
		fmt.Println()
		tree.InOrderTraversal(tree.Nodes[0])
		fmt.Println()
		tree.PostOrderTraversal(tree.Nodes[0])
		fmt.Println()
		tree.LevelOrderTraversal(tree.Nodes[0])
	}

}

type Node struct {
	Data       interface{}
	LeftChild  *Node
	RightChild *Node
}

type Tree struct {
	Nodes []*Node
	sync.RWMutex
}

const DefaultCap = 1024

func NewTree() *Tree {
	return &Tree{
		Nodes: make([]*Node, 0, DefaultCap),
	}
}

func (t *Tree) Insert(val interface{}) bool {
	newNode := &Node{
		Data:       val,
		LeftChild:  nil,
		RightChild: nil,
	}
	t.RLock()

	ok := false
	length := t.Length()
	switch {
	case length == 0:
		t.Nodes = append(t.Nodes, newNode)
		ok = true
	case t.Nodes[length-1] == nil:
		fmt.Println("树最后一个node是空，添加失败")
		ok = false
	case length%2 == 1: //
		t.Nodes = append(t.Nodes, newNode)
		parentIndex := (t.Length() - 1) / 2
		if t.Nodes[parentIndex].LeftChild != nil {
			fmt.Printf("父节点:%d的左孩子节点:%v不为空，添加失败\n", parentIndex, *t.Nodes[parentIndex].LeftChild)
			ok = false
		} else {
			t.Nodes[parentIndex].LeftChild = newNode
			ok = true
		}
	case length%2 == 0: //
		t.Nodes = append(t.Nodes, newNode)
		parentIndex := (t.Length() - 2) / 2
		if t.Nodes[parentIndex].RightChild != nil {
			fmt.Printf("父节点:%d的右孩子节点:%v不为空，添加失败\n", parentIndex, *t.Nodes[parentIndex].RightChild)
			ok = false
		} else {
			t.Nodes[parentIndex].RightChild = newNode
			ok = true
		}
	}
	t.RUnlock()
	return ok
}
func (t *Tree) Remove(val interface{}) bool {
	return true
}
func (t *Tree) Search(val interface{}) (*Node, bool) {
	// 先遍历，再比较返回
	return nil, true
}

func (t *Tree) Length() int {
	return len(t.Nodes)
}

// 层序遍历
// 思路:
// 1：创建一个queue
// 2：将root节点入队
// 3：队列不为空，取出队首节点
// 4：左孩子不为空，leftChild入队；右孩子不为空，rightChild入队
// 5：反复执行3，4步骤，直到队列为空
func (t *Tree) LevelOrderTraversal(root *Node) {
	length := t.Length()
	if length == 0 {
		fmt.Println("树为空,节点数len==0")
		return
	}
	//
	if root != nil {
		queue := list.New()  // 创建一个queue
		queue.PushBack(root) // root节点入队
		for queue.Len() > 0 {
			head := queue.Remove(queue.Front()) //取出队首元素
			tmpNode := head.(*Node)
			fmt.Printf("层序遍历 Level Order Traversal:%v\n", tmpNode.Data)
			if tmpNode.LeftChild != nil {
				queue.PushBack(tmpNode.LeftChild)
			}
			if tmpNode.RightChild != nil {
				queue.PushBack(tmpNode.RightChild)
			}
		}
	}
}
func (t *Tree) PreOrderTraversal(root *Node) {
	length := t.Length()
	if length == 0 {
		fmt.Println("树为空,节点数len==0")
		return
	}
	// 遍历
	if root != nil {
		fmt.Printf("前序遍历 perOrder traversal node.Data:%v\n", root.Data)
		t.PreOrderTraversal(root.LeftChild)
		t.PreOrderTraversal(root.RightChild)
	}
}
func (t *Tree) InOrderTraversal(root *Node) {
	length := t.Length()
	if length == 0 {
		fmt.Println("树为空,节点数len==0")
		return
	}
	if root != nil {
		t.InOrderTraversal(root.LeftChild)
		fmt.Printf("中序遍历 InOrder traversal node.Data:%v\n", root.Data)
		t.InOrderTraversal(root.RightChild)
	}
}
func (t *Tree) PostOrderTraversal(root *Node) {
	length := t.Length()
	if length == 0 {
		fmt.Println("树为空,节点数len==0")
		return
	}
	if root != nil {
		t.PostOrderTraversal(root.LeftChild)
		t.PostOrderTraversal(root.RightChild)
		fmt.Printf("后序遍历 InOrder traversal node.Data:%v\n", root.Data)
	}
}
