package main

import (
	"fmt"

	"github.com/ichi320/lesson-mat/001_matrix/utils"
	"gonum.org/v1/gonum/mat"
)

func main() {
	A := mat.NewDense(2, 3, []float64{0, 1, 2, 3, 4, 5})

	a := A.RowView(0)
	fmt.Printf("%T - %v\n", a, a)
	utils.MatPrint(a)

	a = A.ColView(0)
	fmt.Printf("%T - %v\n", a, a)
	utils.MatPrint(a)

}
