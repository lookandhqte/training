package tbank

import (
	"fmt"
	"training/usecases"
)

func kostyaAndPaper(numsAmount int, operationsAmount int, nums *[]int) (sumDiff int) {
	numsArr := *nums
	for operationsAmount != 0 {
		maxNum := 0
		for i := 0; i < numsAmount; i++ {
			if numsArr[i] > maxNum {
				maxNum = numsArr[i]
			}
		}
		maxCap := 1
		for maxNum/10 != 0 {
			maxCap *= 10
			maxNum /= 10
		}
		minNum := 9999
		minId := 0
		for i := 0; i < len(numsArr); i++ {
			if actualNum := numsArr[i] / maxCap; actualNum != 0 && actualNum < minNum && actualNum != 9 {
				minNum = actualNum
				minId = i
			}
		}
		if maxCap > 1 {
			sumDiff += 9*maxCap + numsArr[minId]%maxCap - numsArr[minId]
			numsArr[minId] = 9*maxCap + numsArr[minId]%maxCap - numsArr[minId]

		} else {
			sumDiff += 9*maxCap - numsArr[minId]
			numsArr[minId] += 9*maxCap - numsArr[minId]
		}
		operationsAmount--
	}
	return
}

func testKostya() {
	test := kostyaAndPaper
	testCases := []usecases.TestCase{
		{
			Input:  &([]int{1, 2, 1, 3, 5}),
			Output: 16,
		},
		{
			Input:  &([]int{99, 5, 85}),
			Output: 10,
		},
	}
	amount := 2
	for _, tc := range testCases {
		numsAmLen := len(*tc.Input.(*[]int))
		output := test(numsAmLen, amount, tc.Input.(*[]int))
		if output != tc.Output {
			fmt.Printf("diff\nyour: %v, needed: %v\n", output, tc.Output)
		} else {
			fmt.Println("well done")
			fmt.Printf("result: %v, needed %v\n", output, tc.Output)
		}
		amount--
	}
}

// func main() {
// 	inputFile, _ := os.Open("input.txt")
// 	defer inputFile.Close()

// 	var n, k int
// 	fmt.Fscan(inputFile, &n, &k)

// 	nums := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		fmt.Fscan(inputFile, &nums[i])
// 	}
// 	var res int
// 	for k != 0 {
// 		maxNum := 0
// 		maxCap := 1
// 		for i := 0; i < n; i++ {
// 			if nums[i] > maxNum {
// 				maxNum = nums[i]
// 				copyNum := maxNum
// 				for copyNum/10 != 0 {
// 					maxCap *= 10
// 					copyNum /= 10
// 				}
// 			}
// 		}
// 		minNum := 9999
// 		minId := 0
// 		for i := 0; i < n; i++ {
// 			if actualNum := nums[i] / maxCap; actualNum != 0 && actualNum < minNum && actualNum != 9 {
// 				minNum = actualNum
// 				minId = i
// 			}
// 		}
// 		if maxCap > 1 {
// 			res += 9*maxCap + nums[minId]%maxCap - nums[minId]
// 			nums[minId] = 9*maxCap + nums[minId]%maxCap - nums[minId]

// 		} else {
// 			res += 9*maxCap - nums[minId]
// 			nums[minId] += 9*maxCap - nums[minId]
// 		}
// 		k--
// 	}
// 	outputFile, _ := os.Create("output.txt")
// 	defer outputFile.Close()
// 	fmt.Fprint(outputFile, res)
// }
