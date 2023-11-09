package main

import "fmt"

// Generic node of list
type Node[TVal any] struct {
	value TVal
	next  *Node[TVal]
}

type SinglyLinkedList[TVal int | string] struct {
	head *Node[TVal]
}

func (list *SinglyLinkedList[TVal]) Add(value TVal) {
	// create new node
	newNode := Node[TVal]{value: value}

	if list.head == nil {
		list.head = &newNode
		return
	}

	for curr := list.head; curr != nil; curr = curr.next {
		if curr.next == nil {
			// set address of previous node to next new node
			curr.next = &newNode
			return
		}
	}
}

func (list *SinglyLinkedList[TVal]) Remove(value TVal) {
	if list.head == nil {
		return
	}

	if list.head.value == value {
		list.head = list.head.next
		return
	}

	for curr, prev := list.head.next, list.head; curr != nil; prev, curr = curr, curr.next {
		if curr.value == value {
			prev.next = curr.next
			return
		}
	}
}

func (list *SinglyLinkedList[TVal]) Print() {
	for curr := list.head; curr != nil; curr = curr.next {
		fmt.Print(curr.value, " ")
	}
	fmt.Println()
}
