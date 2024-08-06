package main

import "fmt"

type LinkedList[A comparable] struct {
	Head *Node[A]
}

type Node[A comparable] struct {
	Value A
	Next  *Node[A]
}

func (l *LinkedList[A]) addFirst(el A) {
	firstNode := &Node[A]{Value: el}
	firstNode.Next = l.Head
	l.Head = firstNode

}

func (l *LinkedList[A]) addLast(el A) {
	if l.Head == nil {
		l.Head = &Node[A]{Value: el}
		return
	}
	currentNode := l.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = &Node[A]{Value: el}

}

func (l *LinkedList[A]) get(el A) (index int) {
	if l.Head == nil {
		return -1
	}

	if l.Head.Value == el {
		return 0
	}

	current := l.Head
	index = 0
	for current != nil {
		index++
		if current.Value == el {
			return index
		}
		current = current.Next
	}
	return -1
}

func (l *LinkedList[A]) remove(el A) {
	if l.Head == nil {
		return
	}
	current := l.Head
	if current.Value == el {
		l.Head = current.Next
		return
	}
	currentNode := l.Head
	for currentNode.Next != nil {
		if currentNode.Next.Value == el {
			currentNode.Next = currentNode.Next.Next
			return
		}
		currentNode = currentNode.Next
	}
}

func (l *LinkedList[A]) printList() {
	currentHead := l.Head
	for currentHead != nil {
		fmt.Printf("references : %v\n", *currentHead)
		currentHead = currentHead.Next
	}
}
