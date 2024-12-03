package slice

import (
	"slices"
	"testing"
)

func TestPopFromSlice(t *testing.T) {
	t.Run("Returns a new slice without the item at index", func(t *testing.T) {
		want := []int{0, 1, 2, 3, 4, 6, 7, 8, 9}
		actual := PopFromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 5)
		if slices.Compare(want, actual) != 0 {
			t.Fatalf(`PopFromSlice(...): want: %v, got: %v`, want, actual)
		}
	})
}
