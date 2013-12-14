package tree

// tree
type Tree struct {
	deg  int
	root *Node
}

// initialize and return tree
func InitTree(deg int) *Tree {
	return &Tree{
		deg:  int,
		root: InitNode(),
	}
}

// insert data
func (self *Tree) Insert(k, v []byte) bool {
	node := search(k, self.root)
	if len(node.records) == self.deg-1 {
		// split
		// leaf := InitNode()
	}
	node.records = append(node.records, Data{k, v})
}

// node
type Node struct {
	children []Index
	records  []Data
	infinity *Node
}

// initialize and return node
func InitNode() *Node {
	return &Node{
		children: make([]Index, 0),
		records:  make([]Data, 0),
		infinity: nil,
	}
}

// check to see if node is leaf
func (self *Node) IsLeaf() bool {
	return len(children) == 0
}

// check to see if node is full
func (self *Node) IsFull(n int) bool {
	return len(children) == n
}

// index struct
type Index struct {
	key   *[]byte
	child *Node
}

// initialize and return index struct
func InitIndex(k *[]byte) Index {
	return Index{
		key:   k,
		child: nil,
	}
}

// data struct
type Data struct {
	key, val []byte
}

// initialize and return data struct
func InitData(k, v []byte) Data {
	return Data{
		key: k,
		val: v,
	}
}
