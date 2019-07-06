package main

import (
	"fmt"
)

func TwoSum(arrays []int,target int) []int {
	hm := make(map[int]int, len(arrays))

	for i, v := range arrays {
		if res, ok := hm[target-v]; ok {
			return []int{res, i}
		}

		hm[v] = i
	}

	return nil
}

// "111" + "1" = "1000"
//TODO
//func TwoSumBinary(s1 string, s2 string) string {
//	return ""
//}

func main() {
	testArray := []int{1,2,3,4,5}
	target := 4
	result := TwoSum(testArray, target)
	fmt.Printf("Target is %d , get index1: %d  index2: %d \n", target, result[0], result[1])

	//s1 := "111"
	//s2 := "1"
	//result = TwoSumBinary(s1, s2)
	//fmt.Printf("Sum %s and %s with Binary is %s. \n", s1, s2, result)
}
