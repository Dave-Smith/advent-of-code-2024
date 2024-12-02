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
	reports := ReadReports(lines)

	t := 0
	for ri, r := range reports {
		if IsReportSafe(ri, r) {
			t++
		}
	}
	return t
}

func IsReportSafe(reportId int, r []int) bool {
	inc, dec := false, false
	for j := 1; j < len(r); j++ {
		for i := j - 1; i < j; i++ {
			diff := r[i] - r[j]
			if ints.Abs(diff) > 3 || diff == 0 {
				return false
			}
			if diff > 0 {
				dec = true
			}
			if diff < 0 {
				inc = true
			}
		}
		if inc && dec {
			return false
		}
	}
	fmt.Printf("ReportId %d is safe: %v\n", reportId, r)
	return true
}

func ReadReports(lines []string) [][]int {
	reports := make([][]int, len(lines))
	for i := 0; i < len(lines); i++ {
		reportVals := strings.Fields(lines[i])
		report := make([]int, len(reportVals))
		for ix, v := range reportVals {
			report[ix] = ints.FromString(v)
		}
		reports[i] = report
	}

	return reports
}

func IsReportSafeWithDampener(r []int) bool {
	for i := 0; i < len(r); i++ {
		rd := make([]int, 0)
		for l := 0; l < len(r); l++ {
			if i == l {
				continue
			}
			rd = append(rd, r[l])
		}
		if IsReportSafe(0, rd) {
			return true
		}
	}
	return false
}

func part2(filename string) int {
	fmt.Printf("Reading file %s\n", filename)
	lines := files.ReadLines(filename)
	reports := ReadReports(lines)

	t := 0
	for ri, r := range reports {
		if IsReportSafe(ri, r) {
			t++
			continue
		}
		if IsReportSafeWithDampener(r) {
			t++
		}
	}
	return t
}
