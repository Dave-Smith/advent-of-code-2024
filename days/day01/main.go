package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Dave-Smith/advent-of-code-2024/internal/files"
	"github.com/Dave-Smith/advent-of-code-2024/internal/ints"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1("./file.txt"))
	fmt.Printf("Part 2: %d\n", part2("./file.txt"))
}

func part1(filename string) int {
	fmt.Printf("Reading file %s\n", filename)
	lines := files.ReadLines(filename)
	left, right := ListsFromFile(lines)

	slices.Sort(left)
	slices.Sort(right)

	t := 0
	for i := 0; i < len(left); i++ {
		t += ints.Abs(left[i] - right[i])
	}

	return t
}

func part2(filename string) int {
	fmt.Printf("Reading file %s\n", filename)
	lines := files.ReadLines(filename)
	left, right := ListsFromFile(lines)

	leftMap, rightMap := make(map[int]int), make(map[int]int)
	for i := 0; i < len(left); i++ {
		leftMap[left[i]] += 1
		rightMap[right[i]] += 1
	}

	similarityScore := 0
	for k, v := range leftMap {
		similarityScore += k * v * rightMap[k]
	}

	return similarityScore
}

func ListsFromFile(lines []string) ([]int, []int) {
	left, right := make([]int, len(lines)), make([]int, len(lines))

	for i := 0; i < len(lines); i++ {
		ids := strings.Fields(lines[i])
		left[i] = ints.FromString(ids[0])
		right[i] = ints.FromString(ids[1])
	}

	return left, right
}
