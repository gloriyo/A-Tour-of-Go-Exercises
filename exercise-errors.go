package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := x / 2
	for {
		p := z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		// stop when value stops changing
		if math.Abs(p-z) < 1e-15 {
			return z, nil
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
