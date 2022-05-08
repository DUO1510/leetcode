package main

import "fmt"

func main() {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1)           // cache is {1=1}
	lRUCache.Put(2, 2)           // cache is {1=1, 2=2}
	fmt.Println(lRUCache.Get(1)) // return 1
	lRUCache.Put(3, 3)           // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	fmt.Println(lRUCache.Get(2)) // returns -1 (not found)
	lRUCache.Put(4, 4)           // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	fmt.Println(lRUCache.Get(1)) // return -1 (not found)
	fmt.Println(lRUCache.Get(3)) // return 3
	fmt.Println(lRUCache.Get(4)) // return 4
}

//LRU
//1. LRUCache(int capacity)初始化LRU缓存
//2. get(k)根据k获取v,如果有返回v否则返回-1
//3. put(k,v),如果已存在则更新，否则添加，如果LRU已满则淘汰最少使用的 - 链表
//4. get,put 要求 O(1) - map

type LRUCache struct {
	capacity, size int
	head, tail     *Node
	cache          map[int]*Node
}

type Node struct {
	value, key int
	pre, next  *Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		size:     0,
		head:     initiateNode(0, 0),
		tail:     initiateNode(0, 0),
		cache:    make(map[int]*Node),
	}
	lru.head.next = lru.tail
	lru.tail.pre = lru.head
	return lru
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.cache[key]; ok {
		l.moveToHead(node)
		return node.value
	}
	return -1
}

func (l *LRUCache) Put(key, value int) {
	if node, ok := l.cache[key]; ok {
		node.value = value
		l.moveToHead(node)
	} else {
		node := initiateNode(key, value)
		l.size++
		l.cache[key] = node
		l.addHeadNode(node)
		if l.size > l.capacity {
			l.size--
			rNode := l.removeTailNode()
			delete(l.cache, rNode.key)
		}
	}
}

func (l *LRUCache) addHeadNode(node *Node) {
	node.next = l.head.next
	node.pre = l.head
	l.head.next.pre = node
	l.head.next = node

}

func (l *LRUCache) removeTailNode() *Node {
	node := l.tail.pre
	removeNode(node)
	return node
}

func (l *LRUCache) moveToHead(node *Node) {
	removeNode(node)
	l.addHeadNode(node)
}

func removeNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func initiateNode(key, value int) *Node {
	node := &Node{
		value: value,
		key:   key,
	}
	return node
}
