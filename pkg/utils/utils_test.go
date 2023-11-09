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
		if !HasElement(output, element) {
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

	for _, v := range table {
		got := ReverseArr(v.Input1)

		if !ArrEqual(got, v.Expect) {
			t.Errorf("For input %v, expected %v; got %v", v.Input1, v.Expect, got)
		}
	}
}

func TestIsEven(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{2, true},
		{0, true},
		{1, false},
		{3, false},
		{-2, true},
		{-1, false},
	}

	for _, tc := range testCases {
		result := IsEven(tc.input)
		if result != tc.expected {
			t.Errorf("For input %d, expected IsEven to be %v, but got %v", tc.input, tc.expected, result)
		}
	}
}

func TestIsOdd(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{2, false},
		{0, false},
		{1, true},
		{3, true},
		{-2, false},
		{-1, true},
	}

	for _, tc := range testCases {
		result := IsOdd(tc.input)
		if result != tc.expected {
			t.Errorf("For input %d, expected IsOdd to be %v, but got %v", tc.input, tc.expected, result)
		}
	}
}

func TestGenOdd(t *testing.T) {
	type GenOddTest struct {
		min      int
		max      int
		expected []int
	}
	table := []GenOddTest{
		{1, 10, []int{1, 3, 5, 7, 9}},
		{2, 10, []int{3, 5, 7, 9}},
		{10, 20, []int{11, 13, 15, 17, 19}},
		{0, 1, []int{1}},
	}

	for _, v := range table {
		result := GenOdd(v.min, v.max)
		if !ArrEqual(result, v.expected) {
			t.Errorf("For range [%d, %d], expected %v, but got %v", v.min, v.max, v.expected, result)
		}
	}
}

func TestGenEven(t *testing.T) {
	type GenEvenTest struct {
		min      int
		max      int
		expected []int
	}
	table := []GenEvenTest{
		{1, 10, []int{2, 4, 6, 8, 10}},
		{2, 10, []int{2, 4, 6, 8, 10}},
		{10, 20, []int{10, 12, 14, 16, 18, 20}},
		{0, 1, []int{}},
	}

	for _, v := range table {
		result := GenEven(v.min, v.max)
		if !ArrEqual(result, v.expected) {
			t.Errorf("For range [%d, %d], expected %v, but got %v", v.min, v.max, v.expected, result)
		}
	}
}

func TestArrEqual(t *testing.T) {
	type ArrEqualTest struct {
		input1   []int
		input2   []int
		expected bool
	}
	table := []ArrEqualTest{
		{[]int{2, 4, 6, 8, 10}, []int{2, 4, 6, 8, 10}, true},
		{[]int{1, 2, 3}, []int{1, 2, 3, 4}, false},
		{[]int{1, 2, 3}, []int{3, 2, 1}, false},
		{[]int{}, []int{}, true},
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
	}

	for _, v := range table {
		result := ArrEqual(v.input1, v.input2)
		if result != v.expected {
			t.Errorf("For slices [%d, %d], expected %v, but got %v", v.input1, v.input2, v.expected, result)
		}
	}
}

func TestHasElement(t *testing.T) {
	type HasElementTest struct {
		input1   []int
		input2   int
		expected bool
	}
	table := []HasElementTest{
		{[]int{2, 4, 6, 8, 10}, 2, true},
		{[]int{1, 3, 5, 7, 9}, 4, false},
		{[]int{1, 2, 3, 4, 5}, 6, false},
		{[]int{}, 42, false},
	}

	for _, v := range table {
		result := HasElement(v.input1, v.input2)
		if result != v.expected {
			t.Errorf("For slice %v and element %d, expected %v, but got %v", v.input1, v.input2, v.expected, result)
		}
	}
}

func TestDeduplicate(t *testing.T) {
	type DeduplicateTest struct {
		input1   []string
		expected []string
	}
	table := []DeduplicateTest{
		{[]string{"2", "4", "6", "8", "10"}, []string{"2", "4", "6", "8", "10"}},
		{[]string{"2", "2", "4", "4", "6", "6"}, []string{"2", "4", "6"}},
		{[]string{"a", "b", "c", "a", "b"}, []string{"a", "b", "c"}},
		{[]string{}, []string{}},
		{[]string{"apple", "banana", "cherry", "date", "fig", "apple"}, []string{"apple", "banana", "cherry", "date", "fig"}},
		{[]string{"x", "x", "y", "x", "z", "y", "z", "y"}, []string{"x", "y", "z"}},
	}

	for _, v := range table {
		result := Deduplicate(v.input1)
		if !ArrEqualStr(result, v.expected) {
			t.Errorf("For slice %v, expected %v, but got %v", v.input1, v.expected, result)
		}
	}
}
