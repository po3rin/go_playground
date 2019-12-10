package trie

const size = 230

type TrieNode struct {
	children   [size]*TrieNode
	endOfWords bool
}

func GetNode() *TrieNode {
	node := &TrieNode{
		endOfWords: false,
	}

	for i := 0; i < size; i++ {
		node.children[i] = nil
	}

	return node
}

func (root *TrieNode) Insert(key string) {
	for i := 0; i < len(key); i++ {
		index := key[i]
		if root.children[index] == nil {
			root.children[index] = GetNode()
		}
		root = root.children[index]
	}
	root.endOfWords = true
}

func Search(root *TrieNode, key string) bool {
	for i := 0; i < len(key); i++ {
		index := key[i]
		if root.children[index] != nil {
			root = root.children[index]
		} else {
			return false
		}
	}
	return (root != nil && root.endOfWords)
}
