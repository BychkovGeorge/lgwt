package loops

import "strings"

func Repeat(symbol string, count int) string {
	var sb strings.Builder
	for range count {
		sb.WriteString(symbol)
	}
	return sb.String()
}
