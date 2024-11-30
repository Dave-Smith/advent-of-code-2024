package ints

import "testing"

func TestSum(t *testing.T) {
	t.Run("Sum all ints in slice", func(t *testing.T) {
		want := 10
		actual := Sum([]int{1, 2, 3, 4})
		if want != actual {
			t.Fatalf("Sum([]int{1, 2, 3, 4}): want: %d, got: %d", want, actual)
		}
	})
}

func TestSumEmpty(t *testing.T) {
	t.Run("Sum all ints in slice", func(t *testing.T) {
		want := 0
		actual := Sum([]int{})
		if want != actual {
			t.Fatalf("Sum([]int{}): want: %d, got: %d", want, actual)
		}
	})
}

func TestFromString(t *testing.T) {
	t.Run("Returns an int from a string", func(t *testing.T) {
		want := 10
		actual := FromString("10")
		if want != actual {
			t.Fatalf(`FromString("10"): want: %d, got: %d`, want, actual)
		}
	})
}

func TestFromStringTrim(t *testing.T) {
	t.Run("Returns an int from a string", func(t *testing.T) {
		want := 10
		actual := FromString(" 10 ")
		if want != actual {
			t.Fatalf(`FromString("10"): want: %d, got: %d`, want, actual)
		}
	})
}