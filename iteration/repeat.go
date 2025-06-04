package iteration

import "strings"

func Repeat(char string, n int) string {
	var res strings.Builder
	for range n {
		res.WriteString(char)
	}
	return res.String()
}
