package files

import (
	"slices"
	"testing"
)

func TestReadLinesEmptyFile(t *testing.T) {
	t.Run("Read blank file from current directory and return empty slice", func(t *testing.T) {
		actual := ReadLines("./input-empty.txt")
		if len(actual) != 0 {
			t.Fatalf(`ReadLines("./input-empty.txt"): got: %d, want: %d`, len(actual), 0)
		}
	})
}

func TestReadLines(t *testing.T) {
	t.Run("Read file from current directory and return all non empty lines", func(t *testing.T) {
		want := []string{"hello", "advent", "of", "code"}
		actual := ReadLines("./input-test.txt")
		if !slices.Equal(want, actual) {
			t.Fatalf(`ReadLines("./input.txt"): got: %v, want: %v`, actual, want)
		}
	})
}
