func findClosestNumber(nums []int) int {
	ans, min := nums[0], nums[0]
	if min < 0 {
		min = 0 - min
	}

	for _, v := range nums {
		if v == min || (0-v) == min {
			if v > ans {
				ans = v
			}
		} else if v < 0 && (0-v) < min {
			ans = v
			min = 0 - v
		} else if v >= 0 && v < min {
			ans = v
			min = v
		}
	}
	return ans
}
