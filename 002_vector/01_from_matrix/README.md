# create Vector from Matrix
>	a := A.RowView(0)
これだと返り値がVector型になり、後続の処理に不都合な場合がある。

> _, c := A.Dims()  
> a := mat.NewVecDense(c, mat.Row(nil, 0, A))
これだと返り値が*VecDense型になる
