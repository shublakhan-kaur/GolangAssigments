package addition

import "fmt"

func Add(a, b float64) float64 {
	fmt.Print(a, " + ", b, " = ")
	return a + b
}
