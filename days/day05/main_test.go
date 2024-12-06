package main

import "testing"

func TestPart1(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		want := 143
		actual := part1("test-input.txt")
		if want != actual {
			t.Fatalf(`part1("test-input.txt"): got: %d, want: %d`, actual, want)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("Part 2", func(t *testing.T) {
		want := 123
		actual := part2("test-input.txt")
		if want != actual {
			t.Fatalf(`part2("test-input.txt"): got: %d, want: %d, error`, actual, want)
		}
	})
}
