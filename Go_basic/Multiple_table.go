package main

import "fmt"

func main() {
	var dan int

	fmt.Scan(&dan)

	for i := 1; i < 10; i++ {
		result := dan * i
		fmt.Printf("%d x %d = %d\n", dan, i, result)
	}
}
