# Go + gonum を使った行列計算まとめ

こんにちは！！Goを愛する皆様におかれましてはビッグデータ解析やニューラルネットワークの実装をGoでやりたいですよね！！そうすると嫌が応にも行列の計算が発生します。そこで今回は Go で 行列計算をやる方法を紹介します。

## gonum/mat

Go で 行列を扱う際には gonum パッケージが鉄板でしょう。gonum は行列だけでなく数値および科学的アルゴリズムの作成を支援するパッケージです。数値計算はこのパッケージに頼りましょう。

<a href="https://github.com/gonum/gonum"><img src="src/gonum.png" width="460px"></a>

## 行列の作成

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

速攻ですね。これで 3×4 の行列の完成です。行列をデバッグのために綺麗に標準出力できるようにしましょう。

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
}
```

実行してみましょう。

```bash
$ go run main.go
⎡1   2   3   4⎤
⎢5   6   7   8⎥
⎣9  10  11  12⎦
```

良いですね。ちなみに全要素0で初期化したい場合は第３引数に```nil```を与えます。

```go
func main() {
    A := mat.NewDense(3, 4, nil)
    // ⎡0  0  0  0⎤
    // ⎢0  0  0  0⎥
    // ⎣0  0  0  0⎦
}
```

以降 A を作るコードは省略します。

## 要素の変更 or 取得

これも簡単でAtメソッドでアクセスできます。また新しい要素のセットはSetメソッドで行えます。

```go
func main() {
    // 行列作成は省略

    a := A.At(0, 2)
	println("A[0, 2]: ", a)
	A.Set(0, 2, -1.5)
	matPrint(A)
}
```

実行するとうまく動いています。

```bash
$ go run main.go
A[0, 2]:  +3.000000e+000
⎡1   2  -1.5   4⎤
⎢5   6     7   8⎥
⎣9  10    11  12⎦
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

    println("Row 1 of A:")
	matPrint(A.RowView(1))

	println("Column 0 of A:")
    matPrint(A.ColView(0))
}
```

実行すると確かに狙ったベクトルを取得できています。

```bash
$ go run main.go
Row 1 of A:
⎡5⎤
⎢6⎥
⎢7⎥
⎣8⎦
Column 0 of A:
⎡1⎤
⎢5⎥
⎣9⎦
```

## 指定した行や列を変更する

指定した行や列をまとめて更新できます。SetRow, SetCol で行ます。

```go
func main() {
    // 行列作成は省略

	row := []float64{10, 9, 8, 7}
    A.SetRow(0, row)
    println("Updated A:")
	matPrint(A)

	col := []float64{3, 2, 1}
    A.SetCol(0, col)
    println("Updated A:")
    matPrint(A)
}
```

```bash
$ go run main.go
Updated A:
⎡10   9   8   7⎤
⎢ 5   6   7   8⎥
⎣ 9  10  11  12⎦
Updated A:
⎡3   9   8   7⎤
⎢2   6   7   8⎥
⎣1  10  11  12⎦
```

## 要素同しの足し引き

ここは少し直感的ではないですが空の行列を作ってそこに計算結果を格納するという方法をとります。ゼロ値で初期化した mat.Dence に対して計算結果を込めます。

```go
func main() {
    // 行列作成は省略

    var B mat.Dense
	B.Add(A, A)
	println("A + A:")
	matPrint(&B)

	var C mat.Dense
	C.Sub(A, A)
	println("A - A:")
	matPrint(&C)
}
```

実行すると要素ごとの足し引きができています。

```bash
$ go run main.go
A + A:
⎡ 2   4   6   8⎤
⎢10  12  14  16⎥
⎣18  20  22  24⎦
A - A:
⎡0  0  0  0⎤
⎢0  0  0  0⎥
⎣0  0  0  0⎦
```

毎回ゼロ値で初期化した mat.Dence を準備するのも面倒です。計算前の行列がいらないなら下記のようにも書けます。

```go
func main() {
    // 行列作成は省略

    A.Add(A, A)
	println("A + A:")
	matPrint(A)

	// var C mat.Dense
	A.Sub(A, A)
	println("A - A:")
	matPrint(A)
}
```

ただ、いくらか関数の形が直感的ではないので独自の関数を用意しておくと便利かもしれません。

```go
func Add(a mat.Matrix, b mat.Matrix) mat.Matrix {
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

    var C mat.Dense
	C.Scale(2, A)
	println("2 * A:")
	matPrint(&C)
}
```

こうなります。

```bash
$ go run main.go
2 * A:
⎡ 2   4   6   8⎤
⎢10  12  14  16⎥
⎣18  20  22  24⎦
```

## 転置行列

転置行列は T メソッドでいけます。

```go
func main() {
    // 行列作成は省略

    B := A.T()
	matPrint(B)
}
```

行列が転置しています。

```bash
$ go run main.go
⎡1  5   9⎤
⎢2  6  10⎥
⎢3  7  11⎥
⎣4  8  12⎦
```

## 行列の積

行列の積はProductメソッドで行ます。

```go
func main() {
    // 行列作成は省略

    var C mat.Dense
    C.Product(A, A.T())
    println("A * A: ")
	matPrint(&C)
}
```

```bash
$ go run main,.go
A * A:
⎡ 30   70  110⎤
⎢ 70  174  278⎥
⎣110  278  446⎦
```

当然、掛け合わせる行数と列数が同じではなければいけません。dimension mismatch というパニックがおきますので注意してください。

