package weekpreparation

func lonelyinteger(a []int32) int32 {
	appearsAmount := make(map[int32]int32)
	n := len(a)
	for i := 0; i < n; i++ {
		appearsAmount[a[i]]++
	}
	for key, val := range appearsAmount {
		if val == 1 {
			return key
		}
	}
	return 0
}

func diagonalDifference(arr [][]int32) int32 {
	var sum int32 = 0
	for i, j := 0, len(arr)-1; i < len(arr); i, j = i+1, j-1 {
		sum += arr[i][j] - arr[i][i]
	}
	if sum < 0 {
		sum *= -1
	}
	return sum
}

func countingSort(arr []int32) []int32 {
	res := make([]int32, 100)
	for i := 0; i < len(arr); i++ {
		res[arr[i]]++
	}
	return res
}

func flippingMatrix(matrix [][]int32) int32 {
	n := len(matrix) / 2
	var sum int32 = 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a := matrix[i][j]
			b := matrix[i][2*n-1-j]
			c := matrix[2*n-1-i][j]
			d := matrix[2*n-1-i][2*n-1-j]

			max := a
			if b > max {
				max = b
			}
			if c > max {
				max = c
			}
			if d > max {
				max = d
			}
			sum += max
		}
	}
	return sum
}
