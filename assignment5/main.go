package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Remove hashing")
	var input, finalString string // := "##12345"
	arr := [3]int64{1234, 1564, 67892}
	fmt.Print("Enter number: ")
	fmt.Scan(&input)
	var flag bool
	//fmt.Print(len(input))
	for _, v := range input {
		if v != '#' {
			finalString = finalString + string(v)
		}
	}
	x, err := strconv.ParseInt(finalString, 0, 0)
	for _, value := range arr {
		if err == nil {
			if value == x {
				flag = true
			}
		}
	}
	if flag {
		fmt.Println("Value exists in array")
	} else {
		fmt.Println("Value does not exists in array")
	}

}
