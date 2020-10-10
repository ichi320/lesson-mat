package main

import (
	"github.com/ichi320/lesson-mat/001_matrix/utils"
	"gonum.org/v1/gonum/mat"
)

func main() {
	A := mat.NewDense(2, 3, []float64{0, 1, 2, 3, 4, 5})

	somefunc := func(i, j int, v float64) float64 {
		return float64(i+j) + v
	}

	var K *mat.Dense
	r, c := A.Dims()
	K = mat.NewDense(r, c, nil) // 空箱を作る必要あり

	K.Apply(somefunc, A)
	utils.MatPrint(K)
}
