package main

func main() {
	list := SinglyLinkedList[int]{}

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	list.Print()

	list.Remove(3)

	list.Print()

	list.Add(5)

	list.Print()

}
