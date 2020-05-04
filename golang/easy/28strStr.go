package easy

func StrStr(haystack string, needle string) int {
	length_h := len(haystack)
	length_n := len(needle)

	if length_h == 0  || length_n == 0 || length_n > length_h {
		return 0
	}

	for i, _ := range haystack {
		if i+length_n > length_h {
			return 0
		}
		if haystack[i:i+length_n] == needle {
			return length_n
		}
	}
	return 0
}
