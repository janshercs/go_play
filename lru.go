package goplay

type LRUCache struct {
	capacity int
	cache    map[int]*node
	head     *node
	tail     *node
}

type node struct {
	left  *node
	right *node
	val   int
	key   int
}

func (n *node) pop() *node {
	n.left.right = n.right
	n.right.left = n.left
	return n
}

func (lru LRUCache) appendHead(n *node) {
	if lru.head.right == nil {
		lru.head.right = n
		n.left = lru.head
		lru.tail.left = n
		n.right = lru.tail
		return
	}
	lru.head.right.left = n
	n.right = lru.head.right
	n.left = lru.head
	lru.head.right = n
}

func (lru LRUCache) dropTail() {
	lru.tail.left = lru.tail.left.left
	lru.tail.left.right = lru.tail
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    map[int]*node{},
		head:     &node{},
		tail:     &node{},
	}
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.cache[key]; ok {
		n := node.pop()
		lru.appendHead(n)
		return node.val
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if lru.Get(key) != -1 { // key exists
		lru.cache[key].val = value
		return
	}

	if len(lru.cache) == lru.capacity {
		delete(lru.cache, lru.tail.left.key)
		lru.dropTail()
	}

	lru.cache[key] = &node{val: value, key: key}
	lru.appendHead(lru.cache[key])
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
