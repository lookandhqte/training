package tbank

import (
	"fmt"
)

func mashaAnyaAndRocks() {
	var n int
	fmt.Scan(&n)
	res := "Anya"
	if n%2 != 0 {
		res = "Masha"
	}
	fmt.Println(res)
}
