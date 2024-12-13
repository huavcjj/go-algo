package algorithm

type TrieNode struct {
	Word     rune
	Children map[rune]*TrieNode
	Term     string
}

type TrieTree struct {
	root *TrieNode
}

func (tree *TrieTree) Add(term string) {
	if len(term) == 0 {
		return
	}

	words := []rune(term)

	if tree.root == nil {
		tree.root = new(TrieNode)
	}

	tree.root.add(words, term)
}

func (node *TrieNode) add(words []rune, term string) {
	for _, word := range words {

		if node.Children == nil {
			node.Children = make(map[rune]*TrieNode)
		}

		if child, exists := node.Children[word]; exists {
			node = child
		} else {
			newNode := &TrieNode{Word: word}
			node.Children[word] = newNode
			node = newNode
		}
	}
	node.Term = term
}

func (tree *TrieTree) Find(prefix string) []string {
	if tree.root == nil || len(tree.root.Children) == 0 {
		return nil
	}
	words := []rune(prefix)
	node := tree.root.walk(words)
	if node == nil {
		return nil
	}

	return node.collect()
}

func (node *TrieNode) walk(words []rune) *TrieNode {
	for _, word := range words {
		if child, exists := node.Children[word]; exists {
			node = child
		} else {
			return nil
		}
	}
	return node
}

func (node *TrieNode) collect() []string {
	var terms []string
	if node.Term != "" {
		terms = append(terms, node.Term)
	}
	for _, child := range node.Children {
		terms = append(terms, child.collect()...)
	}
	return terms
}
