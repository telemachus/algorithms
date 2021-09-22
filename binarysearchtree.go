package algorithms

type Node struct {
	parent *Node
	lChild *Node
	rChild *Node
	Count int
	Value int
}

type Tree struct {
	root *Node
	size int
}

func NewBST() *Tree {
	return &Tree{}
}

func (t *Tree) Size() int {
	return t.size
}
