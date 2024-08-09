package main

import (
	"fmt"
	"linkedlist/linked"
)

func main() {
	list := &linked.LinkedList[int]{}
	list.AddLast(1)
	list.AddLast(2)
	list.AddLast(3)
	list.AddLast(4)
	list.PrintList()

	list.AddFirst(0)
	list.PrintList()
	fmt.Printf("index : %v\n", list.Get(1))

	list.Remove(2)
	list.PrintList()

	secondList := linked.New("A", "B", "C")
	secondList.PrintList()

	thirdList := linked.New(true, false, true)
	thirdList.PrintList()
}
