package main

import (
	"fmt"

	"github.com/ichi320/lesson-mat/001_matrix/utils"
	"gonum.org/v1/gonum/mat"
)

func main() {
	x := mat.NewVecDense(5, []float64{0, 1, 2, 3, 4})
	a := x.AtVec(3)
	fmt.Printf("%T - %v\n", a, a)

	b := x.At(4, 0) // second number is always 0
	fmt.Printf("%T - %v\n", b, b)

	x.SetVec(0, 888)
	utils.MatPrint(x)

}
