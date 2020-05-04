package easy

func MaxSubArray(nums []int) (int, []int) {
	maxCurr := 0
	maxGlobal := nums[0]

	start, end := 0,0
	for i, value := range nums {
		if maxCurr < 0 {
			maxCurr = value
			start = i
		}else {
			maxCurr += value
		}

		if maxCurr > maxGlobal {
			maxGlobal  = maxCurr
			end = i
		}
	}
	return maxGlobal, nums[start:end+1]
}
