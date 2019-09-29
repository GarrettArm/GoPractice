// experiments with the iota generator

package main

import (
	"fmt"
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Multiple int

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
)

func main() {
	fmt.Printf("%[1]v\t%[1]T\n", Sunday)
	fmt.Println("showing KiB, MiB, GiB\n")
	fmt.Printf("%[1]v\t%[1]T\n", KiB)
	fmt.Printf("%[1]v\t%[1]T\n", MiB)
	fmt.Printf("%[1]v\t%[1]T\n", GiB)
}
