package main

import (
	"fmt"
)

type HashMap struct {
	buckets []*Node
}

type Node struct {
	key   string
	value int
	next  *Node
}

func (hashMap *HashMap) put(key string, value int) {
	index := hashMap.getBucketIndex(key)

	// get specific bucket
	node := hashMap.buckets[index]

	if node == nil {
		// create first node of bucket
		hashMap.buckets[index] = &Node{key: key, value: value}
		return
	}

	for ; ; node = node.next {
		if node.key == key {
			// rewrite value of node
			node.value = value
			return
		}
		if node.next == nil {
			break
		}
	}

	// add next node of existing node
	node.next = &Node{key: key, value: value}
}

func (hashMap *HashMap) get(key string) (int, bool) {
	index := hashMap.getBucketIndex(key)

	for node := hashMap.buckets[index]; node != nil; node = node.next {
		if node.key == key {
			return node.value, true
		}
	}

	return 0, false
}

func (hashMap *HashMap) printAll() {
	for i := 0; i < len(hashMap.buckets); i++ {
		if hashMap.buckets[i] == nil {
			fmt.Print("\n\t+----+\n", i, "\t|    |\n\t+----+\n")
			continue
		}
		fmt.Print("\n\t+----+\n", i, "\t| *  |")
		for node := hashMap.buckets[i]; node != nil; node = node.next {
			fmt.Print(" ---> ", node.key, ": ", node.value, " ")
		}
		fmt.Print("\n\t+----+\n")
	}
}

func construct(bucketsLen int) *HashMap {
	// initialize buckets of hash map
	return &HashMap{buckets: make([]*Node, bucketsLen)}
}

func getHash(key string) int {
	hash := 0

	for i := 0; i < len(key); i++ {
		hash += int(key[i])
	}

	return hash
}

func (hashMap *HashMap) getBucketIndex(key string) int {
	return getHash(key) % len(hashMap.buckets)
}

func main() {
	hashMap := construct(5)

	hashMap.put("bob", 10)
	hashMap.put("john", 15)
	hashMap.put("bob", 25)
	hashMap.put("elizabeth", 22)
	hashMap.put("buffalo", 33)

	hashMap.printAll()
}
