package main

import "fmt"

func other(x []int) {
	fmt.Printf("%T\n%v\n", x, x)
	for n := range x {
		fmt.Printf("%T\n%v\n", n, n)
	}
	z := []int{3, 4, 5}
	fmt.Printf("%T\n%v\n", z, z)
	x = append(x, z...)
	fmt.Printf("%T\n%v\n", x, x)
}

func main() {
	origSlice := []int{1, 2, 3, 9}
	other(origSlice)
}
