[TOC]

```go
1:go 切片实现栈 (后进后出)
2:go 切片实现队列 (先进先出)
3: int和[]byte互转
4: go Generate UUID
5: IsEmpty()判断给的值是否为空
6: 二分法对Slice进行插入排序
7: RSA加密解密
8：获取远程客户端的IP,讲IPV4转uint32
```
# 1:go 切片实现栈 (后进后出)
```go
    // 创建栈
    stack:=make([]int,0)
    // push压入
    stack=append(stack,10)
    // pop弹出
    v:=stack[len(stack)-1]
    stack=stack[:len(stack)-1]
    // 检查栈空
    len(stack)==0
```
# 2:go 切片实现队列 (先进先出)
```go
    // 创建队列
    queue:=make([]int,0)
    // enqueue入队
    queue=append(queue,10)
    // dequeue出队
    v:=queue[0]
    queue=queue[1:]
    // 长度0为空
    len(queue)==0
```

# 3：sort排序
```go
    // int排序
    sort.Ints([]int{})
    // 字符串排序
    sort.Strings([]string{})
    // 自定义排序
    sort.Slice(s,func(i,j int)bool{return s[i]<s[j]})
```
