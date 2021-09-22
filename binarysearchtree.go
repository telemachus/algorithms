package algorithms

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	count  int
	value  int
}

type Tree struct {
	root *Node
	size int
}

type fn func(int)

func NewBST() *Tree {
	return &Tree{}
}

func (t *Tree) Size() int {
	return t.size
}

func (t *Tree) Insert(v int) {
	var prevNode *Node
	currNode := t.root

	// Walk the tree until we find a spot to insert a new node.
	for currNode != nil {
		prevNode = currNode
		switch {
		case v < currNode.value:
			currNode = currNode.left
		case v > currNode.value:
			currNode = currNode.right
		default: // Our value is already present, so we can return early.
			currNode.count++
			t.size++
			return
		}
	}

	// Create a new node, place it, and increase the size of the tree.
	newNode := &Node{value: v, parent: prevNode}
	switch {
	case prevNode == nil:
		t.root = newNode
	case newNode.value < prevNode.value:
		prevNode.left = newNode
	default:
		prevNode.right = newNode
	}
	t.size++
}

func (t *Tree) InOrderWalk(f fn) {
	if t == nil {
		return
	}
	t.root.left.inOrderWalk(f)
	f(t.root.value)
	t.root.right.inOrderWalk(f)
	
}

func (n *Node) inOrderWalk(f fn) {
	if n == nil {
		return
	}
	n.left.inOrderWalk(f)
	f(n.value)
	n.right.inOrderWalk(f)
}

func (t *Tree) PreOrderWalk(f fn) {
	if t == nil {
		return
	}
	f(t.root.value)
	t.root.left.preOrderWalk(f)
	t.root.right.preOrderWalk(f)
}

func (n *Node) preOrderWalk(f fn) {
	if n == nil {
		return
	}
	f(n.value)
	n.left.preOrderWalk(f)
	n.right.preOrderWalk(f)
}

func (t *Tree) PostOrderWalk(f fn) {
	if t == nil {
		return
	}
	t.root.left.postOrderWalk(f)
	t.root.right.postOrderWalk(f)
	f(t.root.value)
}

func (n *Node) postOrderWalk(f fn) {
	if n == nil {
		return
	}
	n.left.postOrderWalk(f)
	n.right.postOrderWalk(f)
	f(n.value)
}
