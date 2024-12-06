package main

import (
	"fmt"
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
	rules, updates := rulesAndUpdates(lines)

	t := 0
	for _, v := range updates {
		if updateIsValid(v, rules) {
			t += v[len(v)/2]
		}
	}

	return t
}

func part2(filename string) int {
	fmt.Printf("Reading file %s\n", filename)
	lines := files.ReadLines(filename)
	rules, updates := rulesAndUpdates(lines)

	t := 0
	for _, v := range updates {
		if !updateIsValid(v, rules) {
			pages := orderPages(v, rules)
			t += pages[len(pages)/2]
		}
	}

	return t
}

func orderPages(pageUpdates []int, rules map[int][]int) []int {
	fmt.Printf("orderPages: %v\n", pageUpdates)
	printed := make([]int, len(pageUpdates))
	for i := 0; i < len(pageUpdates); {
		curr := pageUpdates[i]
		print := true
		for printIndex, v := range printed {
			if ints.Includes(rules[curr], v) {
				// swap
				// fmt.Printf("Rule: %v, Printed: %v, Curr: %d\n", rules[curr], printed, curr)
				temp := pageUpdates[i]
				pageUpdates[i] = pageUpdates[printIndex]
				pageUpdates[printIndex] = temp
				fmt.Printf("swapPages: %v\n", pageUpdates)

				printed = make([]int, len(pageUpdates))
				i = 0
				print = false
				break
			}
		}
		if print {
			printed[i] = curr
			fmt.Printf("Printed: %v\n", printed)
			i++
		}
	}
	return pageUpdates
}

func updateIsValid(pageUpdates []int, rules map[int][]int) bool {
	printed := []int{}
	for _, curr := range pageUpdates {
		for _, v := range printed {
			if ints.Includes(rules[curr], v) {
				return false
			}
		}
		printed = append(printed, curr)
	}

	fmt.Printf("Update is valid: %v\n", pageUpdates)
	return true
}

func rulesAndUpdates(lines []string) (map[int][]int, [][]int) {
	var rulesString, updatesString []string

	for i, v := range lines {
		if strings.Contains(v, ",") {
			rulesString = lines[:i]
			updatesString = lines[i:]
			break
		}
	}
	fmt.Printf("RulesString: %s\n", rulesString)
	fmt.Printf("UpdatesString: %s\n", updatesString)

	updates := make([][]int, len(updatesString))

	rules := make(map[int][]int)
	for _, v := range rulesString {
		rule := ints.SliceFromStringDelim(v, '|')
		if _, ok := rules[rule[0]]; !ok {
			rules[rule[0]] = []int{}
		}
		rules[rule[0]] = append(rules[rule[0]], rule[1])
	}
	fmt.Printf("Rules: %v\n", rules)
	for i, v := range updatesString {
		updates[i] = ints.SliceFromStringDelim(v, ',')
	}
	return rules, updates
}
