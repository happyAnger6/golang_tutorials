package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%e %[1]g %[2]e %[2]g\n", math.MaxFloat32, math.MaxFloat64)
	var f float32 = 16777216
	fmt.Printf("%v", f==f+1)

	for x:= 0; x < 8; x++ {
		fmt.Printf("x = %d ex = %8.3f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z);

	n := math.NaN()
	fmt.Println(n==n, n<n, n>n)
}
