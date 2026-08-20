package weekpreparation

import (
	"fmt"
	"strconv"
)

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
	if s[8] == 'A' {
		h, _ := strconv.Atoi(string(s[:2]))
		if h == 12 {
			return "00" + s[2:8]
		}
		return s[:8]
	}
	h, _ := strconv.Atoi(string(s[:2]))
	if h == 12 {
		return s[:8]
	}
	h += 12
	return strconv.Itoa(h) + s[2:8]
}
