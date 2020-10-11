package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewDense(2, 2, []float64{
		4, 0,
		0, 4,
	})
	b := mat.NewDense(2, 3, []float64{
		4, 0, 0,
		0, 0, 4,
	})

	var c mat.Dense
	c.Mul(a, b)

	fc := mat.Formatted(&c, mat.Prefix("    "), mat.Squeeze())
	fmt.Printf("c = %v", fc)
}
