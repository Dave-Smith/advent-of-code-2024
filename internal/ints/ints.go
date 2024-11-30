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
