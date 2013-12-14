package tree

import (
	"bytes"
)

// return proper leaf for given key
func search(k []byte, node *Node) *Node {
	if node.IsLeaf() {
		return node
	}
	for i := range node.children {
		switch bytes.Compare(k, node.children[i].key) {
		case -1, 0:
			return search(k, node.children[i].child)
		}
	}
	return search(k, node.infinity)
}
