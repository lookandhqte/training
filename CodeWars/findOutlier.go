package codewars

// O(n)
func FindOutlier(integers []int) int {
	oddNumbers := make(map[bool]int) // [bool:isOdd]amount
	for i := 0; i < len(integers); i++ {
		if integers[i]%2 == 1 {
			oddNumbers[true] += 1
		} else {
			oddNumbers[false] += 1
		}
	}
	remainder := 0
	if oddNumbers[true] > oddNumbers[false] {
		remainder = 1
	}
	for i := 0; i < len(integers); i++ {
		if integers[i]%2 != remainder {
			return integers[i]
		}
	}
	return 0
}
