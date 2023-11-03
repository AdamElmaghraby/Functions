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

/*func TestAdd(t *testing.T) {
	t.Run("TwoPositive", func(t *testing.T) {
		if Add(1, 1) != 2 {
			t.Error("Adding 1 + 1 should equal 2")
		}
	})

    t.Run("TwoNegative", func(t *testing.T) {
		if Add(-1, -1) != -2 {
			t.Error("Adding -1 + -1 should equal -2")
		}
	})

    t.Run("OneOne", func(t *testing.T) {
		if Add(1, -1) != 0 {
			t.Error("Adding 1 + -1 should equal 0")
		}
	})

    t.Run("OneOne", func(t *testing.T) {
		if Add(1, -1) != 0 {
			t.Error("Adding 1 + -1 should equal 0")
		}
	})





}*/
