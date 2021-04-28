package queue

import (
	"fmt"
)

// 链表实现队列
type LinkedQueue struct {
	length int
	head   *LinkedQueueNode
	tail   *LinkedQueueNode
}

type LinkedQueueNode struct {
	Data interface{}
	Next *LinkedQueueNode
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

// 入队
func (q *LinkedQueue) EnterQueue(e interface{}) {
	enNode := &LinkedQueueNode{
		Data: e,
		Next: nil,
	}

	if q.length == 0 { //当queue为空的时候
		q.head = enNode
		q.tail = enNode
		q.length++
		return
	}

	// 添加新节点 1：直接添加到tail节点，2：或者遍历到最后一个尾节点再添加
	currentNode := q.head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = enNode
	q.tail = enNode
	q.length++
}

// 出队
func (q *LinkedQueue) OutQueue() interface{} {
	if q.length == 0 {
		fmt.Println("队列为空")
		return nil
	}

	tmpNode := q.head
	q.head = tmpNode.Next
	// 若队头是队尾，则出队后将tail指向head
	if q.tail == tmpNode {
		q.tail = q.head
	}
	q.length--
	return tmpNode.Data
}

func (q *LinkedQueue) ength() int {
	return q.length
}
