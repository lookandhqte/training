package other

type Matrix [][]int64

func (A Matrix) Multiply(B Matrix) (C Matrix) {
	return
}

func (A Matrix) Add(B Matrix) (C Matrix) {
	if len(A) != len(B) {
		return
	}
	for i := 0; i < len(A); i++ {
		row := []int64{}
		for j := 0; j < len(B); j++ {
			row = append(row, A[i][j]+B[i][j])
		}
		C = append(C, row)
	}
	return
}

func (A Matrix) Substract(B Matrix) (C Matrix) {
	if len(A) != len(B) {
		return
	}
	for i := 0; i < len(A); i++ {
		row := []int64{}
		for j := 0; j < len(B); j++ {
			row = append(row, A[i][j]-B[i][j])
		}
		C = append(C, row)
	}
	return
}

// func main() {
// 	matrixA := Matrix{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}
// 	matrixB := Matrix{{9, 8, 7}, {6, 5, 4}, {3, 2, 1}}
// 	matrixC := matrixA.Add(matrixB)
// 	for i := 0; i < len(matrixA); i++ {
// 		if i == 1 {
// 			fmt.Println(matrixA[i], "+", matrixB[i], "=", matrixC[i])
// 			continue
// 		}
// 		fmt.Println(matrixA[i], " ", matrixB[i], " ", matrixC[i])
// 	}
// }
