package main

import (
	"fmt"
)

func twoSum(arrays []int,target int) []int {
	hm := make(map[int]int)
	for i, v := range arrays {
		hm[v] = i
	}
	for k,v := range hm {
		res := target - k
		if _,ok := hm[res]; !ok {
			//fmt.Printf("Key %s is not exist", res)
			continue
		}
		return []int{v, hm[res]}
	}
	return nil
}

// "111" + "1" = "1000"
//TODO
func twoSumBinary(s1 string, s2 string) string {
	return ""
}

func main() {
	testArray := []int{1,2,3,4,5}
	target := 4
	result := twoSum(testArray, target)
	fmt.Printf("Target is %d , get index1: %d  index2: %d \n", target, result[0], result[1])

	s1 := "111"
	s2 := "1"
	result = twoSumBinary(s1, s2)
	fmt.Printf("Sum %s and %s with Binary is %si. \n", s1, s2, result)
}
