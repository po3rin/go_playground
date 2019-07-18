# Go + gonum を使った行列計算まとめ

こんにちは！！Goを愛する皆様におかれましてはビッグデータ解析やニューラルネットワークの実装をGoでやりたいですよね！！そうすると嫌が応にも行列の計算が発生します。そこで今回は Go で 行列計算をやる方法を紹介します。

## gonum/mat

Go で 行列を扱う際には gonum パッケージが鉄板でしょう。gonum は行列だけでなく数値および科学的アルゴリズムの作成を支援するパッケージです。数値計算はこのパッケージに頼りましょう。

<a href="https://github.com/gonum/gonum"><img src="src/gonum.png" width="460px"></a>

## 行列の基本

まずは行列の作り方から

```go
package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
    A := mat.NewDense(3, 4, x)
    // ⎡1  2  3  4⎤
    // ⎢5  6  7  8⎥
    // ⎣9 10 11 12⎦
}
```

NewDense は *Dense を生成します。

```go
type Dense struct {
	mat blas64.General

	capRows, capCols int
}
```

Dense の構造は NewDence の中身を見れば分かりやすいです。内部では単純に[]float64のデータや行数、列数、ストライドを保持しているだけです。

```go
func NewDense(r, c int, data []float64) *Dense {
	if r < 0 || c < 0 {
		panic("mat: negative dimension")
	}
	if data != nil && r*c != len(data) {
		panic(ErrShape)
    }
    // nil なら 0 で初期化される
	if data == nil {
		data = make([]float64, r*c)
	}
	return &Dense{
		mat: blas64.General{
			Rows:   r,
			Cols:   c,
			Stride: c,
			Data:   data,
		},
		capRows: r,
		capCols: c,
	}
}
```

*Dense が Matrix インターフェースを実装しています。実装の中では基本的にこの Matrix 型で引き渡していくと便利です。

```go
// Matrix is the basic matrix interface type.
type Matrix interface {
	// Dims returns the dimensions of a Matrix.
	Dims() (r, c int)

	// At returns the value of a matrix element at row i, column j.
	// It will panic if i or j are out of bounds for the matrix.
	At(i, j int) float64

	// T returns the transpose of the Matrix. Whether T returns a copy of the
	// underlying data is implementation dependent.
	// This method may be implemented using the Transpose type, which
	// provides an implicit matrix transpose.
	T() Matrix
}
```

また、行列のデバッグのために整形して標準出力できるようにしておくと便利です。

```go
func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}
```

これの関数を使えばフォーマットされた行列を出力できます。

```go
func main() {
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
    A := mat.NewDense(3, 4, x)
    matPrint(A)
    // ⎡1   2   3   4⎤
    // ⎢5   6   7   8⎥
    // ⎣9  10  11  12⎦
}
```

NewDense の内部実装をみた通り、全要素0で初期化したい場合は第３引数に```nil```を与えます。

```go
func main() {
    A := mat.NewDense(3, 4, nil)
    matPrint(A)
    // ⎡0  0  0  0⎤
    // ⎢0  0  0  0⎥
    // ⎣0  0  0  0⎦
}
```

以降 A を作るコードは基本的に省略します。

## 要素の変更 or 取得

これも簡単でAtメソッドでアクセスできます。また新しい要素のセットはSetメソッドで行えます。

```go
func main() {
    // 行列作成は省略

    a := A.At(0, 2)
    println("A[0, 2]: ", a)
    // A[0, 2]:  +3.000000e+000

	A.Set(0, 2, -1.5)
    matPrint(A)
    // ⎡1   2  -1.5   4⎤
    // ⎢5   6     7   8⎥
    // ⎣9  10    11  12⎦
}
```

## 行だけ列だけを Vector として取り出す

RowView と ColView で行列の中から指定した行 or 列を Vector として取得できます。取得している型は Vector インターフェースです。

```go
type Vector interface {
	Matrix
	AtVec(int) float64
	Len() int
}
```

RowView と ColView の Example です。

```go
func main() {
    // 行列作成は省略

    matPrint(A.RowView(1))
    // ⎡5⎤
    // ⎢6⎥
    // ⎢7⎥
    // ⎣8⎦

    matPrint(A.ColView(0))
    // ⎡1⎤
    // ⎢5⎥
    // ⎣9⎦
}
```

## 指定した行や列を変更する

指定した行や列をまとめて更新できます。SetRow, SetCol で行ます。

```go
func main() {
    // 行列作成は省略

	row := []float64{10, 9, 8, 7}
    A.SetRow(0, row)
    matPrint(A)
    // ⎡10   9   8   7⎤
    // ⎢ 5   6   7   8⎥
    // ⎣ 9  10  11  12⎦

	col := []float64{3, 2, 1}
    A.SetCol(0, col)
    matPrint(A)
    // ⎡3   9   8   7⎤
    // ⎢2   6   7   8⎥
    // ⎣1  10  11  12⎦
}
```

## 要素同しの足し引き

ここは少し直感的ではないですが空の行列を作ってそこに計算結果を格納するという方法をとります。ゼロ値で初期化した mat.Dence に対して計算結果を込めます。計算結果を格納する先の行列のサイズが計算結果と合わないとpanicするので注意してください。

```go
func main() {
    // 行列作成は省略

    B := mat.NewDense(3, 4, nil)
	B.Add(A, A)
    matPrint(&B)
    // ⎡ 2   4   6   8⎤
    // ⎢10  12  14  16⎥
    // ⎣18  20  22  24⎦

	C := mat.NewDense(3, 4, nil)
	C.Sub(A, A)
    matPrint(&C)
    // ⎡0  0  0  0⎤
    // ⎢0  0  0  0⎥
    // ⎣0  0  0  0⎦
}
```

毎回ゼロ値で初期化した mat.Dence を準備するのも面倒です。計算前の行列がいらないなら下記のようにも書けます。

```go
func main() {
    // 行列作成は省略

    A.Add(A, A)
	matPrint(A)

	A.Sub(A, A)
	matPrint(A)
}
```

ただ、いくらか関数の形が直感的ではないので独自の関数を用意しておくと便利かもしれません。

```go
func Add(a mat.Matrix, b mat.Matrix) mat.Matrix {
    // TODO got matrix size from args.
	var B mat.Dense
	B.Add(a, b)
	return &B
}
```

## 各要素の定数倍

各要素の定数倍も簡単です。Scaleメソッドを使います。

```go
func main() {
    // 行列作成は省略

    C = mat.NewDense(3, 4, nil)
	C.Scale(2, A)
    matPrint(&C)
    // ⎡ 2   4   6   8⎤
    // ⎢10  12  14  16⎥
    // ⎣18  20  22  24⎦
}
```

## 転置行列

転置行列は T メソッドでいけます。

```go
func main() {
    // 行列作成は省略

    B := A.T()
    matPrint(B)
    // ⎡1  5   9⎤
    // ⎢2  6  10⎥
    // ⎢3  7  11⎥
    // ⎣4  8  12⎦
}
```

## 逆行列

逆行列は Inverse メソッドです。

```go
func main() {
    A = mat.NewDense(2, 2, []float64{3, 5, 1, 2})
    // ⎡3  5⎤
    // ⎣1  2⎦

    B = mat.NewDense(2, 2, nil)
    err := B.Inverse(A)
    if err != nil {
        log.Fatal("failed to create inverse matrix")
    }
    matPrint(&B)
    // ⎡  1.999999999999999  -4.999999999999997⎤
    // ⎣-0.9999999999999996  2.9999999999999987⎦
}
```

概ね正しく計算できていますが、精度を要求されると厳しいかもしれません。

## 行列の内積

行列の内積はProductメソッドで行ます。(英語で内積は inner product )

```go
func main() {
    // 行列作成は省略

    C = mat.NewDense(3, 3, nil)
    C.Product(A, A.T())
    matPrint(&C)
    // ⎡ 30   70  110⎤
    // ⎢ 70  174  278⎥
    // ⎣110  278  446⎦
}
```

当然、掛け合わせる行数と列数が同じではなければいけません。dimension mismatch というパニックがおきますので注意してください。

## 行列のスライシング

スライシングは行列から指定の箇所を行列として抽出する操作です。

```go
func main() {
    // 行列作成は省略

    S := A.Slice(0, 3, 0, 3)
    matPrint(S)
    // ⎡1   2   3⎤
    // ⎢5   6   7⎥
    // ⎣9  10  11⎦
}
```

3 * 4 の行列から指定した部分だけを抽出しています。

## 各要素に任意の操作を実行する

Applyを使います。第一引数にやりたい操作の関数を渡してあげるだけです。今回は例として要素の値に行番号、列番号を足す処理を定義しています。

```go
func main() {
    // 行列作成は省略

    // 要素ごとに適用する関数
    sumOfIndices := func(i, j int, v float64) float64 {
		return float64(i+j) + v
    }

	var B mat.Dense
	B.Apply(sumOfIndices, A)
    matPrint(&B)
    // ⎡ 1   3   5   7⎤
    // ⎢ 6   8  10  12⎥
    // ⎣11  13  15  17⎦
```

## まとめ

これで一通りの処理はできるはずです。また気が向いたら更新していきます。Go で行列計算やっていきましょう！
