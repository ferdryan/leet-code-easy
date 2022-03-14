package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, D, x1, x2, rp, ip float64
	fmt.Scanf("%f \n", &a)
	fmt.Scanf("%f \n", &b)
	fmt.Scanf("%f \n", &c)
	fmt.Scanf("%f \n", &x1)
	fmt.Scanf("%f \n", &x2)

	D = (b * 2) - (4 * a * c)
	if D >= 0 {
		r1 := ((-1 * b) + math.Sqrt(D)) / (2 * a)
		r2 := ((-1 * b) - math.Sqrt(D)) / (2 * a)
		fmt.Println("Root 1", r1)
		fmt.Println("Root 2", r2)
	} else {
		rp = (-1 * b) / (2 * a)
		ip = math.Sqrt((-1 * D)) / (2 * a)
		fmt.Println(rp + math.J0(ip))
		fmt.Println(rp - math.J0(ip))
	}
}
