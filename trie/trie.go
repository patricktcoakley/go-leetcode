package trie

type node struct {
	children   [26]*node
	isTerminal bool
}

func (n *node) contains(r rune) bool {
	return n.children[r-'a'] != nil
}

func (n *node) get(r rune) *node {
	return n.children[r-'a']
}

func (n *node) put(r rune, v *node) {
	n.children[r-'a'] = v
}

type Trie struct {
	head *node
}

func Constructor() Trie {
	return Trie{head: new(node)}
}

func (t *Trie) Insert(word string) {
	curr := t.head

	for _, v := range word {
		if !curr.contains(v) {
			curr.put(v, new(node))
		}

		curr = curr.get(v)
	}

	curr.isTerminal = true
}

func (t *Trie) Search(word string) bool {
	if prefix := t.searchPrefix(word); prefix != nil {
		return prefix.isTerminal
	}

	return false
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.searchPrefix(prefix) != nil
}

func (t *Trie) searchPrefix(prefix string) *node {
	curr := t.head

	for _, v := range prefix {
		if !curr.contains(v) {
			return nil
		}

		curr = curr.get(v)
	}

	return curr
}
