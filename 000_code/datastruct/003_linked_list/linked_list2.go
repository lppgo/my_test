/**
 * @Author: lucas
 * @Description:
 * @File:  linked_list2
 * @Version: 1.0.0
 * @Date: 2021/3/18 17:41
 */
package linked_list

import (
	"fmt"
)

func NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

func NewNodeWithNext(v int, next *Node) *Node {
	return &Node{
		Value: v,
		Next:  next,
	}
}

func NewLinkedList() *LinkedList {
	head := &Node{
		Value: 0,
		Next:  nil,
	}
	return &LinkedList{Head: head}

}

//
func (l *LinkedList) Append(n *Node) {
	tmp := l.Head
	for {
		if tmp.Next != nil {
			tmp = tmp.Next
		} else {
			break
		}
	}
	tmp.Next = n
}

// 插入链表表头
func (l *LinkedList) InsertFront(n *Node) {
	n.Next = l.Head.Next
	l.Head.Next = n
}

// 插入链表中间
func (l *LinkedList) InsertAfter(n *Node) {

}

// 插入链表尾部
func (l *LinkedList) InsertEnd(n *Node) {

}

// 链表是否为空
func (l *LinkedList) IsEmpty() bool {
	return l.Head.Next == nil
}

// 遍历链表
func (l *LinkedList) Range() {
	if l.IsEmpty() {
		fmt.Println("Linked List is empty !")
		return
	}
	tmp := l.Head.Next
	i := 0
	for {
		i++
		if tmp.Next != nil {
			fmt.Printf("i :%d , data:%d \n", i, tmp.Value)
			tmp = tmp.Next
		} else {
			break
		}
	}
	fmt.Printf("i :%d , data:%d \n", i, tmp.Value)
}

// 寻找链表中是否有对应的值，并返回对应值的node
func (l *LinkedList) Find(value int) *Node {
	if l.IsEmpty() {
		fmt.Println("Linked List is empty !")
		return nil
	}

	tmp := l.Head.Next
	for {
		if tmp != nil {
			if tmp.Value == value {
				return tmp
			}
			tmp = tmp.Next
		} else {
			break
		}
	}
	return nil
}

// 单链表反转(就地逆序)
func (l *LinkedList) Reverse_Locally() {
	var pre *Node       // 前驱节点
	var cur *Node       // 当前节点
	next := l.Head.Next // 后继节点
	for next.Next != nil {
		cur = next.Next // 将后继节点原来的下一个节点保存到当前节点
		next.Next = pre // 将后继节点的Next节点指向前驱节点
		pre = next      //
		next = cur      //

	}
	l.Head.Next = pre
}

//
func (l *LinkedList) Reverse_Recursion() {

}

//
func (l *LinkedList) Reverse_Inserted() {
	var cur *Node  // 当前节点
	var next *Node // 后继节点

	cur = l.Head.Next.Next
	// 设置第一个节点的next为nil
	l.Head.Next.Next = nil
	for cur != nil {
		next = cur.Next
		cur.Next = l.Head.Next
		l.Head.Next = cur
		cur = next
	}
}
