package queue

import "fmt"

const CircleQueueCap = 5

// 实现循环队列
type CircleQueue struct {
	items  []interface{} // 队列中的数据存放数组（切片）
	cap    int           // 队列的容量
	length int           // 队列元素个数
	head   int           // 队首下标
	tail   int           // 队尾下标
}

func NewCircleQueue() *CircleQueue {
	return &CircleQueue{
		items:  make([]interface{}, CircleQueueCap),
		cap:    CircleQueueCap,
		length: 0,
		head:   0,
		tail:   0,
	}
}

// 入队
func (q *CircleQueue) EnterQueue(e interface{}) bool {
	if q.length == q.cap {
		fmt.Println("circleQueue 已经满了")
		return false
	}

	q.tail = q.RealIndex(q.length) //计算入队后tail的真实位置
	q.items[q.tail] = e
	q.length++
	return true
}

// 出队
func (q *CircleQueue) OutQueue() interface{} {
	if q.length == 0 {
		fmt.Println("circle queue 为空")
		return nil
	}

	// 获取队首元素
	head := q.items[q.head]
	q.head = q.RealIndex(1)
	q.length--
	return head
}

// 获取循环队列元素个数
func (q *CircleQueue) Length() int {
	return q.length
}

// 获取循环队列容量
func (q *CircleQueue) Cap() int {
	return q.cap
}

// 获取循环队列数组中的真实索引
func (q *CircleQueue) RealIndex(index int) int {
	return (q.head + index) % q.Cap()
}
