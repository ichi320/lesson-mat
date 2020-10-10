package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	var A *mat.Dense
	A = mat.NewDense(2, 3, nil)
	fmt.Printf("%T\n%v\n", A, A)

	fmt.Println("")

	var data []float64 = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9} // int is error
	B := mat.NewDense(3, 3, data)
	fmt.Printf("%T\n%v\n", B, B)

}
