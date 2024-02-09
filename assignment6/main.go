package main

import "fmt"

func main() {
	fmt.Println("Days of the week")
	var day string
	var k, value int
	fmt.Print("Enter day of the week: ")
	fmt.Scanln(&day)
	fmt.Print("Enter range: ")
	fmt.Scanln(&k)
	week := map[int]string{1: "Mon", 2: "Tues", 3: "Wed", 4: "Thurs", 5: "Fri", 6: "Sat", 7: "Sun"}
	for idx, val := range week {
		// fmt.Printf("%v\t%v\n", idx, val)
		if val == day {
			//fmt.Println(val, " ", idx, " ", k, " ", (idx + k), (idx+k)%7)
			value = (idx + k) % 7
		}
	}
	fmt.Println("Day = ", week[value])
}
