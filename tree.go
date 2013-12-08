// *
// * Copyright 2014, Scott Cagno, All rights reserved.
// * BSD Licensed - sites.google.com/site/bsdc3license
// *
// * B+Tree Implementation
// *
package tree

import (
	"bytes"
)

// tree
type Tree struct {
	deg  int
	root *Node
}

// initialize tree
func InitTree(deg int) *Tree {
	return &Tree{
		deg: deg,
		root: &Node{
			make([]Idx, 0),
			make([]Dat, 0),
		},
	}
}

// search for a record with search-key value k
func (self *Tree) Search(k []byte) *Node {
	if self.root.IsLeaf() {
		return &self.root
	}
	// drill down...
	return nil
}

// insert
func (self *Tree) Insert(k, v []byte) bool {
	return false
}

// return
func (self *Tree) Return(k []byte) ([]byte, bool) {
	return nil, false
}

// delete
func (self *Tree) Delete(k []byte) bool {
	return false
}
