package main

import (
	"fmt"
)

func main() {
	list := SinglyLinkedList[int]{}

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	fmt.Println("Initial list")
	list.Print()

	list.Remove(3)

	fmt.Println("After removed element")
	list.Print()

	list.Add(5)
	fmt.Println("After added element")
	list.Print()

}
