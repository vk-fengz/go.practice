package _08_Implement_Trie_Prefix_Tree

// 实现前缀树;
// 官方解法更符合 api逻辑 见2
func main() {

}

type Trie struct {
	isWord bool
	// 多使用数组
	children map[rune]*Trie
}

func Constructor() Trie {
	return Trie{isWord: true, children: make(map[rune]*Trie)}
}

func (this *Trie) Insert(word string) {
	parent := this
	for _, ch := range word {
		if child, ok := parent.children[ch]; ok {
			parent = child
		} else {
			newChild := &Trie{children: make(map[rune]*Trie)}
			parent.children[ch] = newChild
			parent = newChild
		}
	}
	parent.isWord = true
}

func (this *Trie) Search(word string) bool {
	parent := this
	for _, ch := range word {
		if child, ok := parent.children[ch]; ok {
			parent = child
			continue
		}
		return false
	}
	return parent.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	parent := this
	for _, ch := range prefix {
		if child, ok := parent.children[ch]; ok {
			parent = child
			continue
		}
		return false
	}
	return true
}
