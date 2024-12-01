package utils

func AbsDiff[T Number](a, b T) T {
	var zero T
	diff := a - b
	if diff < zero {
		diff = -diff
	}
	return diff
}
