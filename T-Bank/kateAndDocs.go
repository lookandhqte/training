package tbank

import (
	"fmt"
	"training/usecases"
)

func kateAndDocs(docsAmount int, timeCollegueLeaves int, floors *[]int, idCollegue int) (floorsAmount int) {
	floorArr := *floors
	idCollegue--
	target := floorArr[idCollegue]
	if target-floorArr[0] > timeCollegueLeaves {
		if floorArr[docsAmount-1]-target > timeCollegueLeaves {
			if target-floorArr[0] > floorArr[docsAmount-1]-target {
				for i := idCollegue; i < docsAmount-1; i++ {
					floorsAmount += floorArr[i+1] - floorArr[i]
				}
				floorsAmount += floorArr[docsAmount-1] - floorArr[idCollegue-1]
				for i := idCollegue - 1; i > 0; i-- {
					floorsAmount += floorArr[i] - floorArr[i-1]
				}
			} else {
				for i := idCollegue; i > 0; i-- {
					floorsAmount += floorArr[i] - floorArr[i-1]
				}
				floorsAmount += floorArr[idCollegue+1] - floorArr[0]
				for i := idCollegue + 1; i < docsAmount-1; i++ {
					floorsAmount += floorArr[i+1] - floorArr[i]
				}
			}
		}
	} else if target-floorArr[0] <= timeCollegueLeaves || floorArr[docsAmount-1]-target <= timeCollegueLeaves {
		floorsAmount = floorArr[docsAmount-1] - floorArr[0]
	}
	return
}

func testKate() {
	test := kateAndDocs
	testCases := []usecases.TestCase{
		{
			Input:  usecases.KateInput{DocsAmount: 5, TimeCollegueLeaves: 5, Floors: []int{1, 4, 9, 16, 25}, IdCollegue: 2},
			Output: 24,
		}, {
			Input:  usecases.KateInput{DocsAmount: 6, TimeCollegueLeaves: 4, Floors: []int{1, 2, 3, 6, 8, 25}, IdCollegue: 5},
			Output: 31,
		},
	}
	failedTests := 0
	results := make([]string, len(testCases))
	for _, tc := range testCases {
		input := tc.Input.(usecases.KateInput)
		output := test(input.DocsAmount, input.TimeCollegueLeaves, &input.Floors, input.IdCollegue)
		if output != tc.Output {
			failedTests++
			results = append(results, fmt.Sprintf("Input: %v\nReceived: %v\nWanted: %v\n\n", input, output, tc.Output))
		}
	}
	if failedTests != 0 {
		fmt.Printf("❌ Tests failed. Fail percent %.2f\n", (float64(len(testCases))/float64(failedTests))*100)
		for _, elem := range results {
			fmt.Print(elem)
		}
	} else {
		fmt.Println("✅ All tests passed, congrats 🎉")
	}
}
