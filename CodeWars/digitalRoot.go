package codewars

func DigitalRoot(n int) int {
	res := 0
	for res/10 != 0 || res == 0 {
		for n/10 != 0 {
			res += n % 10
			n /= 10
		}
		res += n % 10
		n = res
		res = 0
		if n/10 == 0 {
			return n
		}
	}
	return res
}
