package main

import (
	"fmt"

	"github.com/ichi320/lesson-mat/001_matrix/utils"
	"gonum.org/v1/gonum/mat"
)

func main() {

	data := []float64{0, 1, 2, 3, 4, 5}
	A := mat.NewDense(2, 3, data)

	a := A.At(0, 1)
	fmt.Printf("%T %v\n\n", a, a)

	A.Set(0, 1, 888)
	utils.MatPrint(A)

}
