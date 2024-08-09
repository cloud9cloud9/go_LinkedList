package linked

import (
	"fmt"
	"strings"
)

type LinkedList[A comparable] struct {
	Head *Node[A]
}

type Node[A comparable] struct {
	Value A
	Next  *Node[A]
}

func New[A comparable](el ...A) *LinkedList[A] {
	list := &LinkedList[A]{}
	for _, e := range el {
		list.AddLast(e)
	}
	return list
}

func (l *LinkedList[A]) AddFirst(el A) {
	firstNode := &Node[A]{Value: el}
	firstNode.Next = l.Head
	l.Head = firstNode

}

func (l *LinkedList[A]) AddLast(el A) {
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

func (l *LinkedList[A]) Get(el A) (index int) {
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

func (l *LinkedList[A]) Remove(el A) {
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

func (l *LinkedList[A]) PrintList() {
	currentHead := l.Head
	for currentHead != nil {
		fmt.Printf("references : %v\n", *currentHead)
		currentHead = currentHead.Next
	}
}

func (l LinkedList[T]) String() string {
	sb := new(strings.Builder)
	sb.WriteRune('{')

	for cur := l.Head; cur != nil; cur = cur.Next {
		sb.WriteString(fmt.Sprintf("%v,", cur.Value))
	}

	return strings.TrimRight(sb.String(), ",") + "}"
}
