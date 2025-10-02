package linkedlist

import "errors"

type Node struct {
	Value    interface{}
	next     *Node
	previous *Node
}

type List struct {
	first *Node
	last  *Node
}

func NewList(elements ...interface{}) *List {
	var list List
	for _, v := range elements {
		list.Push(v)
	}
	return &list
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.previous
}

func (l *List) Unshift(v interface{}) {
	if l.first == nil && l.last == nil {
		var node = Node{Value: v, next: nil, previous: nil}
		l.first = &node
		l.last = &node
		return
	}

	var node = Node{Value: v, next: l.first, previous: nil}
	l.first.previous = &node
	node.next = l.first
	l.first = &node
}

func (l *List) Push(v interface{}) {
	if l.first == nil && l.last == nil {
		var node = Node{Value: v, next: nil, previous: nil}
		l.first = &node
		l.last = &node
		return
	}

	var node = Node{Value: v, next: nil, previous: l.last}
	l.last.next = &node
	node.previous = l.last
	l.last = &node
}

func (l *List) Shift() (interface{}, error) {
	if l == nil {
		return nil, errors.New("empty list")
	}

	value := l.first.Value
	newFirst := l.first.next

	if newFirst == nil {
		l.first = nil
		l.last = nil
	} else {
		newFirst.previous = nil
		l.first = newFirst
	}

	return value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l == nil {
		return nil, errors.New("empty list")
	}

	value := l.last.Value
	newLast := l.last.previous

	if newLast == nil {
		l.last = nil
		l.first = nil
	} else {
		newLast.next = nil
		l.last = newLast
	}

	return value, nil
}

func (l *List) Reverse() {
	if l == nil {
		return
	}

	cur := l.first
	var tmp *Node

	for cur != nil {
		tmp = cur.previous
		cur.previous = cur.next
		cur.next = tmp
		cur = cur.previous
	}

	tmp = l.last
	l.last = l.first
	l.first = tmp
}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}
