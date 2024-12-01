package utils

func Sum[T Number](nums []T) T {
	var sum T
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Average[T Number](nums []T) float64 {
	if len(nums) == 0 {
		return 0
	}
	total := Sum(nums)
	return float64(total) / float64(len(nums))
}

func Max[T Number](nums []T) T {
	if len(nums) == 0 {
		var zero T
		return zero
	}
	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}
	return max
}

func Min[T Number](nums []T) T {
	if len(nums) == 0 {
		var zero T
		return zero
	}
	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	return min
}
