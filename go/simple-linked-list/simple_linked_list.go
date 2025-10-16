package linkedlist

import "fmt"

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	el   int
	next *Node
}

func New(elements []int) *List {
	list := &List{}
	for _, el := range elements {
		list.Push(el)
	}
	return list
}

func (l *List) Size() int {
	current := l.head
	counter := 0

	for current != nil {
		counter++
		current = current.next
	}
	return counter
}

func (l *List) Push(element int) {
	newNode := &Node{el: element}

	if l.head == nil {
		l.head, l.tail = newNode, newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
}

func (l *List) Pop() (int, error) {
	current := l.head
	if current == nil {
		return -1, fmt.Errorf("no element to pop")
	}
	if current.next == nil {
		value := current.el
		l.head = nil
		l.tail = nil
		return value, nil
	}

	for current.next != l.tail {
		current = current.next
	}
	value := l.tail.el
	current.next = nil
	l.tail = current

	return value, nil
}

func (l *List) Array() []int {
	result := make([]int, l.Size())
	current := l.head
	for i := range result {
		result[i] = current.el
		current = current.next
	}
	return result
}

func (l *List) Reverse() *List {
	var prev *Node = nil
	current := l.head
	l.tail = l.head

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
	return l
}
