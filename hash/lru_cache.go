package hash

import "container/list"

type KeyValuePair struct {
	key   int
	value int
}

type LRUCache struct {
	cache    map[int]*list.Element
	list     *list.List
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{cache: make(map[int]*list.Element, capacity), list: list.New(), capacity: capacity}
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.cache[key]; ok {
		l.list.MoveToBack(node)
		return node.Value.(*KeyValuePair).value
	}

	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.cache[key]; ok {
		node.Value.(*KeyValuePair).value = value
		l.list.MoveToBack(node)
		return
	}

	if len(l.cache) == l.capacity {
		head := l.list.Front()
		delete(l.cache, head.Value.(*KeyValuePair).key)
		l.list.Remove(head)
	}

	l.cache[key] = l.list.PushBack(&KeyValuePair{key, value})
}
