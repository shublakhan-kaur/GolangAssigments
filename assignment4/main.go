package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Sorting of array using Slices")
	arr := [6]int{7, 10, 2, 4, 1, 9}
	fmt.Println("Before sorting", arr)
	sort.Slice(arr[:], func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println("Sorted array", arr)

}
