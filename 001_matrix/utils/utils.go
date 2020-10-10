package utils

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// MatPrint print out Matrix
func MatPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
	fmt.Println("")
}

// MatPrintRound print out Matrix rounded
func MatPrintRound(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%.3f\n", fa)
	fmt.Println("")
}
