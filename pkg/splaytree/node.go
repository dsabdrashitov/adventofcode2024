package splaytree

import "fmt"

type node[K any, V any, A any] struct {
	key   K
	value V
	agg   A
	left  *node[K, V, A]
	right *node[K, V, A]
}

func makeNode[K any, V any, A any](key K, value V, agg Aggregator[K, V, A], left *node[K, V, A], right *node[K, V, A]) *node[K, V, A] {
	return &node[K, V, A]{key, value, agg(key, value, nil, nil), left, right}
}

func (n *node[K, V, A]) reassign(agg Aggregator[K, V, A], left *node[K, V, A], right *node[K, V, A]) *node[K, V, A] {
	return makeNode(n.key, n.value, agg, left, right)
}

func (root *node[K, V, A]) splay(key K, cmp Comparator[K], agg Aggregator[K, V, A]) *node[K, V, A] {
	if root == nil {
		return nil
	}
	path := make([]*node[K, V, A], 0)
	for t := root; t != nil; {
		path = append(path, t)
		switch c := cmp(t.key, key); c {
		case GREATER:
			t = t.left
		case LESS:
			t = t.right
		case EQUAL:
			t = nil
		default:
			panic(fmt.Errorf("unknown comparision result: %v", c))
		}
	}
	for len(path) > 1 {
		x0 := path[len(path)-1]
		p0 := path[len(path)-2]
		if len(path) == 2 {
			var x1 *node[K, V, A]
			var p1 *node[K, V, A]
			if cmp(p0.key, key) == GREATER {
				p1 = p0.reassign(agg, x0.right, p0.right)
				x1 = x0.reassign(agg, x0.left, p1)
			} else {
				p1 = p0.reassign(agg, p0.left, x0.left)
				x1 = x0.reassign(agg, p1, x0.right)
			}
			path[len(path)-2] = x1
			path = path[:len(path)-1]
		} else {
			g0 := path[len(path)-3]
			var x1 *node[K, V, A]
			var p1 *node[K, V, A]
			var g1 *node[K, V, A]
			if cmp(g0.key, key) == GREATER {
				if cmp(p0.key, key) == GREATER {
					g1 = g0.reassign(agg, p0.right, g0.right)
					p1 = p0.reassign(agg, x0.right, g1)
					x1 = x0.reassign(agg, x0.left, p1)
				} else {
					g1 = g0.reassign(agg, x0.right, g0.right)
					p1 = p0.reassign(agg, p0.left, x0.left)
					x1 = x0.reassign(agg, p1, g1)
				}
			} else {
				if cmp(p0.key, key) == LESS {
					g1 = g0.reassign(agg, g0.left, p0.left)
					p1 = p0.reassign(agg, g1, x0.left)
					x1 = x0.reassign(agg, p1, x0.right)
				} else {
					g1 = g0.reassign(agg, g0.left, x0.left)
					p1 = p0.reassign(agg, x0.right, p0.right)
					x1 = x0.reassign(agg, g1, p1)
				}
			}
			path[len(path)-3] = x1
			path = path[:len(path)-2]
		}
	}
	// root != nil => len(path) >= 1
	return path[0]
}

func (root *node[K, V, A]) min(agg Aggregator[K, V, A]) *node[K, V, A] {
	if root == nil {
		return nil
	}
	t := root
	for t.left != nil {
		x0 := t.left
		p0 := t
		p1 := p0.reassign(agg, x0.right, p0.right)
		x1 := x0.reassign(agg, x0.left, p1)
		t = x1
	}
	return t
}

func (root *node[K, V, A]) max(agg Aggregator[K, V, A]) *node[K, V, A] {
	if root == nil {
		return nil
	}
	t := root
	for t.right != nil {
		x0 := t.right
		p0 := t
		p1 := p0.reassign(agg, p0.left, x0.left)
		x1 := x0.reassign(agg, p1, x0.right)
		t = x1
	}
	return t
}

func merge[K, V, A any](agg Aggregator[K, V, A], left *node[K, V, A], right *node[K, V, A]) *node[K, V, A] {
	if left == nil {
		return right
	}
	left = left.max(agg)
	return left.reassign(agg, left.left, right)
}
