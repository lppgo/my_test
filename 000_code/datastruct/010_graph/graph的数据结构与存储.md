- [1: Graph图的数据结构与存储](#1-graph图的数据结构与存储)
- [2: 常用图的存储方式](#2-常用图的存储方式)
  - [2.1 邻接矩阵 Adjacency Matrix](#21-邻接矩阵-adjacency-matrix)
    - [2.1.1 Undirected Graph Adjacency (无向图邻接矩阵)](#211-undirected-graph-adjacency-无向图邻接矩阵)
    - [2.1.2 Directed Graph Adjacency (有向图邻接矩阵)](#212-directed-graph-adjacency-有向图邻接矩阵)
  - [2.2 邻接表](#22-邻接表)
  - [2.3 十字链表](#23-十字链表)
  - [2.4 邻接多重表](#24-邻接多重表)
  - [2.5 边集数组](#25-边集数组)
矩阵)
    * [2\.1\.2 Directed Graph Adjacency (有向图邻接矩阵)](#212-directed-graph-adjacency-有向图邻接矩阵)
  * [2\.2 邻接表](#22-邻接表)
  * [2\.3 十字链表](#23-十字链表)
  * [2\.4 邻接多重表](#24-邻接多重表)
  * [2\.5 边集数组](#25-边集数组)

# 1: Graph图的数据结构与存储

```go
ADT 图(Graph)
Data
    顶点的有穷非空集合和边的集合。
Operation
    NewGraph(*G, V, VR):    按照顶点集V和边弧集VR的定义构造图G。
    DestroyGraph(*G):       图G存在则销毁。
    LocateVex(G, u):        若图G中存在顶点u， 则返回图中的位置。
    GetVex(G, v):           返回图G中顶点v的值。
    PutVex(G, v, value):    将图G中顶点v赋值value。
    FirstAdjVex(G, *v):     返回顶点v的一个邻接顶点， 若顶点在G中无邻
    NextAdjVex(G, v, *w):   返回顶点v相对于顶点w的下一个邻接顶点，若w是v的最后一个邻接点则返回“空”。
    InsertVex(*G, v):       在图G中增添新顶点v。
    DeleteVex(*G, v):       删除图G中顶点v及其相关的弧。
    InsertArc(*G, v, w):    在图G中增添弧<v,w>， 若G是无向图， 还需要增
    DeleteArc(*G, v, w):    在图G中删除弧<v,w>， 若G是无向图， 则还删除
    DFSTraverse(G):         对图G中进行深度优先遍历， 在遍历过程对每个
    HFSTraverse(G):         对图G中进行广度优先遍历， 在遍历过程对每个
endADT
```


由上可知，传统的存储结构在实现图时都有弊端：
- 顺序存储：基本上很难实现图的结构
- 多重链表：一个数据域和多个指针域组成的结点表示图中的一个顶点， 此时图的各个顶点度数相差很大，按度数最大的顶点设计结点结构会造成很多存储单元的浪费， 而若按每个顶点自己的度数设计不同的顶点结构， 又带来操作的不便。


# 2: 常用图的存储方式
## 2.1 邻接矩阵 Adjacency Matrix

[邻接矩阵](https://github.com/lppgo/over-algorithm/blob/master/07-%E5%9B%BE/04-1-%E9%82%BB%E6%8E%A5%E7%9F%A9%E9%98%B5.md)

>图由顶点和边/弧 2部分组成。顶点不分主次和大小，边/弧需要维护顶点之间的关系，所以可以将顶点和边分开存储。

**邻接矩阵Adjacency Matrix**：一个一维数组存储顶点信息，一个二维数组存储顶点之间边的信息

### 2.1.1 Undirected Graph Adjacency (无向图邻接矩阵)
如下列的无向图，使用邻接矩阵方式表示后的效果：
![](https://i.bmp.ovh/imgs/2021/04/64cf66025785074b.png)

从上图可以看出，undirected graph adjacency matrix的边的二维数组是一个对称的matrix矩阵.
**无向图邻接矩阵存储可以得出以下结论:**
- 判断任意2顶点有无边，看对应的值matrix[i][j]是否等于1;
- 获取顶点i的度，即所在行数值之和,图中v1=1+0+1+0=2;
- 获取顶点i的所有邻接点，即将其所在行元素循环遍历一遍，值为1的顶点就是matrix[i]的邻接顶点。

### 2.1.2 Directed Graph Adjacency (有向图邻接矩阵)
如下列的有向图，使用邻接矩阵方式表示后的效果：
![](https://i.bmp.ovh/imgs/2021/04/1a131aba050d43cf.png)

> 有向图讲究入度与出度，顶点v1的入度是1,正好是第v1列元素之和；顶点v1的出度是2，即v2行各元素之和.
>
>与无向图同样的办法， 判断顶点vi到vj是否存在弧， 只需要查找矩阵中matrix[i][j]是否为1即可。 要求vi的所有邻接点就是将矩阵第i行元素扫描一遍， 查找arc[i][j]为1的顶点。
## 2.2 邻接表

[邻接表](https://github.com/lppgo/over-algorithm/blob/master/07-%E5%9B%BE/04-2-%E9%82%BB%E6%8E%A5%E8%A1%A8.md)

## 2.3 十字链表
[十字表](https://github.com/lppgo/over-algorithm/blob/master/07-%E5%9B%BE/04-3-%E5%8D%81%E5%AD%97%E9%93%BE%E8%A1%A8.md)
## 2.4 邻接多重表
[邻接多重表](https://github.com/lppgo/over-algorithm/blob/master/07-%E5%9B%BE/04-4-%E9%82%BB%E6%8E%A5%E5%A4%9A%E9%87%8D%E8%A1%A8.md)

## 2.5 边集数组
[边集数组](https://github.com/lppgo/over-algorithm/blob/master/07-%E5%9B%BE/04-5-%E8%BE%B9%E9%9B%86%E6%95%B0%E7%BB%84.md)