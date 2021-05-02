- [1: Graph图的Summary](#1-graph图的summary)
  - [1.1 图的定义](#11-图的定义)
  - [1.2 graph的表示](#12-graph的表示)
  - [1.3 无向图(Undirected Graph)](#13-无向图undirected-graph)
  - [1.4 有向图(Directed Graph)](#14-有向图directed-graph)
    - [1.4.1 有向图顶点的出度和入度](#141-有向图顶点的出度和入度)
  - [1.5 稀疏图(Sparse Graph)和稠密图(Dense Graph)](#15-稀疏图sparse-graph和稠密图dense-graph)
  - [1.6 简单图(Simple Graph)](#16-简单图simple-graph)
  - [1.7 权(Weight)](#17-权weight)
- [2: Graph图的连通性Connectivity](#2-graph图的连通性connectivity)
  - [2.1 SubGraph子图](#21-subgraph子图)

# 1: Graph图的Summary

## 1.1 图的定义
   Graph 由顶点vertex和边edge组成，通常表示为:G(V,E)
   V: 顶点集，是有穷，非空的
   E: 边的集合，表示顶点与顶点之间的关系,可以为空


## 1.2 graph的表示
![图的表示](https://i.bmp.ovh/imgs/2021/04/053b8be97696674f.png)


## 1.3 无向图(Undirected Graph)
无向图：图中任意两个顶点之间都是无向边，则称图为无向图（Undirected graphs）.

完全无向图：如果无向图的任意两个顶点之间都存在边， 则称该图为无向完全图;
完全无向图如果有 n 个顶点，则会有 n(n-1)/2个边.

无向边：顶点$v_i$与顶点$v_j$之间没有方向，则称为无向边(Edge)。此时用无序偶对$(v_i,v_j)$表示.

## 1.4 有向图(Directed Graph)
有向图：如果图中任意两个顶点之间的边都是有向边，则称图为有向图（Directed graphs）.
有向边：顶点$v_i$与顶点$v_j$之间有方向，则称为有向边，也称为弧（Arc），此时用有序偶对$<v_i,v_j>$表示。

有向完全图：如果有向图的任意两个顶点之间都存在方向互为相反的两条弧，则为有向完全图
有向完全图含有 n 个顶点时，有 n(n-1) 条边。

### 1.4.1 有向图顶点的出度和入度
- **出度：Out-degree**，指有多少条边以该顶点为起点
- **入度：In-degree**，指有多少条边以该顶点为终点


## 1.5 稀疏图(Sparse Graph)和稠密图(Dense Graph)
很多场景中，无向图的实际边数、有向图的实际弧数，并没有达到完全图的数量。即具有n个顶点和e条边数的图， 
无向图 0≤e≤n(n-1)/2， 有向图0≤e≤n(n-1)。

有很少条边或弧的图称为稀疏图， 反之称为稠密图。 这里稀疏和稠密是 模糊的概念， 都是相对而言的。
## 1.6 简单图(Simple Graph)
简单图：图中不存在顶点到自身的边，且同一条边不重复出现，则称为简单图。
## 1.7 权(Weight)
有些图的边或弧具有与它相关的数字， 这种与图的边或弧相关的数叫做权（Weight）。

这些权可以表示从一个顶点到另一个顶点的距离或耗费。 这种带权的图通常称为网（Network）。

如下图所示，权即网中的各个顶点之间的距离：
![](https://i.bmp.ovh/imgs/2021/04/d2c9b34de82fecec.png)


# 2: Graph图的连通性Connectivity

[图的连通性](https://github.com/lppgo/over-algorithm/blob/master/07-%E5%9B%BE/02-%E5%9B%BE%E7%9A%84%E8%BF%9E%E9%80%9A%E6%80%A7.md)

连通图：在无向图G中， 如果从顶点v到顶点v'有路径， 则称v和v'是连通的。 如果对于图中任意两个顶点$v_i、 v_j∈ V$， $v_i$和$v_j$都是连通的， 则称G是连通图（Connected Graph） 。

总结：
```go
    图中顶点间存在路径， 两顶点存在路径则说明是连通的  
    如果路径最终回到起始点则称为环，当中不重复叫简单路径。  
    若任意两顶点都是连通的，则图就是连通图，有向则称强连通图。 
    图中有子图， 若子图极大连通则就是连通分量， 有向的则称强连通分量。
    无向图中连通且n个顶点n-1条边叫生成树。 
    有向图中一顶点入度为0其余顶点入度为1的叫有向树。 
    一个有向图由若干棵有向树构成生成森林。
```

## 2.1 SubGraph子图
