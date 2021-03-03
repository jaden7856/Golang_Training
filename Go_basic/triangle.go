package main

import "fmt"

func main() {
	var input int
	//i,j는 두 개의 반복문에 쓰일 변수

	fmt.Scan(&input)

	for i := 0; i < input; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("o ")
		}
		fmt.Printf("* \n")
	}
}
