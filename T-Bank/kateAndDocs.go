package tbank

import (
	"fmt"
	"training/usecases"
)

func kateAndDocs(docsAmount int, timeCollegueLeaves int, floors *[]int, idCollegue int) (floorsAmount int) {
	floorArr := *floors
	target := floorArr[idCollegue]
	// если дойти с нулевого этажа до этажа коллеги мы не успеваем
	if target-floorArr[0] > timeCollegueLeaves {
		// если мы не успеваем дойти и с последнего этажа до этажа коллеги
		if floorArr[docsAmount-1]-target > timeCollegueLeaves {
			//самый неудачный случай
			// мы двигаемся в сторону которая ближе к концу или к началу
			// начинаем с таргет -> конец/начало -> таргет-1 или таргет+1 -> конец/начало
			// если ближе к концу топать
			if target-floorArr[0] > floorArr[docsAmount-1]-target {
				for i := idCollegue; i < docsAmount-1; i++ {
					floorsAmount += floorArr[i+1] - floorArr[i]
				}
				floorsAmount += floorArr[docsAmount-1] - floorArr[idCollegue-1]
				for i := idCollegue - 1; i > 0; i-- {
					floorsAmount += floorArr[i] - floorArr[i-1]
				}
			} else { // если ближе к началу
				for i := idCollegue; i > 0; i-- {
					floorsAmount += floorArr[i] - floorArr[i-1]
				}
				floorsAmount += floorArr[idCollegue+1] - floorArr[0]
				for i := idCollegue + 1; i < docsAmount-1; i++ {
					floorsAmount += floorArr[i+1] - floorArr[i]
				}
			}
		}
		// если все-таки успеваем к коллеге идя последовательно
	} else if target-floorArr[0] <= timeCollegueLeaves || floorArr[docsAmount-1]-target <= timeCollegueLeaves {
		for _, elem := range floorArr {
			floorsAmount += (elem - floorsAmount)
		}
		floorsAmount -= floorArr[0]
	}
	return
}

func main() {
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
	for _, tc := range testCases {
		input := tc.Input.(usecases.KateInput)
		output := test(input.DocsAmount, input.TimeCollegueLeaves, &input.Floors, input.IdCollegue)
		if output != tc.Output {
			failedTests++
		}
	}
	fmt.Printf("Tests passed: %.2f percents\n", float64(len(testCases))/float64(failedTests)*100)
}
