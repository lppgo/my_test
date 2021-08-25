package cmd

import (
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

type Hash func(data []byte) uint32

type HashRing struct {
	sync.RWMutex
	hash     Hash           // 计算hash的函数
	replicas int            // 副本数，这里影响虚拟节点(virtual node)的个数
	keys     []int          // 有序的列表，从大到小排序
	hashMap  map[int]string // 虚拟node和物理节点的映射
}

// New 初始化hash ring .
func New(replicas int, fn Hash) *HashRing {
	hr := &HashRing{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}

	if hr.hash == nil {
		// default use crc32来计算hash值
		hr.hash = crc32.ChecksumIEEE
	}
	return hr
}

//
func (c *HashRing) IsEmpty() bool {
	c.RLock()
	defer c.RUnlock()
	if len(c.keys) > 0 {
		return false
	}
	return true
}

// AddNodes hash ring 添加节点.
func (c *HashRing) AddNodes(keys ...string) {
	c.RLock()
	defer c.RUnlock()
	for _, key := range keys {
		for i := 0; i < c.replicas; i++ {
			// hash 值= hash(i+key)
			hash := int(c.hash([]byte(strconv.Itoa(i) + key)))
			c.keys = append(c.keys, hash)
			c.hashMap[hash] = key
		}
	}
	sort.Ints(c.keys)
}

// GetNode 一致性hash请求，获取真实node .
func (c *HashRing) GetNode(key string) string {
	if c.IsEmpty() {
		return ""
	}

	// 根据输入的key计算一个hash值
	hash := int(c.hash([]byte(key)))
	// 查看hash值在哪个值域范围，选择对应的 virtual node
	f := func(i int) bool {
		return c.keys[i] >= hash
	}
	idx := sort.Search(len(c.keys), f)
	if idx == len(c.keys) {
		idx = 0
	}

	// 选择对应的物理节点
	return c.hashMap[c.keys[idx]]
}
