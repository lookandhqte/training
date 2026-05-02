package other

import (
	"fmt"
)

func findSquareRoot(square float64, guess float64) float64 {
	for guess*guess != square {
		x_g := square / guess
		guess = (guess + x_g) / 2
	}
	return guess
}

func main() {
	// reader := os.Stdin
	// bufio.NewScanner(reader)
	var x float64 = 16.00
	fmt.Println("x is: ", x)
	var g float64 = 3.00
	fmt.Println("Result is: ", findSquareRoot(x, g))
}
