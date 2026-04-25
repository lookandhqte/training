package solved

import "fmt"

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	if len(strs) == 1 {
		result = append(result, strs)
		return result
	}
	isChecked := make(map[int]bool)
	for i := 0; i < len(strs); i++ {
		if isChecked[i] {
			continue
		}
		isChecked[i] = true
		anagramGroup := []string{}
		anagramGroup = append(anagramGroup, strs[i])
		lettersIMap := make(map[rune]int)
		for _, r := range strs[i] {
			lettersIMap[r] += 1
		}
		for j := i + 1; j < len(strs); j++ {
			if len(strs[j]) != len(strs[i]) {
				continue
			}
			isAnagram := true
			lettersJMap := make(map[rune]int)
			for _, r := range strs[j] {
				lettersJMap[r] += 1
			}
			for _, r := range strs[j] {
				if res, ok := lettersIMap[r]; !ok {
					isAnagram = false
				} else if res != lettersJMap[r] {
					isAnagram = false
				}
			}
			if !isAnagram {
				continue
			}
			isChecked[j] = true
			anagramGroup = append(anagramGroup, strs[j])
		}
		result = append(result, anagramGroup)
	}
	return result
}

type testCase struct {
	Input  []string
	Output [][]string
}

func main() {
	test := groupAnagrams
	testCases := []testCase{
		{
			Input:  []string{"act", "pots", "tops", "cat", "stop", "hat"},
			Output: [][]string{{"hat"}, {"act", "cat"}, {"pots", "stop", "tops"}},
		},
		{
			Input:  []string{"x"},
			Output: [][]string{{"x"}},
		},
		{
			Input:  []string{"", "b", ""},
			Output: [][]string{{"", ""}, {"b"}},
		},
		{
			//Считать количество букв, а не просто мапа true
			Input:  []string{"ddddddddddg", "dgggggggggg"},
			Output: [][]string{{"ddddddddddg"}, {"dgggggggggg"}},
		},
	}

	for id, tc := range testCases {
		res := test(tc.Input)
		fmt.Println("--------------------------")
		fmt.Println("Test ", id)
		fmt.Println("Your result: ", res)
		fmt.Println("Right result: ", tc.Output)
	}
}
