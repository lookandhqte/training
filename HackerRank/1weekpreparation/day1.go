package weekpreparation

import "fmt"

func miniMaxSum(arr []int32) {
	n := len(arr)
	var minElem int32 = 2147483647
	var maxElem int32 = 0
	var sum int64 = 0
	for i := 0; i < n; i++ {
		sum += int64(arr[i])
		if arr[i] < minElem {
			minElem = arr[i]
		}
		if arr[i] > maxElem {
			maxElem = arr[i]
		}
	}
	fmt.Printf("%v %v", sum-int64(maxElem), sum-int64(minElem))
}

func timeConversion(s string) string {
	if s[8] == 'P' {
		return s[:8]
	}
	h := s[:2]
	fmt.Println(h)
	return h
}
