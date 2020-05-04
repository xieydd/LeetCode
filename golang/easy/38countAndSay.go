package easy
import (
	"errors"
	"strings"
	"strconv"
)

func CountAndSay(n int) string {
	if n < 0 {
		errors.New("Not Support number < 0")
	}

	if n == 1{
		return "1"
	}

	prev := CountAndSay(n-1)
	count, lastIndex := 1, len(prev)-1
	var result strings.Builder
	for i := range prev[:lastIndex] {
		if prev[i] != prev[i+1] {
			result.WriteString(strconv.Itoa(count))
			result.WriteByte(prev[i])
			count=1
		}else {
			count++
		}
	}
	result.WriteString(strconv.Itoa(count))
	result.WriteByte(prev[lastIndex])

	return result.String()
}
