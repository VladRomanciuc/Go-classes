package main

import (
	"fmt"
)

/*
binary search tree in Go.

A nodes's leftside children should be less than or equal to itself, and the rightside children should be larger then itself.
*/

type node struct {
	left  *node
	right *node
	data  int
}

type tree struct {
	root *node
}

func (t *tree) insert(data int) *tree {
	if t.root == nil {
		t.root = &node{data: data}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *node) insert(data int) {
	if data <= n.data {
		if n.left == nil {
			n.left = &node{data: data}
			return
		}
		n.left.insert(data)
		return
	}
	if n.right == nil {
		n.right = &node{data: data}
		return
	}
	n.right.insert(data)
}

func main() {
	tree := &tree{}
	tree.insert(0)
	tree.insert(100)
	tree.insert(-20)
	tree.insert(-50)
	tree.insert(-15)
	tree.insert(-60)
	tree.insert(50)
	tree.insert(60)
	tree.insert(55)
	tree.insert(85)
	tree.insert(15)
	tree.insert(5)
	tree.insert(-10)
	tree.root.printR("M")
}

func (n *node) printR(label string) {
	if n == nil {
		return
	}

	fmt.Printf("%s: %v\n", label, n.data)
	n.left.printR(fmt.Sprintf("%s-L", label))
	n.right.printR(fmt.Sprintf("%s-R", label))
}