package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name           string `json:"name"`
	Age            int    `json:"age"`
	Phone          string `json:"phone"`
	FavoritePhrase string `json:"custom1"`
}

func main() {
	data, err := os.ReadFile("file.json")

	if err != nil {
		log.Println(err)
	}

	var p Person

	err = json.Unmarshal(data, &p)

	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Printf("Hi, my name is %s! I am %d years old and my phone number is %s. I love saying %s!\n", p.Name, p.Age, p.Phone, p.FavoritePhrase)
}
