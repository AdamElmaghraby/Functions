package main

import "fmt"

func main() {
	intArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Example 1: Looping through an array of int")
	for i := 0; i < len(intArray); i++ {
		fmt.Printf("%d\n", intArray[i])
	}
	fmt.Println("\n")

	stringArray := [3]string{"Adam", "Chris", "Alan"}
	fmt.Println("Example 2: Looping through an array of strings")
	for _, value := range stringArray {
		fmt.Printf("%s\n", value)
	}
	fmt.Println("\n")

	floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println("Example 3: Looping through a slice of floats")
	for _, value := range floatSlice {
		fmt.Printf("%.2f\n", value)
	}
	fmt.Println("\n")

	stringIntMap := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("Example 4 : Looping through a map of strings with ints")
	for key, value := range stringIntMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	valueInt := stringIntMap["two"]

	fmt.Println("\n")

	fmt.Printf("The value at key 'two' is: %d", valueInt)

	fmt.Println("\n")

	intStringMap := map[int]string{1: "one", 2: "two", 3: "three"}
	fmt.Println("Example 4 : Looping through a map of strings with ints")
	for key, value := range intStringMap {
		fmt.Printf("Key: %d, Value: %s\n", key, value)
	}

	valueString := intStringMap[2]

	fmt.Println("\n")

	fmt.Printf("The value at key '2' is: %s", valueString)

}
