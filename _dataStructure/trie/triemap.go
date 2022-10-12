package trie

// 实现前缀树triemap trieset数据结构 相应的api
// 参考: https://mp.weixin.qq.com/s/hGrTUmM1zusPZZ0nA9aaNw

const (
	R = 256
)

// TrieNode TreeNode  the type of val
type TrieNode struct {
	val      interface{}
	children [R]TrieNode
}

type TrieMap struct {
	root TrieNode
	size int
}

func getNode(node TrieNode, key string) *TrieNode {
	p := &node
	// 从节点node开始搜索key
	for _, ch := range key {
		if p == nil {
			return nil
		}
		ch -= 'a'
		p = &p.children[ch]
	}
	return p
}

func NewTrieNode() *TrieNode {
	return &TrieNode{val: nil, children: [R]TrieNode{}}
}

// put 定义：向以 node 为根的 Trie 树中插入 key[i..]，返回插入完成后的根节点
func (t *TrieMap) put(node *TrieNode, key string, val interface{}, i int) *TrieNode {
	if node == nil {
		node = NewTrieNode()
	}
	if i == len(key) {
		node.val = val
		return node
	}
	ch := key[i]
	node.children[ch-'a'] = *t.put(&node.children[ch-'a'], key, val, i+1)
	return node
}

func (t *TrieMap) get(key string) interface{} {
	x := getNode(t.root, key)

}

func (t *TrieMap) contain(key string) bool {
	return t.get(key) != nil
}

func (t *TrieMap) getSize() int {
	return t.size
}
