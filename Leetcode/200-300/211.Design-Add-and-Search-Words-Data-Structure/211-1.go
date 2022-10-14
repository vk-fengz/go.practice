package main

import "fmt"

// 类型太多了, map实现trie 和 数组实现

func main() {
	wordDictionary := NewTrie()
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")
	fmt.Println(wordDictionary.Search("pad")) // 返回 False
	fmt.Println(wordDictionary.Search("bad")) // 返回 True
	fmt.Println(wordDictionary.Search(".ad"))
	fmt.Println(wordDictionary.Search("b.."))
	wordDictionary.Search("bad")
	wordDictionary.Search(".ad") // 返回 True
	wordDictionary.Search("b..") // 返回 True
}

// ===================

type WordDictionary = Trie

type Trie struct {
	isWord bool
	// 多使用数组
	children map[rune]*Trie
}

func NewTrie() Trie {
	return Trie{isWord: true, children: make(map[rune]*Trie)}
}

func (this *WordDictionary) AddWord(word string) {
	this.Insert(word)
}

func (this *WordDictionary) Search(word string) bool {
	m := &TrieMap{root: this, size: 0}
	return m.hasKeyWithPattern(word)
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

const (
	R = 256
)

// 借助 triemap;
// TrieNode TreeNode  the type of val
type TrieNode struct {
	val      interface{}
	children [R]*TrieNode
}

type TrieMap struct {
	root TrieNode
	size int
}

func NewTrieNode() *TrieNode {
	return &TrieNode{val: nil, children: [R]*TrieNode{}}
}

func NewTrieMap() *TrieMap {
	return &TrieMap{root: *NewTrieNode(), size: 0}
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
