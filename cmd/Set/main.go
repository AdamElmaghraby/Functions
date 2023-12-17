package main

import (
	"fmt"
)

type Set map[string]bool 

func (s Set) Add(element string) {
	s[element] = true
}


func (s Set) Contains(element string) bool {
	return s[element]
}

func (s Set) Remove(element string) {
	delete(s, element)
}

func (s Set) Size() int {
	return len(s)
}

func main() {
	mySet := make(Set)

	mySet.Add("Apple")
	mySet.Add("Banana")
	mySet.Add("Orange")

	fmt.Println("Apple Exists?:", mySet.Contains("Apple"))
	fmt.Println("Grape Exists?:", mySet.Contains("Grape"))

	mySet.Remove("Apple")

	fmt.Println("Apple Exists?:", mySet.Contains("Apple"))

	fmt.Println("Elements in the set:", mySet)
	fmt.Println("Size of set:", mySet.Size())
}