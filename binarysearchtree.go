package algorithms

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	count  int
	value  int
}

func (n *Node) Value() int {
	return n.value
}

func (n *Node) Count() int {
	return n.count
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
	// Walk the tree until we find a spot to insert a new node.
	var prevNode *Node
	for currNode := t.root; currNode != nil; {
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
	newNode := &Node{value: v, parent: prevNode, count: 1}
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
	for i := t.root.count; i > 0; i-- {
		f(t.root.value)
	}
	t.root.right.inOrderWalk(f)

}

func (n *Node) inOrderWalk(f fn) {
	if n == nil {
		return
	}
	n.left.inOrderWalk(f)
	for i := n.count; i > 0; i-- {
		f(n.value)
	}
	n.right.inOrderWalk(f)
}

func (t *Tree) PreOrderWalk(f fn) {
	if t == nil {
		return
	}
	for i := t.root.count; i > 0; i-- {
		f(t.root.value)
	}
	t.root.left.preOrderWalk(f)
	t.root.right.preOrderWalk(f)
}

func (n *Node) preOrderWalk(f fn) {
	if n == nil {
		return
	}
	for i := n.count; i > 0; i-- {
		f(n.value)
	}
	n.left.preOrderWalk(f)
	n.right.preOrderWalk(f)
}

func (t *Tree) PostOrderWalk(f fn) {
	if t == nil {
		return
	}
	t.root.left.postOrderWalk(f)
	t.root.right.postOrderWalk(f)
	for i := t.root.count; i > 0; i-- {
		f(t.root.value)
	}
}

func (n *Node) postOrderWalk(f fn) {
	if n == nil {
		return
	}
	n.left.postOrderWalk(f)
	n.right.postOrderWalk(f)
	for i := n.count; i > 0; i-- {
		f(n.value)
	}
}

func (t *Tree) InOrderIterativeWalk(f fn) {
	if node := t.root; node != nil {
		for node.left != nil {
			node = node.left
		}
		for node != nil {
			for i := node.count; i > 0; i-- {
				f(node.value)
			}
			if node.right != nil {
				node = node.right
				for node.left != nil {
					node = node.left
				}
			} else {
				for node.parent != nil && node.parent.right == node {
					node = node.parent
				}
				node = node.parent
			}
		}
	}
}

func (t *Tree) PreOrderIterativeWalk(f fn) {
	if node := t.root; node != nil {
		for node.left != nil {
			for i := node.count; i > 0; i-- {
				f(node.value)
			}
			node = node.left
		}
		f(node.value)
		for node != nil {
			if node.right != nil {
				node = node.right
				for i := node.count; i > 0; i-- {
					f(node.value)
				}
				for node.left != nil {
					node = node.left
					for i := node.count; i > 0; i-- {
						f(node.value)
					}
				}
			} else {
				for node.parent != nil && node.parent.right == node {
					node = node.parent
				}
				node = node.parent
			}
		}
	}
}

func (t *Tree) PostOrderIterativeWalk(f fn) {
	if node := t.root; node != nil {
		for node.left != nil {
			node = node.left
		}
		for node != nil {
			if node.right != nil {
				node = node.right
				for node.left != nil {
					node = node.left
				}
			} else {
				for node.parent != nil && node.parent.right == node {
					for i := node.count; i > 0; i-- {
						f(node.value)
					}
					node = node.parent
				}
				for i := node.count; i > 0; i-- {
					f(node.value)
				}
				node = node.parent
			}
		}
	}
}

func (t *Tree) Search(v int) *Node {
	return t.root.search(v)
}

func (n *Node) search(v int) *Node {
	for n != nil && v != n.value {
		switch {
		case v < n.value:
			n = n.left
		case v > n.value:
			n = n.right
		}
	}
	return n
}

func (t *Tree) Min() *Node {
	return t.root.min()
}

func (n *Node) min() *Node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func (t *Tree) Max() *Node {
	return t.root.max()
}

func (n *Node) max() *Node {
	for n.right != nil {
		n = n.right
	}
	return n
}

func (t *Tree) Successor(currNode *Node) *Node {
	if currNode.right != nil {
		return currNode.right.min()
	}

	prevNode := currNode.parent
	for prevNode != nil && currNode == prevNode.right {
		currNode = prevNode
		prevNode = prevNode.parent
	}
	return prevNode
}

func (t *Tree) Predecessor(currNode *Node) *Node {
	if currNode.left != nil {
		return currNode.left.max()
	}

	prevNode := currNode.parent
	for prevNode != nil && currNode == prevNode.left {
		currNode = prevNode
		prevNode = prevNode.parent
	}
	return prevNode
}

func (t *Tree) Delete(v int) {
	deadNode := t.Search(v)
	if deadNode == nil {
		return
	}

	switch {
	case deadNode.count > 1:
		deadNode.count--
	case deadNode.left == nil:
		t.transplant(deadNode, deadNode.right)
	case deadNode.right == nil:
		t.transplant(deadNode, deadNode.left)
	default:
		succNode := deadNode.right.min()
		if succNode.parent != deadNode {
			t.transplant(succNode, succNode.right)
			succNode.right = deadNode.right
			succNode.right.parent = succNode
		}
		t.transplant(deadNode, succNode)
		succNode.left = deadNode.left
		succNode.left.parent = succNode
	}
	t.size--
}

func (t *Tree) transplant(deadNode, succNode *Node) {
	switch {
	case deadNode.parent == nil:
		t.root = succNode
	case deadNode == deadNode.parent.left:
		deadNode.parent.left = succNode
	default:
		deadNode.parent.right = succNode
	}
	if succNode != nil {
		succNode.parent = deadNode.parent
	}
}
