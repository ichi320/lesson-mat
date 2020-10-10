package main

import (
	"fmt"

	"github.com/ichi320/lesson-mat/001_matrix/utils"
	"gonum.org/v1/gonum/mat"
)

func main() {

	a := mat.NewVecDense(4, []float64{0, 1, 2, 3})
	fmt.Printf("%T\n", a)
	utils.MatPrint(a)

	data := []float64{9, 8, 7, 6, 5, 4}
	b := mat.NewVecDense(len(data), data)
	fmt.Printf("%T\n", b)
	utils.MatPrint(b)

}
