package leetcode

func majorityElement(nums []int) []int {
	elemsAmount := make(map[int]int) // map[num]amount
	n := len(nums)
	res := []int{}
	maxTimes := n / 3
	for i := 0; i < n; i++ {
		elemsAmount[nums[i]]++
	}
	for num, amount := range elemsAmount {
		if amount > maxTimes {
			res = append(res, num)
		}
	}
	return res
}
