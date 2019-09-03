// experiments with pointers

package main

import (
	"fmt"
)

/*
x := 1
p := &x	// p points to x, p is of type *int
*p = 2	// same as writing x = 2
*/

func main() {
	orig := 1
	fmt.Printf("orig is:\t%[2]v\t\t%[1]v\t\t%[1]T\n", orig, &orig)
	// fmt.Printf("&orig is:\t%[1]v\t\t%[1]T\n", &orig)
	copy1, copy2 := orig, orig
	fmt.Printf("copy1 is:\t%[2]v\t\t%[1]v\t\t%[1]T\n", copy1, &copy1)
	fmt.Printf("copy2 is:\t%[2]v\t\t%[1]v\t\t%[1]T\n", copy2, &copy2)
	// fmt.Printf("&copy1 is:\t%[1]v\t\t%[1]T\n", &copy1)
	// fmt.Printf("&copy2 is:\t%[1]v\t\t%[1]T\n", &copy2)
	copy1 = 2
	fmt.Printf("copy1 = \"hi\"\n")
	fmt.Printf("copy1 is:\t%[2]v\t\t%[1]v\t\t%[1]T\n", copy1, &copy1)
	fmt.Printf("copy2 is:\t%[2]v\t\t%[1]v\t\t%[1]T\n", copy2, &copy2)
	// fmt.Printf("copy1 is:\t%[1]v\t\t%[1]T\n", copy1)
	// fmt.Printf("copy2 is:\t%[1]v\t\t%[1]T\n", copy2)
	// fmt.Printf("&copy1 is:\t%[1]v\t\t%[1]T\n", &copy1)
	// fmt.Printf("&copy2 is:\t%[1]v\t\t%[1]T\n", &copy2)

	int1 := 1
	fmt.Println("int1 = ", int1)
	incrByAddress(&int1)
	fmt.Println("int1 = ", int1)

	outsideN := returnPointer()
	fmt.Printf("outsideN is:\t%[1]v\t\t%[1]T\n", outsideN)
	fmt.Printf("*ousideN is:\t%[1]v\t\t%[1]T\n", *outsideN)
	outsideN = &copy1
	fmt.Printf("outsideN is:\t%[1]v\t\t%[1]T\n", outsideN)
	fmt.Printf("*ousideN is:\t%[1]v\t\t%[1]T\n", *outsideN)

	outsideH := returnValue()
	fmt.Printf("outsideH is:\t%[1]v\t\t%[1]T\n", outsideH)
	fmt.Printf("&outsideH is:\t%[1]v\t\t%[1]T\n", &outsideH)
	outsideH = 6
	fmt.Printf("outsideH is:\t%[1]v\t\t%[1]T\n", outsideH)
	fmt.Printf("&outsideH is:\t%[1]v\t\t%[1]T\n", &outsideH)
}

func incrByAddress(w *int) {
	*w++
}

func returnPointer() *int {
	insideN := 5
	fmt.Printf("insideN is:\t%[1]v\t\t%[1]T\n", insideN)
	fmt.Printf("&insideN is:\t%[1]v\t\t%[1]T\n", &insideN)
	return &insideN
}

func returnValue() int {
	insideH := 5
	fmt.Printf("insideH is:\t%[1]v\t\t%[1]T\n", insideH)
	fmt.Printf("&insideH is:\t%[1]v\t\t%[1]T\n", &insideH)
	return insideH
}
