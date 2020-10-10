package main

import (
	"github.com/ichi320/lesson-mat/001_matrix/utils"
	"gonum.org/v1/gonum/mat"
)

func main() {
	data := []float64{0, 1, 2, 3, 4, 5}
	A := mat.NewDense(2, 3, data)
	utils.MatPrint(A)

	row := []float64{53, 52, 51}
	A.SetRow(0, row)

	utils.MatPrint(A)

	col := []float64{98, 97}
	A.SetCol(1, col)

	utils.MatPrint(A)

}
