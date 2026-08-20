package main

import (
	"fmt"
)

func timeConversion(s string) string {
	if s[8] == 'P' {
		return s[:8]
	}
	h := s[:2]
	fmt.Println(h)
	return h
}

func main() {
	fmt.Println(timeConversion("12:01:00AM"))
}
