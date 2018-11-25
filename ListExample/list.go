package main

import (
	"fmt"
)

type Node struct {
	next, prev *Node
	Value      interface{}
}

// Node constructor
func CreateNode(prev *Node, v interface{}, next *Node) *Node {
	node := new(Node)
	node.prev = prev
	node.next = next
	node.Value = v
	return node
}

func (n *Node) clear() interface{} {
	v := n.Value
	n.next = nil
	n.prev = nil
	n.Value = nil
	return v
}

type List struct {
	len        int
	head, tail *Node
}

func (l *List) String() string {
	s := "List["
	node := l.head
	for idx := 0; node != nil; idx++ {
		s += fmt.Sprint(node.Value, " ")
		node = node.next
	}
	s += "\b]"
	return s
}

func (l *List) linkFirst(v interface{}) {
	f := l.head
	nn := CreateNode(nil, v, f)
	l.head = nn

	if f == nil {
		l.tail = nn
	} else {
		f.prev = nn
	}
	l.len++
}

func (l *List) linkLast(v interface{}) {
	t := l.tail
	nn := CreateNode(t, v, nil)
	l.tail = nn

	if t == nil {
		l.head = nn
	} else {
		t.next = nn
	}
	l.len++
}

func (l *List) linkBefore(v interface{}, successor *Node) {
	predecessor := successor.prev
	nn := CreateNode(predecessor, v, successor)
	successor.prev = nn

	if predecessor == nil {
		l.head = nn
	} else {
		predecessor.next = nn
	}
	l.len++
}

func (l *List) unlink(node *Node) interface{} {
	v := node.Value
	next := node.next
	prev := node.prev

	if prev == nil {
		l.head = next
	} else {
		prev.next = next
	}

	if next == nil {
		l.tail = prev
	} else {
		next.prev = prev
	}

	l.len--
	node.clear()
	return v
}

func (l *List) unlinkFirst() interface{} {
	node := l.head

	v := node.Value
	next := node.next

	node.clear()

	l.head = next
	if next == nil {
		l.tail = nil
	} else {
		next.prev = nil
	}
	l.len--
	return v
}

func (l *List) unlinkLast() interface{} {
	node := l.tail

	v := node.Value
	prev := node.prev

	node.clear()

	l.tail = prev
	if prev == nil {
		l.head = nil
	} else {
		prev.next = nil
	}
	l.len--
	return v
}

func (l *List) size() int {
	return l.len
}

func (l *List) add(v interface{}) {
	l.linkLast(v)
}

func (l *List) remove(v interface{}) bool {
	for x := l.head; x != nil; x = x.next {
		if v == x.Value {
			l.unlink(x)
			return true
		}
	}
	return false
}

func main() {
	l := new(List)

	l.add(10)
	l.add(20)
	l.add(30)
	l.add(40)
	l.add(50)
	fmt.Println(l)

	fmt.Printf("Size: %d\n", l.size())
	if l.remove(40) {
		fmt.Println("Removed successfully.")
	} else {
		fmt.Println("No such element in List")
	}

	fmt.Println(l)
	if l.remove(60) {
		fmt.Println("Removed successfully.")
	} else {
		fmt.Println("No such element in List")
	}
}

