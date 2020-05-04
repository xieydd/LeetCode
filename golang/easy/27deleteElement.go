package easy

func DeleteElement(nums []int, val int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	i := 0
	for j:=0; j < length;j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
