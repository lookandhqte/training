package tbank

import (
	"fmt"
)

func sashaAndTests() {
	var l, r int
	fmt.Scan(&l, &r)
	var res int
	for i := l; i < r+1; i++ {
		if i/11 == 0 {
			res++
		}
	}
	fmt.Println(res)
}
