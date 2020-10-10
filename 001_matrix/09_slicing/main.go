package main

import (
	"fmt"

	"github.com/ichi320/lesson-mat/001_matrix/utils"
	"gonum.org/v1/gonum/mat"
)

func main() {
	var data []float64
	var A *mat.Dense
	var B mat.Matrix // なぜか*Dense型で受けるとこける

	data = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	A = mat.NewDense(3, 4, data)
	B = A.Slice(0, 2, 1, 3)

	fmt.Printf("%T\n", B) // *Dense型なのに...
	utils.MatPrint(B)

	var C *mat.Dense
	var D, E mat.Dense
	data = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	C = mat.NewDense(3, 4, data)
	D.Stack(A, C)
	fmt.Println("stack")
	utils.MatPrint(&D)
	E.Augment(A, C)
	fmt.Println("Augment")
	utils.MatPrint(&E)

}
