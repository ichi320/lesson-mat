package main

import (
	"log"
	"math/rand"

	"github.com/ichi320/lesson-mat/001_matrix/utils"
	"gonum.org/v1/gonum/mat"
)

// MatMul is MatMul node
type MatMul struct {
	W  mat.Matrix
	B  mat.Vector
	X  mat.Matrix
	DW mat.Matrix
	DB mat.Vector
}

// Layer is interface for layers
type Layer interface {
	Forward(mat.Matrix) mat.Matrix
	Backword(mat.Matrix) mat.Matrix
}

func dot(x mat.Matrix, w mat.Matrix) mat.Matrix {
	n, _ := x.Dims()
	_, c := w.Dims()
	y := mat.NewDense(n, c, nil)
	y.Mul(x, w)
	return y
}

// ConstAdd 行列に定数を加算
func constAdd(x mat.Matrix, b mat.Vector) mat.Matrix {
	r, c := x.Dims()
	if c != b.Len() {
		log.Fatalln("Matrix rank error.")
	}
	var B *mat.Dense
	B = mat.NewDense(r, c, nil)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			B.Set(i, j, b.AtVec(j))
		}
	}

	m := mat.NewDense(r, c, nil)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			m.Set(i, j, x.At(i, j)+B.At(i, j))
		}
	}
	return m
}

func (layer *MatMul) Forward(x mat.Matrix) mat.Matrix {
	layer.X = x

	var y mat.Matrix
	y = dot(layer.X, layer.W)
	y = constAdd(y, layer.B)
	return y
}

func main() {

	var layer MatMul

	n := 4 // data count
	r := 2 // input vector demension
	var data []float64
	data = []float64{1, 1, 0, 0, 0, 1, 1, 0}
	X := mat.NewDense(n, r, data)

	c := 2 // output vector demension
	data = make([]float64, r*c)
	for i, _ := range data {
		data[i] = rand.NormFloat64()
		// data[i] = float64(i) // for test
	}
	layer.W = mat.NewDense(r, c, data)

	data = make([]float64, c)
	for i := 0; i < c; i++ {
		data[i] = float64(1)
	}
	layer.B = mat.NewVecDense(c, data)

	y := layer.Forward(X)

	// data = []float64{1, 0, 1, 0, 0, 1, 0, 1}
	// y := mat.NewDense(n, c, data)

	utils.MatPrint(layer.X)
	utils.MatPrintRound(layer.W)
	utils.MatPrintRound(y)

}
