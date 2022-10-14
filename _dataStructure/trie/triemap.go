package trie

// 实现前缀树triemap trieset数据结构 相应的api
// 参考: https://mp.weixin.qq.com/s/hGrTUmM1zusPZZ0nA9aaNw

const (
	R = 256
)

// TrieNode TreeNode  the type of val
type TrieNode struct {
	val      interface{}
	children [R]*TrieNode
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
		p = p.children[ch]
	}
	return p
}

func NewTrieNode() *TrieNode {
	return &TrieNode{val: nil, children: [R]*TrieNode{}}
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
	node.children[ch-'a'] = t.put(node.children[ch-'a'], key, val, i+1)
	return node
}

func (t *TrieMap) get(key string) interface{} {
	x := getNode(t.root, key)

}

// hasKeyWithPattern 判断是和否存在前缀为 prefix 的键
func (t *TrieMap) hasKeyWithPattern(pattern string) bool {
	return hasKeyWithPattern(&t.root, pattern, 0)
}

func hasKeyWithPattern(node *TrieNode, pattern string, i int) bool {
	// 树枝不存在，即匹配失败
	if node == nil {
		return false
	}
	// 模式串走到头了，看看匹配到的是否是一个键
	if i == len(pattern) {
		return node.val == nil
	}

	c := pattern[i]

	if c != '.' {
		return hasKeyWithPattern(node.children[c-'a'], pattern, i+1)
	}
	// 遇到通配符
	for j := 0; j < R; j++ {
		if hasKeyWithPattern(node.children[j], pattern, i+1) {
			return true
		}
	}

	return false
}

func (t *TrieMap) contain(key string) bool {
	return t.get(key) != nil
}

func (t *TrieMap) getSize() int {
	return t.size
}
