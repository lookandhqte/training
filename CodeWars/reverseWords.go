package codewars

func ReverseWords(str string) string {
	res := ""
	str += " "
	startPosition := 0
	for i := 0; i < len(str); i++ {
		if string(str[i]) == " " {
			for j := i - 1; j >= startPosition; j-- {
				res += string(str[j])
			}
			if i != len(str)-1 {
				res += " "
			}
			startPosition = i + 1
		}
	}
	return res
}
