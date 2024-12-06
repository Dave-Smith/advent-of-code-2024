package ints

import (
	"strconv"
	"strings"
)

func Sum(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}

func FromString(number string) int {
	val, err := strconv.Atoi(strings.TrimSpace(number))
	if err != nil {
		panic(err)
	}
	return val
}

func SliceFromString(numbers string) []int {
	nums := []int{}
	for _, v := range strings.Fields(numbers) {
		if num, err := strconv.Atoi(v); err == nil {
			nums = append(nums, num)
		}
	}
	return nums
}

func SliceFromStringDelim(numbers string, delim rune) []int {
	nums := []int{}
	delimFunc := func(r rune) bool {
		return r == delim
	}
	for _, v := range strings.FieldsFunc(numbers, delimFunc) {
		if num, err := strconv.Atoi(v); err == nil {
			nums = append(nums, num)
		}
	}
	return nums
}

func Abs(n int) int {
	if n > 0 {
		return n
	}

	return 0 - n
}

func Includes(nums []int, num int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == num {
			return true
		}
	}
	return false
}
