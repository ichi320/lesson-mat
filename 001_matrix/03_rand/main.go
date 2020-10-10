package main

import (
	"math/rand"

	"github.com/ichi320/lesson-mat/001_matrix/utils"
	"gonum.org/v1/gonum/mat"
)

func main() {
	r := 2
	c := 3

	data := make([]float64, r*c)
	for i, _ := range data {
		data[i] = rand.NormFloat64()
	}

	A := mat.NewDense(r, c, data)
	utils.MatPrintRound(A)
}
