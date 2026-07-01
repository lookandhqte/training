package tbank

import (
	"fmt"
)

func secretInformation() {
	var n int
	fmt.Scan(&n)
	var prev, curr int
	res := 0
	fmt.Scan(&prev)
	for i := 1; i < n; i++ {
		fmt.Scan(&curr)
		if curr > prev {
			res += curr - prev
		}
		prev = curr
	}
	fmt.Println(res)
}
