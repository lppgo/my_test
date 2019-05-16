package main

import (
	"fmt"
	"sync"
)

//set 无序，不重复。无序--指的是是否安装添加顺序存储
// set是一个集合，本质上是一个list，但set里面的元素不能重复
//实现set
type Set struct {
	m            map[interface{}]bool //key可以保存所有类型
	sync.RWMutex                      //读写锁，保证线程安全
}

// New：返回一个set实例
func New() *Set {
	return &Set{
		m: map[interface{}]bool{},
	}
}

// Add：增加一个元素
func (s *Set) Add(item interface{}) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

// Remove：删除一个元素
func (s *Set) Remove(item interface{}) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

// Has：是否存在指定的元素
func (s *Set) Has(item interface{}) bool {
	//允许读
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

// List：转换为List
func (s *Set) List() []interface{} {
	s.RLock()
	defer s.RUnlock()
	var l []interface{}
	for value := range s.m {
		l = append(l, value)
	}
	return l
}

// Len：返回set元素个数
func (s *Set) Len() int {
	return len(s.m)
}

// Clear：清除set
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[interface{}]bool{}
}

// IsEmpty：set是否为空
func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

// Main测试
func main() {
	s := New()
	s.Add(1)
	s.Add("2")
	s.Add(3.5)
	s.Add(4)
	s.Add("5")

	fmt.Println("将set转成List:", s.List())

	fmt.Println("set的长度：", s.Len())

	fmt.Println("set是否存在指定元素：", s.Has("5"))

	fmt.Println("set删除元素：")
	s.Remove(4)

	fmt.Println("set是否为空：", s.IsEmpty())

	fmt.Println("将set转成List(end):", s.List())

}
