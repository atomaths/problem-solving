package main

import "fmt"

func main() {
	var cases int

	fmt.Scanf("%d", &cases)

	for cases > 0 {
		// XXX: algospot.com not accept println of Go built-in fuction.
		fmt.Println("Hello Algospot!")
		cases--
	}
}
