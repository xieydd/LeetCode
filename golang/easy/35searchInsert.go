package easy

func SearchInsert(nums []int, value int) int {
	length := len(nums)
	if length < 0 {
		return 0
	}
	start := 0
	end := length-1
	if value < nums[start] {
		return 0
	}else if value > nums[end] {
		return length
	}
	for start <= end{
		mid := (start+end)/2
		if end-start==1 && nums[start]<value && nums[end]>value {
			return mid+1
		}else if end == start && nums[start] > value {
			return start
		}else if end == start && nums[start] < value {
			return start+1
		}
		if nums[mid] < value {
			start = mid
		}else if nums[mid] > value {
			end = mid
		}else{
			return mid
		}
	}
	return 0
}
