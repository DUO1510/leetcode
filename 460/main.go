package main

import "fmt"

func main() {
	l := Constructor(2)
	l.Put(1, 1)
	l.Put(2, 2)
	fmt.Println(l.Get(1))
	l.Put(3, 3)
	fmt.Println(l.Get(2))
	fmt.Println(l.Get(3))
	l.Put(4, 4)
	fmt.Println(l.Get(1))
	fmt.Println(l.Get(3))
	fmt.Println(l.Get(4))

	//l.Put(3, 1)
	//l.Put(2, 1)
	//l.Put(2, 2)
	//l.Put(4, 4)
	//fmt.Println(l.Get(2))
	//l.Put(0, 0)
	//fmt.Println(l.Get(0))
}

type Node struct {
	key, value, count int
	pre, next         *Node
}

func initNode(k, v int) *Node {
	return &Node{k, v, 1, nil, nil}
}

type DoubleLinkList struct {
	head  *Node
	tail  *Node
	total int
}

func NewDLList() *DoubleLinkList {
	d := &DoubleLinkList{
		head: initNode(0, 0),
		tail: initNode(0, 0),
	}

	d.head.next = d.tail
	d.tail.pre = d.head
	return d
}

type LFUCache struct {
	capacity, size, min int
	//保存k对应的节点
	cache map[int]*Node
	//保存次数对应的节点
	cntMap map[int]*DoubleLinkList
}

func Constructor(capacity int) LFUCache {
	l := LFUCache{
		capacity: capacity,
		size:     0,
		min:      0,
		cache:    make(map[int]*Node),
		cntMap:   make(map[int]*DoubleLinkList),
	}

	return l
}

func (d *DoubleLinkList) removeTailNode() *Node {
	node := d.tail.pre
	d.removeNode(node)
	return node
}

func (d *DoubleLinkList) addHeadNode(node *Node) {
	node.next = d.head.next
	node.pre = d.head
	d.head.next.pre = node
	d.head.next = node
	d.total++
}

func (d *DoubleLinkList) removeNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
	d.total--
}

func (l *LFUCache) Get(key int) int {
	if node, ok := l.cache[key]; ok {
		//存在，移除对应次数链表的节点
		l.cntMap[node.count].removeNode(node)
		//该节点次数就是最小值&&链表为空
		if l.min == node.count && l.cntMap[node.count].total == 0 {
			delete(l.cntMap, node.count)
			l.min = node.count + 1
		}

		node.count++
		if _, ok := l.cntMap[node.count]; !ok {
			l.cntMap[node.count] = NewDLList()
		}
		l.cntMap[node.count].addHeadNode(node)
		return node.value
	}
	return -1
}

func (l *LFUCache) Put(key int, value int) {
	//如果存在 更新访问次数
	if l.Get(key) > -1 {
		l.cache[key].value = value
		return
	}

	//如果不存在 且缓存满了，需要删除
	if l.size == l.capacity {
		if l.capacity == 0 {
			return
		}
		dl := l.cntMap[l.min]
		node := dl.removeTailNode()
		delete(l.cache, node.key)
		l.size--
	}
	// 新建节点
	//当前个数+1
	l.size++
	//初始化一个新节点
	node := initNode(key, value)
	//节点放到cache中
	l.cache[key] = node
	//带有次数1的节点链表存在
	if _, ok := l.cntMap[1]; !ok {
		//新建一个节点列表
		l.cntMap[1] = NewDLList()
	}
	l.cntMap[1].addHeadNode(node)
	l.min = 1
}
