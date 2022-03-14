package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d\n", &a)
	fmt.Scanf("%d\n", &b)
	fmt.Scanf("%d\n", &c)

	if a > b {
		if a > c {
			fmt.Printf("A the largest number")
		} else {
			fmt.Printf("C the largest number")
		}
	} else {
		if b > c {
			fmt.Printf("B the largest number")
		} else {
			fmt.Printf("C the largest number")
		}
	}
}
