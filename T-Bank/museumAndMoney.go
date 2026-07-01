package tbank

import (
	"fmt"
)

func museumAndMoney() {
	var n, x int
	fmt.Scan(&n, &x)
	values := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&values[i])
	}
	res := 0
	for i := n - 1; i >= 0; i-- {
		count := x / values[i]
		res += count
		x = x % values[i]
	}
	fmt.Println(res)
}
