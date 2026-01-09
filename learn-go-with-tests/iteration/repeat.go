package iteration

import "strings"

func Repeat(char string, repetitions int) string {
	var repeated strings.Builder
	for range repetitions {
		repeated.WriteString(char)
	}
	return repeated.String()
}
