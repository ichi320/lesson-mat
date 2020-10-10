package main

import (
	"fmt"

	"github.com/ichi320/lesson-mat/001_matrix/utils"
	"gonum.org/v1/gonum/mat"
)

func main() {
	var data []float64
	data = []float64{0, 1, 2, 3, 4, 5}
	A := mat.NewDense(2, 3, data)
	fmt.Println("A")
	utils.MatPrint(A)
	data = []float64{10, 11, 12, 13, 14, 15}
	B := mat.NewDense(2, 3, data)
	fmt.Println("B")
	utils.MatPrint(B)

	C := mat.NewDense(2, 3, nil)
	fmt.Println("addition")
	C.Add(A, B)
	utils.MatPrint(C)

	fmt.Println("subtraction")
	C.Sub(B, A)
	utils.MatPrint(C)

	fmt.Println("scalar times")
	C.Scale(2, A)
	utils.MatPrint(C)
}
