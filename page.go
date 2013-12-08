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
