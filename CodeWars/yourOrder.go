package codewars

import "strings"

func Order(sentence string) string {
	afterFields := strings.Fields(sentence)
	result := make([]string, len(afterFields))
	for _, word := range afterFields {
		for _, r := range word {
			if r >= '0' && r <= '9' {
				result[int(r)-'0'-1] = word
			}
		}
	}
	return strings.Join(result, " ")
}
