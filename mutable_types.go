// those containing pointers, behave like python "mutable arguments"
// includes:  slices, pointer, map, function, or channel
// pass pointers as arguments when the argument is a very large thing, or when you want the function to alter the original argument.  Argurments get sent as copies unless they are pointers or are wrappers areound pointers.

package main

import "fmt"

func AlterString(x string) string {
	return "altered" + x
}

func AlterInt(x int) int {
	return x + 5
}

func AlterSlice(x []int) []int {
	x[0] = 5
	return x
}

func AlterArray(x [3]int) [3]int {
	x[0] = 5
	return x
}

func AlterStruct(x StructA) StructA {
	x.X = 5
	return x
}

type StructA struct {
	X int
	Y int
}

func main() {
	str_a := "hi"
	int_a := 5
	slice_a := []int{0, 0, 0}
	array_a := [3]int{0, 0, 0}
	struct_a := StructA{0, 0}
	fmt.Printf("before <hi>\t\tafter <%v>\torig after <%v>\n",
		AlterString(str_a),
		str_a,
	)
	fmt.Printf("before <5>\t\tafter <%v>\t\torig after <%v>\n",
		AlterInt(int_a),
		int_a,
	)
	fmt.Printf("before <[0,0,0]>\tafter <%v>\t\torig after <%v>\n",
		AlterSlice(slice_a),
		slice_a,
	)
	fmt.Printf("before <[0,0,0]>\tafter <%v>\t\torig after <%v>\n",
		AlterArray(array_a),
		array_a,
	)
	fmt.Printf("before {0, 0}\t\tafter <%v>\t\torig after <%v>\n",
		AlterStruct(struct_a),
		struct_a,
	)
}T
