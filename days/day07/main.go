package main

import (
	"fmt"
	"strings"

	"github.com/Dave-Smith/advent-of-code-2024/internal/files"
	"github.com/Dave-Smith/advent-of-code-2024/internal/ints"
)

type Calibration struct {
	TestValue int64
	Values    []int
}

func main() {
	fmt.Printf("Part 1: %d\n", part1("./file.txt"))
	fmt.Printf("Part 2: %d\n", part2("./file.txt"))
}

func part1(filename string) int64 {
	fmt.Printf("Reading file %s\n", filename)
	cals := ParseCalibrations(filename)

	t := int64(0)

	ops := []func(int64, int64) int64{mult, add}
	for i := 0; i < len(cals); i++ {
		if IsCalCorrectPartTwo(cals[i].TestValue, 0, cals[i].Values, ops) {
			fmt.Printf("Correct: %d: %v\n", cals[i].TestValue, cals[i].Values)
			t += cals[i].TestValue
		}
	}
	// 2299996598890
	return t
}

func ParseCalibrations(filename string) []Calibration {
	lines := files.ReadLines(filename)
	cals := make([]Calibration, len(lines))
	for i := 0; i < len(lines); i++ {
		colonIndex := strings.IndexAny(lines[i], ":")
		test := ints.FromSubstring(lines[i], 0, colonIndex)
		vals := ints.SliceFromString(string(lines[i][colonIndex:]))
		cals[i] = Calibration{TestValue: test, Values: vals}
	}
	return cals
}

func part2(filename string) int64 {
	fmt.Printf("Reading file %s\n", filename)
	cals := ParseCalibrations(filename)

	t := int64(0)

	ops := []func(int64, int64) int64{mult, add, ints.Concat64}
	for i := 0; i < len(cals); i++ {
		fmt.Printf("testing %d: with %v\n", cals[i].TestValue, cals[i].Values)
		if IsCalCorrectPartTwo(cals[i].TestValue, 0, cals[i].Values, ops) {
			t += cals[i].TestValue
		}
	}
	return t //362646859298554
}

func isCalibrationCorrect(number int64, numbers []int) bool {
	if number < 0 {
		return false
	}
	if len(numbers) == 0 {
		return number == 0
	}
	v := numbers[len(numbers)-1]
	body := numbers[0 : len(numbers)-1]
	if number%int64(v) == 0 {
		return isCalibrationCorrect(number/int64(v), body) ||
			isCalibrationCorrect(number-int64(v), body)
	} else {
		return isCalibrationCorrect(number-int64(v), body)
	}
}

func IsCalCorrectPartTwo(target, num int64, nums []int, ops []func(int64, int64) int64) bool {
	if num > target {
		return false
	}

	if len(nums) == 0 {
		return num == target
	}

	next := int64(nums[0])
	rest := []int{}
	if len(nums) > 0 {
		rest = nums[1:]
	}

	for _, op := range ops {
		if IsCalCorrectPartTwo(target, op(num, next), rest, ops) {
			return true
		}
	}
	return false
	// return IsCalCorrectPartTwo(target, num*next, rest, ops) ||
	// 	IsCalCorrectPartTwo(target, num+next, rest, ops) ||
	// 	IsCalCorrectPartTwo(target, ints.Concat64(num, next), rest, ops)
}

func add(a, b int64) int64 {
	return a + b
}
func mult(a, b int64) int64 {
	return a * b
}

func isCalibrationCorrectWithConcat(number int64, numbers []int) bool {
	if number < 0 {
		return false
	}

	fmt.Printf("Checking Calibration Step %d: Test: %d, numbers: %v\n", len(numbers), number, numbers)
	if len(numbers) == 0 {
		if number == 0 {
			fmt.Println("correct")
			return true
		}
		return false
	}
	body := numbers[:len(numbers)-1]
	v := numbers[len(numbers)-1]

	concat := int64(0)
	if len(numbers) > 1 {
		concat = int64(ints.Concat(numbers[len(numbers)-2], numbers[len(numbers)-1]))
	}

	if number%int64(v) == 0 {
		return isCalibrationCorrectWithConcat(number/int64(v), body) ||
			isCalibrationCorrectWithConcat(number-int64(v), body) ||
			len(numbers) > 1 && isCalibrationCorrectWithConcat(number-concat, numbers[:len(numbers)-2])
	} else {
		return isCalibrationCorrectWithConcat(number-int64(v), body) ||
			len(numbers) > 1 && isCalibrationCorrectWithConcat(number-concat, numbers[:len(numbers)-2])
	}
}
