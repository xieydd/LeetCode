package easy

func LengthOfLastWord(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	count := 0
	for i:=length-1;i>=0;i-- {
		if string(s[i]) != " " {
			count++
		}

		if string(s[i]) == " " && count != 0 {
			break
		}
	}

	return count
}
