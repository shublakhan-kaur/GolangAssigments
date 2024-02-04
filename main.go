package main

import (
	"GolangAssigments/addition"
	"GolangAssigments/division"
	"GolangAssigments/multiply"
	"GolangAssigments/subtraction"
	"fmt"
)

func main() {

	option := 2
	switch option {
	case 1:
		fmt.Println("Calculator")
		fmt.Println(addition.Add(2, 4))
		fmt.Println(subtraction.Subtract(10, 7))
		fmt.Println(multiply.Multiply(2, 9))
		fmt.Println(division.Divide(18, 2))
	case 2:
		var n int
		fmt.Println("Fibonacci Series")
		fmt.Print("Enter Number : ")
		fmt.Scanf("%d", &n)
		Fibonacci(n)

	case 3:
		arr := [5]int{1, 2, 3, 4, 5}
		var num int
		fmt.Print("Enter the number to find the position ", arr, " : ")
		fmt.Scanf("%d", &num)
		findPosition(arr, num)

	}

}

func Fibonacci(num int) {
	for a, b := 0, 1; a <= num; a, b = b, a+b {
		fmt.Printf("%d\t", a)
	}
}

func findPosition(arr [5]int, num int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			fmt.Println("Position of", num, "is", i)
		}
	}

}
