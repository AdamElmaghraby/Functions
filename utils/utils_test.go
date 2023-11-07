package utils

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type AddTest struct {
		Expect int
		Input1 int
		Input2 int
	}

	table := []AddTest{
		{
			Input1: 2,
			Input2: 3,
			Expect: 5,
		},
		{
			Input1: -2,
			Input2: 3,
			Expect: 1,
		},
		{
			Input1: 2,
			Input2: -3,
			Expect: -1,
		},
		{
			Input1: -2,
			Input2: -3,
			Expect: -5,
		},
		{
			Input1: 0,
			Input2: 0,
			Expect: 0,
		},
		{
			Input1: 1,
			Input2: 2,
			Expect: 3,
		},
	}

	for _, v := range table {
		got := Add(v.Input1, v.Input2)

		if got != v.Expect {
			t.Errorf("for inputs [%d, %d] expected %d; got %d", v.Input1, v.Input2, v.Expect, got)
		}
	}
}

func TestMax(t *testing.T) {
	type MaxTest struct {
		Expect int
		Input1 int
		Input2 int
	}

	table := []MaxTest{
		{
			Input1: 2,
			Input2: 3,
			Expect: 3,
		},
		{
			Input1: -2,
			Input2: 3,
			Expect: 3,
		},
		{
			Input1: 2,
			Input2: -3,
			Expect: 2,
		},
		{
			Input1: -2,
			Input2: -3,
			Expect: -2,
		},
		{
			Input1: -1999,
			Input2: 1000000,
			Expect: 1000000,
		},
		{
			Input1: -20000,
			Input2: -19000,
			Expect: -19000,
		},
	}

	for _, v := range table {
		got := Max(v.Input1, v.Input2)

		if got != v.Expect {
			t.Errorf("for inputs [%d, %d] expected %d; got %d", v.Input1, v.Input2, v.Expect, got)
		}
	}
}

func TestMin(t *testing.T) {
	type MinTest struct {
		Expect int
		Input1 int
		Input2 int
	}

	table := []MinTest{
		{
			Input1: 2,
			Input2: 3,
			Expect: 2,
		},
		{
			Input1: -2,
			Input2: 3,
			Expect: -2,
		},
		{
			Input1: 2,
			Input2: -3,
			Expect: -3,
		},
		{
			Input1: -2,
			Input2: -3,
			Expect: -3,
		},
		{
			Input1: -1999,
			Input2: 1000000,
			Expect: -1999,
		},
		{
			Input1: -20000,
			Input2: -19000,
			Expect: -20000,
		},
	}

	for _, v := range table {
		got := Min(v.Input1, v.Input2)

		if got != v.Expect {
			t.Errorf("for inputs [%d, %d] expected %d; got %d", v.Input1, v.Input2, v.Expect, got)
		}
	}
}

func TestAbsValue(t *testing.T) {
	type AbsTest struct {
		Expect int
		Input1 int
	}

	table := []AbsTest{
		{
			Input1: 1,
			Expect: 1,
		},
		{
			Input1: -2,
			Expect: 2,
		},
		{
			Input1: 0,
			Expect: 0,
		},
		{
			Input1: 100000000,
			Expect: 100000000,
		},
		{
			Input1: -100000000,
			Expect: 100000000,
		},
	}

	for _, v := range table {
		got := AbsValue(v.Input1)

		if got != v.Expect {
			t.Errorf("for inputs [%d] expected %d; got %d", v.Input1, v.Expect, got)
		}
	}
}

func TestDistance(t *testing.T) {
	type DisTest struct {
		Expect int
		Input1 int
		Input2 int
	}

	table := []DisTest{
		{
			Input1: 0,
			Input2: 0,
			Expect: 0,
		},
		{
			Input1: -2,
			Input2: -2,
			Expect: 0,
		},
		{
			Input1: 1,
			Input2: 1,
			Expect: 0,
		},
		{
			Input1: 2,
			Input2: 3,
			Expect: 1,
		},
		{
			Input1: 1,
			Input2: 5,
			Expect: 4,
		},
		{
			Input1: -2,
			Input2: 1,
			Expect: 3,
		},
	}

	for _, v := range table {
		got := Distance(v.Input1, v.Input2)

		if got != v.Expect {
			t.Errorf("for inputs [%d, %d] expected %d; got %d", v.Input1, v.Input2, v.Expect, got)
		}
	}
}

func TestRandNumber(t *testing.T) {
	n := 6
	max := 1000

	nums := RandNumber(n, max)

	if len(nums) != n {
		t.Errorf("length of result is %d instead of %d", len(nums), n)
	}

	for _, v := range nums {
		if v < 0 || v > max {
			t.Errorf("result is outside valid range; got %d", v)
		}
	}
}

func TestShuffle(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	output := Shuffle(input)

	if len(input) != len(output) {
		t.Errorf("Length of input and output slices do not match")
	}

	for _, element := range input {
		if !hasElement(output, element) {
			t.Errorf("Element %d missing in the output slice", element)
		}
	}

}

func TestReverseArr(t *testing.T) {
	type RvsArrTest struct {
		Expect []int
		Input1 []int
	}

	table := []RvsArrTest{
		{
			Input1: []int{1, 2, 3, 4, 5},
			Expect: []int{5, 4, 3, 2, 1},
		},
		{
			Input1: []int{5, 5, 5, 5, 5},
			Expect: []int{5, 5, 5, 5, 5},
		},
		{
			Input1: []int{-1, -2, -3, -4, -5},
			Expect: []int{-5, -4, -3, -2, -1},
		},
		{
			Input1: []int{1, 0, 3, 4, 5},
			Expect: []int{5, 4, 3, 0, 1},
		},
		{
			Input1: []int{1, 0, -3, 4, 5},
			Expect: []int{5, 4, -3, 0, 1},
		},
	}

	for _, test := range table {
		reversedSlice := ReverseArr(test.Input1)

		if !arrEqual(reversedSlice, test.Expect) {
			t.Errorf("For input %v, expected %v; got %v", test.Input1, test.Expect, reversedSlice)
		}
	}
}
