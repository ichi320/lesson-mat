package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewVecDense(3, []float64{1, 2, 3})
	b := mat.NewVecDense(3, []float64{9, 8, 7})
	c := mat.Dot(a, b)
	fmt.Printf("%T\n%v\n", c, c)

}
