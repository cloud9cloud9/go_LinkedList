package main

type LinkedInterface[T any] interface {
	addFirst(el T)
	addLast(el T)
	get(el T) (index int)
	remove(el T)
	printList()
}
