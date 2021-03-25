package linked_list

// 链表 .
type Node struct {
	Value int
	Next  *Node //后一个node地址
	// Per *Node //前一个node地址
}

type LinkedList struct {
	Head *Node
}

// 各种应用场景
func Application() {
	l := NewLinkedList()
	node1 := NewNode(10)
	node2 := NewNode(20)
	node3 := NewNode(30)
	node4 := NewNode(40)
	node5 := NewNode(50)

	l.Append(node1)
	l.Append(node2)
	l.Append(node3)
	l.Append(node4)
	l.Append(node5)

	node0 := NewNode(5)
	l.InsertFront(node0)

	l.Range()

	// 查询链表是否有对应的value，并返回对应node
	// fmt.Println(l.Find(30))

	// 	TODO 1: 单链表反转,就当逆序
	l.Reverse_Locally()
	l.Range()
	// 	TODO 2: 单链表反转,递归
	l.Reverse_Recursion()
	l.Range()
	// 	TODO 3: 单链表反转,插入法
	l.Reverse_Inserted()
	l.Range()
}
