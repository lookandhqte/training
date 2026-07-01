package tbank

import (
	"fmt"
)

func sashaAndPalyndrome() {
	var s string
	fmt.Scan(&s)

	left, right := 0, len(s)-1

	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			if s[right] == 'a' {
				right--
			} else {
				fmt.Println("No")
				return
			}
		}
	}

	fmt.Println("Yes")
}
