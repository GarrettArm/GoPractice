package main

import (
	"fmt"
)

type A struct {
	X int
	Y int
}

func (a A) MultiplyX(n int) int {
	return (a.X * n)
}

func (a A) MultiplyY(p *int) int {
	return (a.Y * *p)
}

func main() {
	y := A{4, 5}
	fmt.Printf("%v\n", y)
	z := y.MultiplyX(2)
	fmt.Printf("%v\n", z)
	n := 2
	q := y.MultiplyY(&n)
	fmt.Printf("%v\n", q)
}
