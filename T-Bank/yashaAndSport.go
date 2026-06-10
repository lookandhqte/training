package tbank

import "fmt"

func yashaAndSport() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	exceptionArr := []int{-1, -1}
	counter := 0
	for i := 0; i < n; i += 2 {
		if counter == 1 {
			exceptionArr = []int{-1, -1}
			break
		}
		if arr[i]%2 != 1 {
			exceptionArr[counter] = i + 1
			counter++
		}
	}
	for i := 1; i < n; i += 2 {
		if counter == 2 {
			exceptionArr = []int{-1, -1}
			break
		}
		if arr[i]%2 != 0 {
			exceptionArr[counter] = i + 1
			counter++
		}
	}
	fmt.Println(exceptionArr)
}
