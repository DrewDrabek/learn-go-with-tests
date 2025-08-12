package iteration

import "strings"

func Repeat(character string, countAmount int) string {
	var repeatCount = countAmount
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
