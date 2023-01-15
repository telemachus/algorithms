package algorithms

import (
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func RecursiveConcurrentWalk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walk(t.Right, ch)
	}
}

// IterativeConcurrentWalk walks the tree t sending all values from the tree to
// the channel ch.
func IterativeConcurrentWalk(t *tree.Tree, ch chan int) {
	stack := []*tree.Tree{}
	current := t
	for {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else if len(stack) > 0 {
			current, stack = pop(stack)
			ch <- current.Value
			current = current.Right
		} else {
			break
		}
	}
	close(ch)
}

func pop(stack []*tree.Tree) (*tree.Tree, []*tree.Tree) {
	n := len(stack) - 1
	t := stack[n]
	stack[n] = nil
	return t, stack[:n]
}

func MorrisIterativeConcurrentWalk(t *tree.Tree, ch chan int) {
	current := t
	for current != nil {
		if current.Left == nil {
			ch <- current.Value
			current = current.Right
		} else {
			pre := current.Left
			for pre.Right != nil && pre.Right != current {
				pre = pre.Right
			}
			if pre.Right == nil {
				pre.Right = current
				current = current.Left
			} else {
				pre.Right = nil
				ch <- current.Value
				current = current.Right
			}
		}
	}
	close(ch)
}
