package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for it := 0;; it++ {
		zL := z
		z = z - (((z * z) - x) / (2 * z))
                if (math.Abs(zL - z) < 0.000000001) {
			fmt.Println(it)
			break
		}
	}
	return z
}

func main() {
	for val := float64(0); val < 10; val++ {
		fmt.Println("Square Root of", val, "is:", Sqrt(val), math.Sqrt(val))
	}
}
