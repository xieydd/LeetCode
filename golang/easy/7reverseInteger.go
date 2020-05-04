package main

import (
	"fmt"
)

// input: 123 output: 321 Notice: -321 and 120
const (
	MININT32 int32 = 1- 1<<31
	MAXINT32 int32 = 1<<31 - 1
)
// Push the last to the first directly is wrong, because don`t consider the overflow situation
func ReverseInteger(in int32) int32 {
	var rev int32 = 0
	for ;in != 0; in = in/10 {
		last := in % 10
		if rev > MAXINT32/10 || (rev == MAXINT32/10 && last > 7) {
			fmt.Printf("Above Overflow\n")
			return 0
		}
		if rev < MININT32/10 || (rev == MININT32/10 && last < -8) {
			fmt.Printf("Bottom Overflow\n")
			return 0
		}
		rev = rev*10 + last
	}
	return rev
}

func main() {
	var test int32 = 4541549
	result := ReverseInteger(test)
	fmt.Printf("Reverse %d is %d\n", test, result)
}
