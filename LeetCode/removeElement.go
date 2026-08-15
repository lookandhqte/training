package leetcode

import "fmt"

func removeElement(nums []int, val int) int {
	n := len(nums)
	amount := 0
	newArr := make([]int, n)
	for i := 0; i < n; i++ {
		if nums[i] != val {
			newArr[amount] = nums[i]
			amount++
		}
	}
	for i := 0; i < amount; i++ {
		nums[i] = newArr[i]
	}
	return amount
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
}
