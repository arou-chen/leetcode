package lrucache_146

type LRUCache struct {
	capacity int
	length   int
	data     map[int]*ListNode
	head     *ListNode
	tail     *ListNode
}

type ListNode struct {
	val  int
	prev *ListNode
	next *ListNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		length:   0,
		data:     make(map[int]*ListNode),
		head:     nil,
		tail:     nil,
	}
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.data[key]; ok {
		return node.val
	}
}

func (l *LRUCache) Put(key int, value int) {

}
