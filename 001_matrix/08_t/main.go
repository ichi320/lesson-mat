package main

import (
	"fmt"
	"log"

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

	fmt.Println("Tranpose")
	C := A.T()
	utils.MatPrint(C)

	data = []float64{1, 2, 3, 4}
	Din := mat.NewDense(2, 2, data)
	fmt.Println("Din")
	utils.MatPrint(Din)
	var Dout mat.Dense // なぜかポインタ型で受けると計算がこける
	err := Dout.Inverse(Din)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Inverse")
	utils.MatPrintRound(&Dout)
	fmt.Printf("Din: %T  Dout: %T\n\n", Din, Dout)

	var E mat.Dense // なぜかポインタ型で受けると計算がこける
	E.Product(A, A.T())
	fmt.Println("Product")
	utils.MatPrint(&E)

	F := mat.NewDense(2, 3, nil)
	F.MulElem(A, B)
	fmt.Println("Multiplication")
	utils.MatPrint(F)

}
