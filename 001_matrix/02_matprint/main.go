package main

import (
	"github.com/ichi320/lesson-mat/001_matrix/utils"
	"gonum.org/v1/gonum/mat"
)

func main() {
	A := mat.NewDense(2, 3, nil)
	utils.MatPrintRound(A)
}
