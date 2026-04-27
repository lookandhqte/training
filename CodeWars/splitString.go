package codewars

func Solution(str string) []string {
	result := make([]string, 0)
	if len(str)%2 != 0 {
		str += "_"
	}
	for i := 0; i < len(str); i += 2 {
		result = append(result, str[i:i+1])
	}
	return result
}
