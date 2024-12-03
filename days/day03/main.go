package main

import (
	"fmt"
	"regexp"

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
	mulExp, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	numExp, _ := regexp.Compile(`\d{1,3}`)

	t := 0
	for _, line := range lines {
		muls := mulExp.FindAllString(line, -1)
		for _, mul := range muls {
			fmt.Printf("mul: %s\n", mul)
			nums := numExp.FindAllString(mul, -1)
			t += ints.FromString(nums[0]) * ints.FromString(nums[1])
		}
	}
	return t
}

func part2(filename string) int {
	fmt.Printf("Reading file %s\n", filename)
	lines := files.ReadLines(filename)
	mulExp, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	numExp, _ := regexp.Compile(`\d{1,3}`)
	doExp, _ := regexp.Compile(`do\(\)`)
	dontExp, _ := regexp.Compile(`don\'t\(\)`)

	t := 0
	enabled := true
	for _, line := range lines {
		for len(line) > 0 {
			if enabled {
				mulIndex := mulExp.FindStringIndex(line)
				dontIndex := dontExp.FindStringIndex(line)

				// neither dont() or mul() appear in the rest of the line
				if dontIndex == nil && mulIndex == nil {
					break
				}

				// dont() appears before next mul(). Advance down the line
				if dontIndex != nil && dontIndex[0] < mulIndex[0] {
					line = line[dontIndex[1]:]
					enabled = false
					continue
				}

				// mul() appears before next dont(). Add to total and move down the line
				mul := line[mulIndex[0]:mulIndex[1]]
				nums := numExp.FindAllString(mul, -1)

				t += ints.FromString(nums[0]) * ints.FromString(nums[1])
				line = line[mulIndex[1]:]
			}

			if !enabled {
				doIndex := doExp.FindStringIndex(line)

				// do() is not on the current line. Move to next line
				if doIndex == nil {
					break
				}

				// do() is on the current line. Enable and move down the line
				line = line[doIndex[1]:]
				enabled = true
			}
		}
	}

	return t
}
