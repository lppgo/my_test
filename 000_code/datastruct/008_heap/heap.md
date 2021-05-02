- [Headp 堆](#headp-堆)
  - [1: Complete Binary Tree](#1-complete-binary-tree)
  - [2: Binary Heap 二叉堆](#2-binary-heap-二叉堆)
    - [2.1：二叉堆的定义](#21二叉堆的定义)
    - [2.2：二叉堆的特性](#22二叉堆的特性)
    - [2.3：Go使用Heap堆](#23go使用heap堆)

# Headp 堆

## 1: Complete Binary Tree

## 2: Binary Heap 二叉堆

### 2.1：二叉堆的定义

### 2.2：二叉堆的特性

### 2.3：Go使用Heap堆

**Heap堆最重要的性质:**
- 如果任意节点的值总是>=子节点的值，称为最大堆，大根堆;
- 如果任意节点的值总是<=子节点的值，称为最小堆，小根堆.

**堆的时间复杂度:**
- 获取最大/小值: O(1);
- 删除最大/小值: O(logn);
- 添加元素:O(logn);




实现 container/heap 里面的接口

Len()
Less()
Swap()
Push()
Pop()
