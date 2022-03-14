package main

import (
	"fmt"
)

func main() {
	var num1, num2, sum int

	fmt.Print("Input Num 1: ")
	fmt.Scanf("%d\n", &num1)
	fmt.Print("Input Num 2: ")
	fmt.Scanf("%d\n", &num2)
	sum = num1 + num2

	fmt.Printf("Sum : %d \n", sum)
}
