package main

import (
	"fmt"
)

// Shouldn`t convert int to string
// Can use ReverseInteger get result, get input == result
func PalindromeNumber(num int32) bool {
	if num < 0 || (num % 10 == 0 && num != 0) {
		return false
	}
	var res int32 = 0
	for ; num > res; num = num/10 {
		res = res*10 + num%10
		if num == res {
			return true
		}
	}
	return false
}

func main() {
	var num int32 = 121
	if !PalindromeNumber(num) {
		fmt.Printf("Frue\n")
		return
	}
	fmt.Printf("True\n")
}
