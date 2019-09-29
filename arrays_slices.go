// experiments with arrays and slices

package main

import (
	"fmt"
)

func createAndRead() {

	// array of all months
	months := [...]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
	// slices of months
	spring := months[3:6]
	hot := months[5:13]

	// new slice from thin air
	colors := []string{"red", "blue", "orange"}

	fmt.Printf("%v\n", spring)
	fmt.Println(colors[0])

	for _, i := range hot {
		fmt.Println(i)
	}
}

func appendTricks() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Printf("cap=%d\t%v\t\t%p\n", cap(x), x, &x)
	for i := 1; i < 10; i++ {
		x = append(x, i)
		fmt.Printf("cap=%d\t%v\t\t%p\n", cap(x), x, &x)
	}
}

func main() {
	createAndRead()
	appendTricks()
}
