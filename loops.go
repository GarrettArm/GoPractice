// experimenting with for loops

package main

import (
	"bufio"
	"fmt"
	"os"
)

func printArgs(args []string) {
	fmt.Println(args)
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func main() {
	args := os.Args[1:]
	printArgs(args)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Printf("%v\n\n", scanner)
		input := scanner.Text()
		fmt.Println(input)
	}
}
