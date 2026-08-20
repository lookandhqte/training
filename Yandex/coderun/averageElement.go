package coderun

import (
	"cmp"
	"fmt"
	"slices"
)

func averageElement() {
	nums := make([]int, 3)

	for i := 0; i < len(nums); i++ {
		fmt.Scan(&nums[i])
	}

	slices.SortFunc(nums, func(a, b int) int {
		return cmp.Compare(a, b)
	})

	fmt.Println(nums[1])
}
