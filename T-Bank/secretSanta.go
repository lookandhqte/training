package tbank

import "fmt"

func secretSanta() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	count := make([]int, n+2)
	for _, elem := range arr {
		if elem > 0 && elem < n+1 {
			count[elem]++
		}
	}
	isFunctional := true
	for _, elem := range count {
		if elem != 1 {
			isFunctional = false
		}
	}
	if isFunctional {
		fmt.Println("-1, -1")
		return
	}

	res := []int{-1, -1}

	fmt.Println(res)
}
