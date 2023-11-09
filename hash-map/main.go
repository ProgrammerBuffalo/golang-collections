package main

func main() {
	hashMap := construct(5)

	hashMap.put("bob", 10)
	hashMap.put("john", 15)
	hashMap.put("bob", 25)
	hashMap.put("elizabeth", 22)
	hashMap.put("buffalo", 33)

	hashMap.printAll()
}
