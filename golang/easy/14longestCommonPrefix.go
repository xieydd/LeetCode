package main

import (
	"fmt"
)

func LongestCommonPrefiex(strs []string) string {
	if strs == nil || len(strs) == 0 { return "" }
	for i, value := range []byte(strs[0]) {
		for _, str := range strs {
			if len(str) == i || []byte(str)[i] != value {
				return string([]byte(str)[:i])
			}
		}
	}
	return strs[0]
}

func main() {
	strs := []string{"flower", "fly", "flight"}
	output := LongestCommonPrefiex(strs)
	fmt.Printf("Longet Common Prefiex is %s\n", output)
}
