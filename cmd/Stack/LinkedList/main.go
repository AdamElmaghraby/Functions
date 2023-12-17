package main

import (
	"fmt"
)

type Toy struct {
	Name string
	Color string
	Next *Toy
}

type Box struct {
	TopToy *Toy
}

func (b *Box) AddToy(name, color string) {
	newToy  := &Toy{Name: name, Color: color, Next:nil}

	if b.TopToy == nil {
		b.TopToy = newToy
		return
	}

	current := b.TopToy
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newToy

}


func (b *Box) DisplayToys() {
	current := b.TopToy 
	for current != nil {
		fmt.Printf("Toy: %s (Color: %s) -> ", current.Name, current.Color)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	myBox := Box{}

	myBox.AddToy("Teddy Bear", "Brown")
	myBox.AddToy("Building Blocks", "MultiColor")
	myBox.AddToy("Doll", "Pink")

	fmt.Println("Toys in the box:")
	myBox.DisplayToys()
}