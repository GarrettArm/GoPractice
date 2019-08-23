package main

import (
	"fmt"
)

type Thing struct {
	Flavor string
	Count  int
}

func (t Thing) Describe() []string {
	return []string{t.Flavor, string(t.Count)}
}

type myInt int

func (i myInt) PlusTwo() int {
	return int(i + 2)
}

func main() {
	s := Thing{"Guatemala", 1}
	fmt.Printf("%v\n", s)
	fmt.Println(s.Describe())
	i := myInt(5)
	fmt.Println(i.PlusTwo())
}
