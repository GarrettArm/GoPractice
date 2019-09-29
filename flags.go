// demonstrates command-line flags parsing

package main

import (
	"flag"
	"fmt"
)

func main() {
	var verbose = flag.Bool("v", false, "show verbose messages")
	var quick = flag.Bool("q", false, "do quickly")
	var timeout = flag.Int("t", 0, "timeout in seconds")

	flag.Parse()
	fmt.Println("printing commandline flags' values")
	fmt.Println("verbose\t", *verbose)
	fmt.Println("quick\t", *quick)
	fmt.Println("timeout\t", timeout)
}
