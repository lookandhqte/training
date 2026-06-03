package main

import (
	"fmt"
)

func main() {
	var l, r int
	fmt.Scan(&l, &r)
	var res int
	for i := l; i < r+1; i++ {
		copyNum := i
		digits := make(map[int]bool)
		isGood := true
		for copyNum != 0 {
			digits[copyNum%10] = true
			copyNum /= 10
			if len(digits) > 1 {
				isGood = false
				break
			}
		}
		if !isGood {
			continue
		} else {
			res++
		}
	}
	fmt.Println(res)
}
