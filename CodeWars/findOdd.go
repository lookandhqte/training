package codewars

func FindOdd(seq []int) (result int) {
	for _, num := range seq {
		result ^= num
	}
	return
}
