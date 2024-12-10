package ints

import (
	"fmt"
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

func MultiplyAll(numbers []int) int64 {
	t := int64(1)
	for i := 0; i < len(numbers); i++ {
		t *= int64(numbers[i])
	}
	return t
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

func Includes64(nums []int64, num int64) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == num {
			return true
		}
	}
	return false
}

func IncludesAny64(a, b []int64) bool {
	var empty struct{}
	bMap := make(map[int64]struct{})
	for i := 0; i < len(b); i++ {
		bMap[b[i]] = empty
	}

	for i := 0; i < len(a); i++ {
		if _, ok := bMap[a[i]]; ok {
			return true
		}
	}
	return false
}

func FromSubstring(str string, from int, til int) int64 {
	val, err := strconv.ParseInt(str[from:til], 10, 64)
	if err != nil {
		panic(err)
	}
	return val
}
func FindFactors(n int64) []int64 {
	factors := []int64{}
	for i := int64(1); i <= n; i++ {
		if n%i == 0 && i != 1 && i != n {
			factors = append(factors, i)
		}
	}
	return factors
}

func Concat(a, b int) int {
	v, e := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	if e != nil {
		panic(e)
	}
	return v
}

func Concat64(a, b int64) int64 {
	v, e := strconv.ParseInt(fmt.Sprintf("%d%d", a, b), 10, 64)
	if e != nil {
		panic(e)
	}
	return v
}

func GetLastDigit(num int) int {
	return (num / 10) % 10
}

func GetLastTwoDigits(num int) int {
	return (num / 100) % 100
}
