package leetcode

import "fmt"

func summaryRanges(nums []int) []string {
	n := len(nums)
	res := make([]string, 0)
	if n == 0 {
		return res
	}
	firstValue := nums[0]
	prevValue := nums[0]
	for i := 1; i < n; i++ {
		if nums[i]-1 == prevValue {
			prevValue = nums[i]
			continue
		} else {
			var s string
			if firstValue != prevValue {
				s = fmt.Sprintf("%v->%v", firstValue, prevValue)

			} else {
				s = fmt.Sprintf("%v", firstValue)
			}
			res = append(res, s)
			firstValue = nums[i]
			prevValue = nums[i]
		}
	}
	if firstValue != prevValue {
		res = append(res, fmt.Sprintf("%v->%v", firstValue, prevValue))
	} else {
		res = append(res, fmt.Sprintf("%v", firstValue))
	}
	return res
}
