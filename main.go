package main

import (
	"fmt"
)

//Console input FMT

func main() {
	fmt.Print("Hello!")
	// scanner := bufio.NewScanner(os.Stdin)
}

// //CLOSURES
// func getType(v any) string {
// 	t := fmt.Sprintf("%T", v)
// 	return t
// }

// func addTypesData() func(v any) map[string]any {
// 	typesMap := make(map[string]any)
// 	return func(v any) map[string]any {
// 		typeOfV := getType(v)
// 		typesMap[typeOfV] = v
// 		return typesMap
// 	}
// }
// //TESTS
// func main() {
// 	testing := addTypesData()
// 	test := func() {
// 		testing(42)
// 		testing("hihi")
// 		testing(42.23)
// 		testing(false)
// 		testing('a')
// 		fmt.Println(testing(28i))
// 	}
// 	test()
// }

// // Assignment
// // The Textio API needs a very robust error-logging system
// // so we can see when things are going awry in the back-end
// // system. We need a function that can create a custom
// // "logger" (a function that prints to the console)
// // given a specific formatter.
// //
// // Complete the getLogger function. It should take as input
// // a formatter function and return a new function. The new
// // logger function takes as input two strings and passes
// // them to the formatter, then prints the result. Keep
// // the order of the strings.
// //
// // getLogger takes a function that formats two strings into
// // a single string and returns a function that formats two strings but prints
// // the result instead of returning it
// func getLogger(formatter func(string, string) string) func(string, string) {
// 	newStrings := formater()
// }

// // don't touch below this line

// func test(first string, errors []error, formatter func(string, string) string) {
// 	defer fmt.Println("====================================")
// 	logger := getLogger(formatter)
// 	fmt.Println("Logs:")
// 	for _, err := range errors {
// 		logger(first, err.Error())
// 	}
// }

// func colonDelimit(first, second string) string {
// 	return first + ": " + second
// }
// func commaDelimit(first, second string) string {
// 	return first + ", " + second
// }

// func main() {
// 	dbErrors := []error{
// 		errors.New("out of memory"),
// 		errors.New("cpu is pegged"),
// 		errors.New("networking issue"),
// 		errors.New("invalid syntax"),
// 	}
// 	test("Error on database server", dbErrors, colonDelimit)

// 	mailErrors := []error{
// 		errors.New("email too large"),
// 		errors.New("non alphanumeric symbols found"),
// 	}
// 	test("Error on mail server", mailErrors, commaDelimit)
// }

// func CreateTempFile() {
// 	f, _ := os.Create("temp-42.txt")
// 	defer os.Remove(f.Name()) // executed second
// 	defer f.Close()           // executed first

// 	fmt.Fprintln(f, "How many roads must a man walk down?")
// 	time.Sleep(5 * time.Second)
// }

// //анонимные функции, замыкания, функции как value

// func conversions(converter func(int) string, x, y int) (string, string) {
// 	convertedX := converter(x)
// 	convertedY := converter(y)
// 	return convertedX, convertedY
// }

// func main() {
// 	a, b := 0, 0
// 	fmt.Print("Enter your numbers\nNumber 1: ")
// 	fmt.Scan(&a)
// 	fmt.Print("Well! Number 2: ")
// 	fmt.Scan(&b)
// 	binA, binB := conversions(func(a int) string {
// 		binaryArr := []int{}
// 		for a > 1 {
// 			binaryArr = append(binaryArr, a%2)
// 			a /= 2
// 		}
// 		binaryArr = append(binaryArr, a)
// 		result := ""
// 		for i := len(binaryArr) - 1; i >= 0; i-- {
// 			switch binaryArr[i] {
// 			case 0:
// 				result += "0"
// 			case 1:
// 				result += "1"
// 			}
// 		}

// 		return result
// 	}, a, b)

// 	fmt.Printf("Your converted numbers: %v, %v", binA, binB)
// }

//MATRIX
// func getMatrixSize() (size int) {
// 	fmt.Print("Введите размерность матрицы: ")
// 	fmt.Scan(&size)
// 	return
// }

// func fillTheMatrix(size int) [][]int {
// 	matrix := make([][]int, size) // [] len=0, cap=3 [...]
// 	// Заполнение матрицы
// 	for i := 0; i < size; i++ {
// 		fmt.Print("Введите коэффициенты уравнения: ")
// 		matrixI := make([]int, size)
// 		for j := 0; j < size; j++ {
// 			fmt.Scan(&matrixI[j])
// 		}
// 		matrix[i] = matrixI
// 	}
// 	return matrix
// }

// func fillTheAnswer(size int) []int {
// 	answer := make([]int, size)
// 	fmt.Print("Введите решение уравнения: ")
// 	for i := 0; i < size; i++ {
// 		fmt.Scan(&answer[i])
// 	}
// 	return answer
// }

// func main() {
// 	// 2x1 + 5x2 - 8x3 = 8
// 	// 4x1 + 3x2 - 9x3 = 9
// 	// 2x1 3x2 -5x3 = 7
// 	// x1 = 3 x2 = 2 x3 = 1
// 	size := getMatrixSize()
// 	matrix := fillTheMatrix(size)
// 	answer := fillTheAnswer(size)
// 	//алгоритм Гаусса для СЛАУ
// 	fmt.Println(matrix)
// 	fmt.Println(answer)

// }

//LEETCODE
// func groupAnagrams(strs []string) [][]string {
// 	result := [][]string{}
// 	// act pots tops cat stop hat

// 	slices.SortFunc(strs, func(i, j string) int {
// 		return strings.Compare(i, j)
// 	})

// 	//Bubble Sort
// 	// for i := 0; i < len(strs); i++ {
// 	// 	var tmp string
// 	// 	for j := 0; j < len(strs)-1; j++ {
// 	// 		if i == j {
// 	// 			continue
// 	// 		}
// 	// 		if len(strs[j]) > len(strs[j+1]) {
// 	// 			tmp = strs[j+1]
// 	// 			strs[j+1] = strs[j]
// 	// 			strs[j] = tmp
// 	// 		}
// 	// 	}
// 	// }

// 	fmt.Println(strs)

// 	// act cat hat pots tops stop
// 	for i := 0; i < len(strs); i++ {
// 		resultSubArray := []string{}
// 		iWordLetters := make(map[rune]bool)
// 		for _, r := range strs[i] {
// 			iWordLetters[r] = true
// 		}
// 		resultSubArray = append(resultSubArray, strs[i])
// 		count := 0 //i=3 j=4 true //i=3 j=5 true
// 		for j := i + 1; j < (len(strs)); j++ {
// 			if len(strs[j]) == len(strs[i]) {
// 				isAnagram := true
// 				for _, r := range strs[j] {
// 					if _, ok := iWordLetters[r]; !ok {
// 						isAnagram = false
// 						break
// 					}
// 				}
// 				if isAnagram {
// 					count++
// 					resultSubArray = append(resultSubArray, strs[j])
// 				}
// 			}

// 		}
// 		if count+i < len(strs) {
// 			i += count
// 		}
// 		result = append(result, resultSubArray)
// 	}

// 	return result[:]
// }

// func main() {
// 	strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}
// 	rightResult := [][]string{{"hat"}, {"act", "cat"}, {"stop", "pots", "tops"}}
// 	result := groupAnagrams(strs)
// 	fmt.Println(rightResult)
// 	fmt.Println(result)
// }

// func groupAnagrams(strs []string) [][]string {
// 	result := [][]string{}
// 	//сортировка - сначала строки поменьше длины, в конце длинные
// 	//цикл пока мы не до конца
// 	//создание мапы букв слова и внесение слова в промежуточный массив
// 	//

// 	// //wasInArray мапа проверенных на анаграмму индексов
// 	// wasInArray := make(map[int]bool, 0)

// 	// for i := 0; i < len(strs); i++ {
// 	// 	if !wasInArray[i] {
// 	// 		wasInArray[i] = true
// 	// 		//resultNow массив промежуточного результата [act, cat]
// 	// 		resultNow := []string{}
// 	// 		resultNow = append(resultNow, strs[i])
// 	// 		//lettersMap мапа букв анаграмм
// 	// 		lettersMap := make(map[rune]bool, 0)

// 	// 		for _, r := range strs[i] {
// 	// 			lettersMap[r] = true
// 	// 		}

// 	// 		for j := 0; j < len(strs); j++ {
// 	// 			if !wasInArray[j] {
// 	// 				wasInArray[j] = true
// 	// 				if i == j || len(strs[j]) != len(strs[i]) {
// 	// 					continue
// 	// 				}
// 	// 				isAnagram := true
// 	// 				for _, rj := range strs[j] {
// 	// 					if _, ok := lettersMap[rj]; !ok {
// 	// 						isAnagram = false
// 	// 						break
// 	// 					}
// 	// 				}
// 	// 				if isAnagram {
// 	// 					resultNow = append(resultNow, strs[j])
// 	// 				}
// 	// 			}
// 	// 		}
// 	// 		result = append(result, resultNow)
// 	// 	}
// 	// }
// 	return result
// }

// func compareResults(result []int, rightResult []int) bool {
// 	if len(result) != len(rightResult) {
// 		return false
// 	}
// 	for i := 0; i < len(result); i++ {
// 		if result[i] != rightResult[i] {
// 			return false
// 		}
// 	}
// 	return true
// }
