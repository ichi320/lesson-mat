package main

import (
	"fmt"

	"github.com/ichi320/lesson-mat/001_matrix/utils"
	"gonum.org/v1/gonum/mat"
)

func main() {
	A := mat.NewDense(2, 3, nil)
	r, c := A.Dims()
	fmt.Println("row cols: ", r, c)

	A.Set(0, 0, 888)
	utils.MatPrintRound(A)
}
