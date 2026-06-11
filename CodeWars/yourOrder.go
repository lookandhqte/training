package codewars

// Your task is to sort a given string.
// Each word in the string will
// contain a single number.
// This number is the position
// the word should have in the result.

// Note: Numbers can be from 1 to 9.
// So 1 will be the first word (not 0).

// If the input string is empty,
// return an empty string.
// The words in the input
// String will only contain valid
// consecutive numbers.

// Examples
// "is2 Thi1s T4est 3a"  -->
// "Thi1s is2 3a T4est"
// "4of Fo1r pe6ople g3ood th5e the2"  -->
// "Fo1r the2 g3ood 4of th5e pe6ople"
// ""  -->  ""

func Order(sentence string) (result string) {
	if len(sentence) == 0 {
		return ""
	}
	ourOrder := []int{}
	words := []string{}
	word := ""
	for _, r := range sentence {
		if r == 32 {
			words = append(words, word)
			word = ""
			continue
		}
		word += string(r)
		if r > 47 && r < 58 {
			ourOrder = append(ourOrder, int(r))
		}
	}
	words = append(words, word)
	for i := 0; i < len(ourOrder); i++ {
		min := 1000
		minId := 0
		for id, elem := range ourOrder {
			if elem < min {
				min = elem
				minId = id
			}
		}
		if i == len(ourOrder)-1 {
			result += words[minId]
		} else {
			result += words[minId] + " "
		}
		ourOrder[minId] = 1000
	}
	return
}
