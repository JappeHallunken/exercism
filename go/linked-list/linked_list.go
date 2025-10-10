package linkedlist

import "fmt"

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type Node struct {
	Value any
	next  *Node
	prev  *Node
}

type List struct {
	head *Node
	tail *Node
}

func NewList(elements ...any) *List {
	l := &List{}
	for _, el := range elements {
		l.Push(el)
	}
	return l
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v any) {
	newNode := &Node{Value: v}

	if l.head == nil {
		l.head, l.tail = newNode, newNode
		return
	}
	newNode.next = l.head
	l.head.prev, l.head = newNode, newNode
}

func (l *List) Push(v any) {
	newNode := &Node{Value: v}

	if l.tail == nil {
		l.head, l.tail = newNode, newNode
		return
	}
	newNode.prev = l.tail
	l.tail.next, l.tail = newNode, newNode
}

func (l *List) Shift() (any, error) {
	if l.head == nil {
		return nil, fmt.Errorf("list has no head")
	}
	rem := l.head
	value := rem.Value

	if l.head.next != nil {
		l.head = l.head.next
		l.head.prev = nil
	} else {
		l.head, l.tail = nil, nil
	}
	rem.next, rem.prev = nil, nil

	return value, nil
}

func (l *List) Pop() (any, error) {
	if l.tail == nil {
		return nil, fmt.Errorf("list has no tail")
	}
	rem := l.tail
	val := rem.Value

	if l.tail.prev != nil {
		l.tail = l.tail.prev
		l.tail.next = nil
	} else {
		l.tail, l.head = nil, nil
	}
	rem.next, rem.prev = nil, nil
	return val, nil
}

func (l *List) Reverse() {
	if l.head == nil || l.head == l.tail {
		return
	}
	current := l.head
	l.head, l.tail = l.tail, l.head

	for current != nil {
		current.next, current.prev = current.prev, current.next
		current = current.prev
	}
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
