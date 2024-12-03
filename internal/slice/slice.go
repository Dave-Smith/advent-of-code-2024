package slice

func PopFromSlice[T any](s []T, index int) []T {
	if index >= len(s) {
		cp := make([]T, len(s))
		copy(s, cp)
		return cp
	}
	if index < 0 {
		return []T{}
	}
	return append(s[:index], s[index+1:]...)
}
