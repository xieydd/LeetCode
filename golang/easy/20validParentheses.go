package easy

import (
	"github.com/xieydd/LeetCode/golang/datastruct"
)
func ValidParentheses(s string) bool {
	b := []rune(s)
	maps := make(map[rune]rune)
	maps[')']='('
	maps['}']='{'
	maps[']']='['

	stack := datastruct.NewStack()

	for _,value := range b {
		 if r, ok := maps[value]; ok {
			var topElement rune
			if stack.IsEmpty() {
				topElement = '#'
			}else {
				ele,_ := stack.Pop()
				topElement = ele.(rune)
			}
			if topElement != r {
				return false
			}

		 } else {
			stack.Push(value)
		 }
	}
	return stack.IsEmpty()
}
