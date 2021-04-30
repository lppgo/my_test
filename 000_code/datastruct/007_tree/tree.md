[toc]

# 1: 树形结构

## 1.1 Tree的相关术语
根节点: 
父节点：
子节点：
节点的度 Degree: 该节点拥有的子树的个数
叶子节点 Leaf: 度为0的节点，也就是最下层没有子节点的节点
非叶子节点: 度不为0的节点,根节点除外
节点的深度: 从上往下定义，根节点到这个节点所经历的边的个数。(根节点深度是0，也可以从1开始)
树的深度: 
节点的高度: 从下往上定义，从该节点到叶子节点的最长边的个数。
树的高度: 


## 1.2 树的存储结构
1: **双亲表示法** 
树的每个节点不一定有孩子，但是一定有且仅有一个双亲节点（根节点除外），所以可以通过其双亲表示一个完整的树。
```go
type Node struct{
    Data interface{} //数据域:存储节点里面的数据
    Parent int //指针域；存储双亲节点在数组的下标
}

type Tree struct{
    Nodes [ ]*Node //节点指针数组
    Root int //根节点位置
    Num int //节点数
}
```
数据存储格式如下:
![双亲表示法---存储树](https://ftp.bmp.ovh/imgs/2021/04/4f8ab30aa94adc3b.png)

2: **孩子表示法**
把每个节点的孩子节点排列起来，以链表存储，如果树有 n 个节点，就会有 n 个链表，如果是叶节点，则此单链表为空链表，然后链表的头指针放在一个数组中。
```go
// 链表中的每个节点存储的不再是数据本身，而是数据在数组中的下标
type Node struct {
    Child   int                // 数据域：数组下标
    Next    *Node              // 指针域：指向该节点的下一个孩子节点的指针
}

// 表头结构
type Head struct {
    Data    interface{}       // 数据域：存储节点数据
    Head   *Node              // 头指针域：存储该节点的孩子链表的头指针
}

// 树结构
type Tree struct {
    Nodes []*Node
    Root int
    Num int
}
```
![](https://ftp.bmp.ovh/imgs/2021/04/4c13b6ee14e0f6b0.png)


3: **孩子兄弟表示法**
任意一棵树，它的节点的第一个孩子如果存在就是唯一的，它的右兄弟如果存在也是唯一的，因此设置两个指针，分别指向该节点的第一个孩子和此节点的右兄弟即可。
```go
type Node struct{
    Data interface{} // 数据域
    LeftChild  *Node // 指向左孩子
    RightChild *Node // 指向右孩子
}
```

# 2：二叉树

1. 二叉树: 树的每个节点的度最大值是2，即最多拥有2棵子树

2. 二叉树的五种形态
    (1). 空二叉树  
    (2). 只有一个根节点
    (3). 根节点只有左子树
    (4). 根节点只有右子树
    (5). 根节点既有左子树又有右子树

3. 特殊的二叉树
    (1). 斜二叉树: 有左斜树(只有左孩子)，右斜树(只有右孩子)
    (2). 真二叉树: 所有节点的度要么为0，要么为2. Proper Binary Tree  
    (3). 满二叉树: 除了叶子节点外，所有节点度都是2.Full Binary Tree
    (4). 完全二叉树: Complete Binary Tree
4. **非空二叉树的特性**
    (1). 每层节点数：二叉树第i层上，最多有 2^(i-1) 个节点
    (2). 全部节点数：高度为h的二叉树，最多有 (2^h)-1 个节点
    (3). 叶子节点与度的关系：叶子节点数=度为2的节点数+1
    (4). 

## 2.1 满二叉树
## 2.1 完全二叉树
完全二叉树（Complete Binary Tree）的特点：
- 度为1的节点只有左子树
- 度为1的节点要么是1个，要么是0个
- 同样节点数的二叉树，完全二叉树深度最小
- 如果节点度为1，则该节点只有孩子节点，不存在只有右子树的情况

假设完全二叉树的高度为 h ，h >= 1，则：
- 则至少有 2h-1 个节点（2$^0$ + 2$^1$ + ... + 2$^h$$^-$$^2$ + 1）
- 则最多有 2h - 1 个节点 (2$^0$ + 2$^1$ + ... + 2$^h$$^-$$^1$，满二叉树)

一颗有n(n>0)个节点的完全二叉树，从上到下，从左到右，给节点从0编号，则第i个节点有：
- i=0，是根节点
- i>2，其父节点编号是 floor((i-1)/2)
- 其左子节点编号为2i+1 
- 其右子节点编号为2i+2
  
## 2.3 二叉树的 Traversal

### 2.3.1 层序遍历 Level Order Traversal
> (广度优先搜索)
```go
/**
实现思路：无法用递归实现
1 将各节点入队
2 循环执行以下操作，直到队列为空
	取出队头节点出队，进行访问
	将队头节点的左子节点入队
	将队头节点的右子节点入队
*/
func LevelOrderTraverse(node *Node) {
	if node == nil {
		return
	}

	queue := list.New()		// 制作一个队列
	queue.PushBack(node)

	for queue.Len() != 0 {
		queueHead := queue.Remove(queue.Front())	// 队首出队
		tempNode := queueHead.(*Node)				// 类型断言
		fmt.Printf("%v ", tempNode.data)
		if tempNode.left != nil {
			queue.PushBack(tempNode.left)
		}
		if tempNode.right != nil {
			queue.PushBack(tempNode.right)
		}
	}
}
```
### 2.3.2 前序遍历 PerOrder Traversal
> (深度优先搜索)
```go
func PreOrderTraverse(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%v ", node.data)			// 前序遍历就是从node开始遍历，所以要先打印
	PreOrderTraverse(node.left)
	PreOrderTraverse(node.right)
}
```
### 2.3.3 中序遍历 InOrder Traversal
> (深度优先搜索)
```go
func InOrderTraverse(node *Node) {
	if node == nil {
		return
	}
	// 会产生式升序结果
	InOrderTraverse(node.left)
	fmt.Printf("%v ", node.data)
	InOrderTraverse(node.right)

	// 会产生降序结果
	//InOrderTraverse(node.right)
	//fmt.Printf("%v ", node.data)
	//InOrderTraverse(node.left)
}
```
### 2.3.4 后序遍历 PostOrder Traversal
> (深度优先搜索)
```go
func PostOrderTraverse(node *Node) {
	if node == nil {
		return
	}
	PostOrderTraverse(node.left)
	PostOrderTraverse(node.right)
	fmt.Printf("%v ", node.data)
}
```

## 2.4 二叉搜索树 BST
### 2.4.1 BST的特点
二叉搜索树可以为空。如果不为空，则满足：
- 非空左子树的所有键值小于其根节点的键值
- 非空右子树的所有键值大于其根节点的键值
- 左、右子树本身也都是二叉搜索树

二叉搜索树的概念是一致的，利用二分查找的思想

### 2.4.2 二叉搜索树与哈希表的对比
1. Hash表需要一个很大的数组，会造成一定的空间浪费
2. Hash表的数据是无序的，二叉搜索树利用的是有序数组二分查找的思想
3. 二叉搜索树的时间复杂度是O(logn)，但是当BST不是平衡二叉树的时候，最坏时间复杂度是O(n)

### 2.4.3 AVL 平衡二叉树
### 2.4.4 Red Black Tree 红黑树

### 2.4.5 B树 (平衡二叉搜索树)
### 2.4.6 B+树





# 3：二叉堆