package main

import (
	"fmt"

	"github.com/Dave-Smith/advent-of-code-2024/internal/files"
	"github.com/Dave-Smith/advent-of-code-2024/internal/ints"
)

var empty struct{}

type Point struct {
	X, Y int
}

type Board struct {
	Height, Width int
	Points        map[Point]int
}

var moves = []func(Point) Point{up, right, down, left}

func main() {
	fmt.Printf("Part 1: %d\n", part1("./file.txt"))
	fmt.Printf("Part 2: %d\n", part2("./file.txt"))
}

func part1(filename string) int {
	fmt.Printf("Reading file %s\n", filename)

	lines := files.ReadLines(filename)
	points := make(map[Point]int)

	for y, line := range lines {
		for x, cell := range line {
			points[Point{X: x + 1, Y: y + 1}] = ints.FromString(string(cell))
		}
	}
	board := Board{Height: len(lines), Width: len(lines[0]), Points: points}
	fmt.Printf("Board: %v\n", board)

	t := 0
	for k, v := range board.Points {
		if v == 0 {
			t += TrailheadPeaksScore(k, board)
		}
	}
	return t
}

func TrailheadPeaksScore(th Point, board Board) int {
	return travel(th, th, board, 0, make(map[Point]struct{}))
}
func TrailheadTrailScore(th Point, board Board) int {
	return travelTrails(th, th, board, 0)
}

func travel(th, curr Point, board Board, score int, peaks map[Point]struct{}) int {
	height := board.Points[curr]
	if _, ok := peaks[curr]; !ok && height == 9 {
		score++
		peaks[curr] = empty
		return score
	}
	for _, m := range moves {
		next := m(curr)
		if _, ok := board.Points[next]; ok {
			if board.Points[next] == height+1 {
				score = travel(th, next, board, score, peaks)
			}
		}
	}
	return score
}

func travelTrails(th, curr Point, board Board, score int) int {
	height := board.Points[curr]
	if height == 9 {
		score++
		return score
	}
	for _, m := range moves {
		next := m(curr)
		if _, ok := board.Points[next]; ok {
			if board.Points[next] == height+1 {
				score = travelTrails(th, next, board, score)
			}
		}
	}
	return score
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

func part2(filename string) int {
	lines := files.ReadLines(filename)
	points := make(map[Point]int)

	for y, line := range lines {
		for x, cell := range line {
			points[Point{X: x + 1, Y: y + 1}] = ints.FromString(string(cell))
		}
	}
	board := Board{Height: len(lines), Width: len(lines[0]), Points: points}
	fmt.Printf("Board: %v\n", board)

	t := 0
	for k, v := range board.Points {
		if v == 0 {
			t += TrailheadTrailScore(k, board)
		}
	}
	return t
}
