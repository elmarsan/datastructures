package main

type Tree struct {
	v    int
	r, l *Tree
}

// Returns max depth of the Tree.
func (t *Tree) depth() int {
	if t.r == nil && t.l == nil {
		return 1
	}

	rightDepth := t.r.depth()
	leftDepth := t.l.depth()

	if rightDepth > leftDepth {
		return rightDepth + 1
	}

	return leftDepth + 1
}

// Returns the sum of all node values.
func (t *Tree) Sum() int {
	sum := t.v

	if t.l != nil {
		sum += t.l.Sum()
	}

	if t.r != nil {
		sum += t.r.Sum()
	}

	return sum
}
