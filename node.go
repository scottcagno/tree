// *
// * Copyright 2014, Scott Cagno, All rights reserved.
// * BSD Licensed - sites.google.com/site/bsdc3license
// *
// * B+Tree Implementation
// *
package tree

// index page
type Idx struct {
	key   *[]byte
	child *Node
}

// data page
type Dat struct {
	key, val []byte
}

// node
type Node struct {
	idx []Idx // children aka. index page
	rec []Dat // records aka. data page
}

// check to see if node is leaf
func (self *Node) IsLeaf() bool {
	return len(self.idx) == 0
}
