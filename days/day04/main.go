package main

import (
	"fmt"

	"github.com/Dave-Smith/advent-of-code-2024/internal/files"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1("./file.txt"))
	fmt.Printf("Part 2: %d\n", part2("./file.txt"))
}

func part1(filename string) int {
	fmt.Printf("Reading file %s\n", filename)
	lines := files.ReadLines(filename)
	t := 0
	for x := 0; x < len(lines[0]); x++ {
		for y := 0; y < len(lines); y++ {
			if lines[y][x] != 'X' {
				continue
			}
			if left(lines, []rune{'M', 'A', 'S'}, x, y) {
				t++
			}
			if right(lines, []rune{'M', 'A', 'S'}, x, y) {
				t++
			}
			if down(lines, []rune{'M', 'A', 'S'}, x, y) {
				t++
			}
			if up(lines, []rune{'M', 'A', 'S'}, x, y) {
				t++
			}
			if diagUpRight(lines, []rune{'M', 'A', 'S'}, x, y) {
				t++
			}
			if diagUpLeft(lines, []rune{'M', 'A', 'S'}, x, y) {
				t++
			}
			if diagDownRight(lines, []rune{'M', 'A', 'S'}, x, y) {
				t++
			}
			if diagDownLeft(lines, []rune{'M', 'A', 'S'}, x, y) {
				t++
			}
		}
	}

	return t
}

func left(lines []string, letters []rune, x, y int) bool {
	if len(letters) == 0 {
		return true
	}
	if x-1 < 0 {
		return false
	}
	if rune(lines[y][x-1]) != letters[0] {
		return false
	}
	return left(lines, letters[1:], x-1, y)
}

func right(lines []string, letters []rune, x, y int) bool {
	if len(letters) == 0 {
		return true
	}
	if x+1 > len(lines[0])-1 {
		return false
	}
	if rune(lines[y][x+1]) != letters[0] {
		return false
	}
	return right(lines, letters[1:], x+1, y)
}

func up(lines []string, letters []rune, x, y int) bool {
	if len(letters) == 0 {
		return true
	}
	if y-1 < 0 {
		return false
	}
	if rune(lines[y-1][x]) != letters[0] {
		return false
	}
	return up(lines, letters[1:], x, y-1)
}

func down(lines []string, letters []rune, x, y int) bool {
	if len(letters) == 0 {
		return true
	}
	if y+1 > len(lines)-1 {
		return false
	}
	if rune(lines[y+1][x]) != letters[0] {
		return false
	}
	return down(lines, letters[1:], x, y+1)
}

func diagUpRight(lines []string, letters []rune, x, y int) bool {
	if len(letters) == 0 {
		return true
	}
	if y-1 < 0 || x+1 > len(lines[y])-1 {
		return false
	}
	if rune(lines[y-1][x+1]) != letters[0] {
		return false
	}
	return diagUpRight(lines, letters[1:], x+1, y-1)
}

func diagUpLeft(lines []string, letters []rune, x, y int) bool {
	if len(letters) == 0 {
		return true
	}
	if y-1 < 0 || x-1 < 0 {
		return false
	}
	if rune(lines[y-1][x-1]) != letters[0] {
		return false
	}
	return diagUpLeft(lines, letters[1:], x-1, y-1)
}

func diagDownRight(lines []string, letters []rune, x, y int) bool {
	if len(letters) == 0 {
		return true
	}
	if y+1 > len(lines)-1 || x+1 > len(lines[y])-1 {
		return false
	}
	if rune(lines[y+1][x+1]) != letters[0] {
		return false
	}
	return diagDownRight(lines, letters[1:], x+1, y+1)
}

func diagDownLeft(lines []string, letters []rune, x, y int) bool {
	if len(letters) == 0 {
		return true
	}
	if y+1 > len(lines)-1 || x-1 < 0 {
		return false
	}
	if rune(lines[y+1][x-1]) != letters[0] {
		return false
	}
	return diagDownLeft(lines, letters[1:], x-1, y+1)
}

func part2(filename string) int {
	lines := files.ReadLines(filename)
	t := 0
	for x := 1; x < len(lines[0])-1; x++ {
		for y := 1; y < len(lines)-1; y++ {
			if lines[y][x] != 'A' {
				continue
			}
			if xmas(lines, x, y) {
				t++
			}
		}
	}
	return t
}

func xmas(lines []string, x, y int) bool {
	line1 := string([]byte{lines[y-1][x-1], lines[y][x], lines[y+1][x+1]})
	line2 := string([]byte{lines[y+1][x-1], lines[y][x], lines[y-1][x+1]})

	if line1 != "MAS" && rev(line1) != "MAS" {
		return false
	}
	if line2 != "MAS" && rev(line2) != "MAS" {
		return false
	}
	return line1 == line2 || rev(line1) == rev(line2) || rev(line1) == line2 || line1 == rev(line2)
}

func rev(str string) string {
	result := ""
	for _, v := range str {
		result = string(v) + result
	}
	return result
}
