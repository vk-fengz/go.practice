package _46_LRUCache

// leetcode 解法, 构建自己的数据结构
// Least Recently Used (LRU) cache
// 最近使用的

type LRUCache struct {
	head, tail *Node
	Keys       map[int]*Node
	Cap        int
}

// Node 双向链表中的 node
type Node struct {
	Key, Value int
	Prev, Next *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Keys: make(map[int]*Node),
		Cap:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Keys[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Keys[key]; ok {
		node.Value = value
		this.Remove(node)
		this.Add(node)
		return
	} else {
		node := &Node{
			Key:   key,
			Value: value,
		}
		this.Keys[key] = node
		this.Add(node)
	}
	if len(this.Keys) > this.Cap {
		delete(this.Keys, this.tail.Key)
		this.Remove(this.tail)
	}

}

// Add 添加到双向链表的头  时间先-->后
func (this *LRUCache) Add(node *Node) {
	node.Prev = nil
	node.Next = this.head
	if this.head != nil {
		this.head.Prev = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
		this.tail.Next = nil
	}
}

// Remove the last node from the LRUCache   //
func (this *LRUCache) Remove(node *Node) {
	if node == this.head {
		this.head = this.head.Next
		node.Next = nil
		return
	}
	if node == this.tail {
		this.tail = node.Prev
		node.Prev.Next = nil
		node.Prev = nil
		return
	}

	// 中间
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// leetcode submit region end(Prohibit modification and deletion)
