package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := x / 2
	for {
		p := z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		if p == z {
			return z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(49))
}
