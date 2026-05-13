package leetcode

import "fmt"

func removeElement(nums []int, val int) int {
	removedNumsAmount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			removedNumsAmount++
			for j := i; j < len(nums)-removedNumsAmount; j++ {
				nums[j] = nums[j+1]
			}
			nums[len(nums)-removedNumsAmount] = 0
		}
	}
	return removedNumsAmount
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(nums, val))
}
