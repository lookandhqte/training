package hexlet

import (
	"fmt"
	"training/usecases"
)

// Доделать
func solution(s string) bool {
	var ans bool
	lettersMap := make(map[rune]int)
	for _, r := range s {
		lettersMap[r]++
	}
	evenAmount := 0
	for _, i := range lettersMap {
		if i%2 == 1 {
			evenAmount++
		}
	}
	if evenAmount <= 1 {
		ans = true
	}
	return ans
}

func main() {
	testCases := []usecases.TestCase{
		{
			Input:  "abcabc",
			Output: true,
		}, {
			Input:  "a",
			Output: true,
		}, {
			Input:  "aa",
			Output: true,
		}, {
			Input:  "aab",
			Output: true,
		}, {
			Input:  "aabb",
			Output: true,
		}, {
			Input:  "aabbc",
			Output: true,
		}, {
			Input:  "racecar",
			Output: true,
		}, {
			Input:  "carerac",
			Output: true,
		}, {
			Input:  "aabbcc",
			Output: true,
		}, {
			Input:  "aabbccc",
			Output: true,
		}, {
			Input:  "aabbccdd",
			Output: true,
		}, {
			Input:  "aabbccdde",
			Output: true,
		}, {
			Input:  "aaa",
			Output: true,
		}, {
			Input:  "aaaa",
			Output: true,
		}, {
			Input:  "aaaab",
			Output: true,
		}, {
			Input:  "aabbccddee",
			Output: true,
		}, {
			Input:  "aabbccddeef",
			Output: true,
		}, {
			Input:  "aabbccddeeff",
			Output: true,
		}, {
			Input:  "xxyyzz",
			Output: true,
		}, {
			Input:  "xyzxyz",
			Output: true,
		}, {
			Input:  "noon",
			Output: true,
		}, {
			Input:  "",
			Output: true,
		}, {
			Input:  "ab",
			Output: false,
		}, {
			Input:  "aabbcd",
			Output: false,
		}, {
			Input:  "aaabbb",
			Output: false,
		}, {
			Input:  "abc",
			Output: false,
		}, {
			Input:  "abcdef",
			Output: false,
		}, {
			Input:  "aaabbb",
			Output: false,
		}, {
			Input:  "aaabbbc",
			Output: false,
		}, {
			Input:  "abcd",
			Output: false,
		}, {
			Input:  "xyz",
			Output: false,
		},
	}
	test := solution
	wrongTests := make(map[int]bool) //id result
	res := false
	for id, tc := range testCases {
		res = test(tc.Input.(string))
		if res != tc.Output {
			wrongTests[id] = res
		}
	}
	if len(wrongTests) == 0 {
		fmt.Println("All tests passed!")
	} else {
		fmt.Printf("%.2f percents of tests passed\n", float64(len(testCases)/len(wrongTests)))
		for id, result := range wrongTests {
			fmt.Printf("Test %v\n", id)
			fmt.Printf("Expected: %v\n", testCases[id].Output)
			fmt.Printf("Received: %v\n", result)
		}
	}
}
