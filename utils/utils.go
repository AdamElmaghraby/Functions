package utils

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

func AbsValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Distance(x, y int) int {
	return AbsValue(x - y)
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

func ReverseArr(r []int) []int {
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}

func ReverseString(s string) string {
	rvs := []rune(s)
	for i, j := 0, len(rvs)-1; i < j; i, j = i+1, j-1 {
		rvs[i], rvs[j] = rvs[j], rvs[i]
	}
	return string(rvs)
}

func IsEven(x int) bool {
	return x%2 == 0
}

func IsOdd(x int) bool {
	return x%2 != 0
}

func GenOdd(min, max int) []int {
	var list []int

	if min%2 == 0 {
		min++
	}

	for i := min; i <= max; i += 2 {
		list = append(list, i)
	}

	return list
}

func GenEven(min, max int) []int {
	var list []int

	if min == 0 {
		min = 2
	} else if min%2 != 0 {
		min++
	}

	for i := min; i <= max; i += 2 {
		list = append(list, i)
	}

	return list
}

func ArrEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func ArrEqualStr(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func HasElement(arr []int, element int) bool {
	for _, i := range arr {
		if i == element {
			return true
		}
	}
	return false
}

func Deduplicate(arr []string) []string {

	m := make(map[string]bool)

	var result []string

	for _, value := range arr {
		if !m[value] {
			m[value] = true
			result = append(result, value)
		}
	}

	return result
}
