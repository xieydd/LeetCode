package easy
import (
	"errors"
)

func DeleteSortedArray(nums []int) int{
	if len(nums) == 0 {
		print(errors.New("The Array is empty"))
		return 0
	}

	i := 0
	size := len(nums)
	for j := 1; j < size; j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i+1
}
