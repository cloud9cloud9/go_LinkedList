package main

import "fmt"

func main() {
	list := &LinkedList[int]{}
	list.addLast(1)
	list.addLast(2)
	list.addLast(3)
	list.addLast(4)
	list.printList()

	list.addFirst(0)
	list.printList()
	fmt.Printf("index : %v\n", list.get(1))

	list.remove(2)
	list.printList()
}
