package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left *Tree
	Value int
	Right *Tree
}

func traverse(t *Tree) {
	if t == nil {
		return
	}

	fmt.Print(t.Value, " ")
	traverse(t.Left)
	traverse(t.Right)
}

func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())

	for i := 0; i < 2 * n; i++ {
		tmp := rand.Intn(n)
		t = insert(t, tmp)
	}

	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}

	if v == t.Value {
		return t
	}

	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}

	t.Right = insert(t.Right, v)

	return t
}

func main() {
	tree := create(30)
	traverse(tree)
	fmt.Println()
	fmt.Println("Tree Root ", tree.Value)
}
