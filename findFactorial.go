package main

import "fmt"

func main() {
	var n int
	factorial := 1
	fmt.Scanf("%d\n", &n)
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	fmt.Println(factorial)
}
