package main

import (
	"fmt"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

func main() {
	r, c := 3, 7

	data := []float64{}
	var value float64
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			value = rand.Float64()
			data = append(data, value)
		}
	}
	A := mat.NewDense(r, c, data)
	fA := mat.Formatted(A, mat.Prefix("    "), mat.Squeeze())
	fmt.Printf("A = %.3f", fA)

	fmt.Println("")

	x := mat.NewVecDense(c, []float64{1, 0, 0, 0, 0, 0, 0})
	fmt.Printf("x: %T - %v\n\n", x, x)

	var y mat.VecDense
	y.MulVec(A, x)
	fy := mat.Formatted(&y, mat.Prefix("    "), mat.Squeeze())
	fmt.Printf("y = %.3f", fy)
}
