package main

import (
	"fmt"

	"github.com/Dave-Smith/advent-of-code-2024/internal/files"
)

type Point struct {
	X, Y int
}

type Board struct {
	Height, Width int
	Lines         []string
}

type Walker struct {
	Dirs     []func(Point) Point
	DirIndex int
}

var empty struct{}

func main() {
	fmt.Printf("Part 1: %d\n", part1("./file.txt"))
	fmt.Printf("Part 2: %d\n", part2("./file.txt"))
}

func part1(filename string) int {
	fmt.Printf("Reading file %s\n", filename)

	lines := files.ReadLines(filename)
	width := len(lines[0])
	height := len(lines)
	obstructions := FindObstructions(lines)
	visited := make(map[Point]struct{})

	curr := FindStartingPoint(lines)
	dirs := []func(Point) Point{up, right, down, left}
	dirIndex := 0
	forward := dirs[dirIndex]
	for curr.X > 0 && curr.X <= width && curr.Y > 0 && curr.Y <= height {
		visited[curr] = empty
		if obstructions[forward(curr)] {
			fmt.Printf("Turning right: %v\n", curr)
			dirIndex++
			forward = dirs[dirIndex%4]
		}
		curr = forward(curr)
		fmt.Printf("Moved forward to %v\n", curr)
	}

	return len(visited)
}

func up(p Point) Point {
	p.Y--
	return p
}

func right(p Point) Point {
	p.X++
	return p
}

func down(p Point) Point {
	p.Y++
	return p
}
func left(p Point) Point {
	p.X--
	return p
}

func FindStartingPoint(lines []string) Point {
	for row, _ := range lines {
		for cell, _ := range lines[row] {
			if lines[row][cell] == '^' {
				return Point{cell + 1, row + 1}
			}
		}
	}
	return Point{}
}

func FindObstructions(lines []string) map[Point]bool {
	obstructions := make(map[Point]bool)
	for row, _ := range lines {
		for cell, _ := range lines[row] {
			if lines[row][cell] == '#' {
				obstructions[Point{cell + 1, row + 1}] = true
			}
		}
	}
	fmt.Printf("Obstructions: %v\n", obstructions)
	return obstructions
}

func part2(filename string) int {
	fmt.Printf("Reading file %s\n", filename)
	lines := files.ReadLines(filename)
	board := Board{Height: len(lines), Width: len(lines[0]), Lines: lines}
	obstructions := FindObstructions(lines)
	visited := make(map[Point]struct{})
	newBlocks := []Point{}

	curr := FindStartingPoint(lines)
	w := Walker{Dirs: []func(Point) Point{up, down, left, right}, DirIndex: 0}
	dirs := []func(Point) Point{up, right, down, left}
	dirIndex := 0
	forward := dirs[dirIndex]

	for inBounds(board, curr) {
		if obstructions[forward(curr)] {
			fmt.Printf("Turning right: %v\n", curr)
			dirIndex = (dirIndex + 1) % 4
			forward = dirs[dirIndex]
		}
		if loopTest(curr, board, w) {
			newBlocks = append(newBlocks, forward(curr))
		}
		// if runsIntoBlock(curr, board, dirs[(dirIndex+1)%4]) && runsIntoBlock(curr, board, dirs[(dirIndex+2)%4]) {
		// 	newBlocks = append(newBlocks, forward(curr))
		// }
		visited[curr] = empty
		curr = forward(curr)
		fmt.Printf("Moved forward to %v\n", curr)
	}

	fmt.Printf("NewBlocks:%v\n", newBlocks)
	return len(newBlocks)
}

func runsIntoBlock(curr Point, board Board, next func(Point) Point) bool {
	for inBounds(board, next(curr)) {
		if isObs(next(curr), board) {
			fmt.Printf("Runs into block at %v\n", next(curr))
			return true
		}
		curr = next(curr)
	}
	return false
}

func loopTest(curr Point, board Board, w Walker) bool {
	start := curr
	for inBounds(board, curr) {
		if isObs(curr, board) {
			//fmt.Printf("Turning right: %v\n", curr)
			w.TurnRight()
		}
		curr = w.Next(curr)
		if curr == start {
			return true
		}

	}
	return false
}

func inBounds(board Board, curr Point) bool {
	return curr.X > 0 && curr.X <= board.Width && curr.Y > 0 && curr.Y <= board.Height
}

func isObs(curr Point, board Board) bool {
	return board.Lines[curr.Y-1][curr.X-1] == '#'
}

func (w *Walker) TurnRight() {
	w.DirIndex = w.DirIndex % 4
}

func (w *Walker) Next(curr Point) Point {
	return w.Dirs[w.DirIndex](curr)
}
