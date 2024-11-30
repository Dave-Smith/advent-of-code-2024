package main

import "testing"

func TestPart1(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		want := 0
		actual := part1("test-input.txt")
		if want != actual {
			t.Fatalf(`part1("test-input.txt") = %d, want %d, error`, actual, want)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("Part 2", func(t *testing.T) {
		want := 0
		actual := part1("test-input.txt")
		if want != actual {
			t.Fatalf(`part2("test-input.txt") = %d, want %d, error`, actual, want)
		}
	})
}
