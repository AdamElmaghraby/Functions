package functions

import (
	"math/rand"
)

func Add(x, y int) int {
	return x + y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Distance (x, y int) int {
	difference := x - y
	if difference < 0 {
		return -difference
	}
	return difference
}

func RandNumber(n, max int) []int {
	var numbers []int

	for i := 0; i < n; i++ {
		numbers = append(numbers, rand.Intn(max))

	}
	return numbers
}

func Shuffle(s []int) []int {

	for i := 0; i < len(s); i++ {
		x, y := rand.Intn(len(s)), rand.Intn(len(s))

		s[x], s[y] = s[y], s[x]
	}
	return s
}

func ReverseArr(r []int) {
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}

func ReverseString() {

}

func IsEven() {

}

func IsOdd() {

}

func GenOdd(a, b int) {

}

func GenEven(a, b int) {

}
