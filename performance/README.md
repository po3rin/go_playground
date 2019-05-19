# High Performance Go Workshop の要約

下記のページの自分なりの要約&翻訳になっています。非常に勉強になりました。最高です。
[High Performance Go Workshop](https://dave.cheney.net/high-performance-go-workshop/gopherchina-2019.html)

## Goでベンチマークをとる際に知っておくべきこと

```go
func Fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return Fib(n-1) + Fib(n-2)
	}
}
```

```go
func BenchmarkFib20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(20) // run the Fib function b.N times
	}
}
```

各ベンチマーク関数はb.N回反復されます。デフォルトではb.Nは1から始まりますが、ベンチマーク機能が1秒以内に完了した場合b.Nが増加してベンチマーク機能が再度実行されます。 ｂ．Ｎは近似シーケンスで増加していきます。下記のようにベンチマークが実行できます。

```
$ go test -bench=.  ./
goos: darwin
goarch: amd64
pkg: github.com/po3rin/performance
BenchmarkFib20-12   	  500000	     35934 ns/op
PASS
ok  	github.com/po3rin/performance	18.340s
```

### --cpu

先ほどのベンチマークの結果を見ると```-12```というsuffixがついています。これはこのテストを実行するために使用されたGOMAXPROCSの値です。この数はデフォルトで起動時にGoプロセスに見えるCPUの数になります。この値は、ベンチマークを実行するための値のリストを取る-cpuフラグで変更できます。-cpuの使い方を見ていきましょう。

```
$ go test -bench=. -cpu=1,2,4 ./
goos: darwin
goarch: amd64
pkg: github.com/po3rin/performance
BenchmarkFib20     	   50000	     36892 ns/op
BenchmarkFib20-2   	   50000	     36000 ns/op
BenchmarkFib20-4   	   50000	     35878 ns/op
BenchmarkFib20-8   	   50000	     35634 ns/op
PASS
ok  	github.com/po3rin/performance	8.687s
```

今回の例のベンチマークは完全にシーケンシャルな処理であるため、フラグによる結果への影響はほとんど無いのが確認できます。 チーム間でベンチマークを使い回す時はCPUの数がベンチマークに影響を与えないように注意が必要です。

### -benchtime

反復回数を増やす為に-benchtimeフラグを使用してベンチマーク時間を増やすことができます。

```
$ go test -bench=. -benchtime=10s ./
goos: darwin
goarch: amd64
pkg: github.com/po3rin/performance
BenchmarkFib20-12    	  500000	     35934 ns/op
PASS
ok  	github.com/po3rin/performance	18.340s
```

また Go1.12から-benchtimeフラグは反復回数を指定できます。

```
go test -bench=Fib20 -benchtime=20x ./
goos: darwin
goarch: amd64
pkg: github.com/po3rin/performance
BenchmarkFib20-12    	      20	     37894 ns/op
PASS
ok  	github.com/po3rin/performance	0.006s
```

### -count

ここで別のベンチマークを追加してみましょう。

```go
func BenchmarkFib1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(1) // run the Fib function b.N times
	}
}
```

ベンチマークをとってみます。

```
$ go test -bench=Fib1 ./
goos: darwin
goarch: amd64
pkg: github.com/po3rin/performance
BenchmarkFib1-12    	2000000000	         1.53 ns/op
PASS
ok  	github.com/po3rin/performance	3.219s
```

実は上の例のように数百万または数十億の反復で実行されるベンチマークが1マイクロ秒またはナノ秒の範囲の操作あたりの時間になる場合、ベンチマーク数は熱スケーリング、メモリ局所性、バックグラウンド処理、gcアクティビティなどのため不安定になります。1回の操作で10桁または1桁のナノ秒単位で測定される場合、命令の並べ替えやコード配置の相対的な影響がベンチマーク時間に影響を与えます。

よってこのような場合は-countフラグを使用して、ベンチマークを複数回実行することでベンチマークの分散も含めて確認するのが得策です。

```
go test -bench=Fib1 -count=10 ./
goos: darwin
goarch: amd64
pkg: github.com/po3rin/performance
BenchmarkFib20-12    	   50000	     36186 ns/op
BenchmarkFib20-12    	   50000	     36714 ns/op
BenchmarkFib20-12    	   50000	     35819 ns/op
BenchmarkFib20-12    	   50000	     36048 ns/op
BenchmarkFib20-12    	   50000	     36291 ns/op
BenchmarkFib20-12    	   50000	     36326 ns/op
BenchmarkFib20-12    	   50000	     36338 ns/op
BenchmarkFib20-12    	   50000	     36242 ns/op
BenchmarkFib20-12    	   50000	     36845 ns/op
BenchmarkFib20-12    	   50000	     36131 ns/op
PASS
ok  	github.com/po3rin/performance	18.340s
```

テストに適用するデフォルト設定を特定のパッケージに合わせて調整する必要があることがわかった場合は、ベンチマークを実行したいすべての人が同じ設定を使用できるように、Makefileでそれらの設定を体系化することが求められます。

## ベンチマークの安定性

前のセクションでは、ベンチマーク結果の分散に惑わされないために複数回ベンチマークを実行することを提案しました。電力管理、バックグラウンドプロセス、および熱管理の影響のため、これはどのベンチマークに対しても有効な手段です。一方でベンチマークの安定度を見るのに便利なツールがあります。Russ Coxによるbenchstatというツールを紹介します。

```
$ go get golang.org/x/perf/cmd/benchstat
```

benchstatは一連のベンチマークテストを実行して、それらがどれほど安定しているかを教えてくれます。まずはベンチマーク結果をファイルに出力します。

```
go test -bench=Fib20 -count=10 ./ | tee old.txt
goos: darwin
goarch: amd64
pkg: github.com/po3rin/performance
BenchmarkFib20-12    	   50000	     35823 ns/op
BenchmarkFib20-12    	   50000	     35655 ns/op
BenchmarkFib20-12    	   50000	     35932 ns/op
BenchmarkFib20-12    	   50000	     35763 ns/op
BenchmarkFib20-12    	   50000	     35861 ns/op
BenchmarkFib20-12    	   50000	     35723 ns/op
BenchmarkFib20-12    	   50000	     35681 ns/op
BenchmarkFib20-12    	   50000	     35836 ns/op
BenchmarkFib20-12    	   50000	     36004 ns/op
BenchmarkFib20-12    	   50000	     36248 ns/op
PASS
ok  	github.com/po3rin/performance	21.576s
```

```
$ benchstat old.txt
name      time/op
Fib20-12  35.9µs ± 1%
```

この結果は平均が```39.1µs```で```±1%```で結果が変動することを示しています。これはかなり安定していますね。

## ベンチマーク間の比較
ベンチマーク間のパフォーマンスの差を判断するのは面倒で須賀、benchstatはこの問題も解決します。

早速Fib()を改良したいのですが、コードを改良した後で、もう一度改良前のベンチマークを取りたい時が出てきたらどうしましょうか。実は```go test```には前回のベンチマーク結果を生成したバイナリを保存しておくことができる機能を提供する```-c```が存在します。

```
$ go test -c
```

改良前のバイナリ名は.testから.goldenに変更しておくと良いでしょう。

```
$ mv performance.test performance.golden
```

早速今のFib()関数を改善しましょう。以前の関数は、フィボナシ級数の0番目と1番目の数値に対してハードコーディングされた値を持っていました。その後、コードはそれ自身を再帰的に呼び出します。再帰のコストについては後ほどお話ししますが、特に私たちのアルゴリズムが指数関数的な時間を使うので、コストがかかると仮定します。 これを簡単に修正するには、フィボナッチ数列から別の番号をハードコーディングして、各再帰呼び出しの深さを1つ減らします。


```go
func Fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	default:
		return Fib(n-1) + Fib(n-2)
	}
}
```

これでコードを改良しました。変更前と後でどれほどパフォーマンス改善ができたのでしょうか。```benchstat```で確認しましょう。

```
$ go test -c
$ ./performance.golden -test.bench=. -test.count=10 > old.txt
$ ./performance.test -test.bench=. -test.count=10 > new.txt
$ benchstat old.txt new.txt
name      old time/op  new time/op  delta
Fib20-12  35.8µs ± 1%  22.0µs ± 2%  -38.52%  (p=0.000 n=10+9)
Fib1-12   1.55ns ± 2%  1.67ns ± 4%   +8.02%  (p=0.000 n=10+10)
```

改良後のベンチマーク結果の分散は3%なのでまあまあ信頼できる結果のようです。分散が大きいベンチマークを比較するときは間違った結論を導いてしまう可能性があるので注意してください。Fib(1)は改良は見られませんでしたが、Fib(20)では-38.45%の改良が確認できます。

nの値は、古いサンプルと新しいサンプルのうち、何個のデータを有効なデータであると見なしたかを表します。今回の結果を見ると```-count=10```を実行したにも関わらず、9つだけが報告されています。データの棄却率が10％以下であれば問題無いようですが、10％を超えると、セットアップが不安定になり、比較するサンプルが少なすぎる可能性があるようです。

p値が```0.05```を超えることはベンチマークが統計的に有意ではないことを意味します。p値については下記が詳しいです。
[統計学的検定のP値、統計学的に有意、有意差、有意水準とは何か？](http://toukei.link/basicstatistics/pvalue_and_significance/)

## ベンチマークコストの回避

ベンチマークの実行にセットアップコストがかかる場合があります。 b.ResetTimer()は、セットアップで発生した時間を回避するために使用します。

```go
func BenchmarkExpensive(b *testing.B) {
        boringAndExpensiveSetup()
        b.ResetTimer()
        for n := 0; n < b.N; n++ {
                // function under test
        }
}
```

ループの反復ごとにコストが高いセットアップがある場合は、b.StopTimer()およびb.StartTimer()を使用してベンチマークタイマーを一時停止します。

```go
func BenchmarkComplicated(b *testing.B) {
        for n := 0; n < b.N; n++ {
                b.StopTimer()
                complicatedSetup()
                b.StartTimer()
                // function under test
        }
}
```

## アロケーションの測定

アロケーションの数とサイズは、ベンチマーク時間と強く相関しています。アロケーションの数を記録するようにコード内で指示できます。

```go
func BenchmarkRead(b *testing.B) {
        b.ReportAllocs()
        for n := 0; n < b.N; n++ {
                // function under test
        }
}
```

下記はbufioパッケージのベンチマークを使った結果です。
アロケーションサイズと回数が表示されます。

```
$ go test -bench=. bufio

// ...

BenchmarkWriterEmpty-6            	  500000	      3180 ns/op	    4096 B/op	       1 allocs/op
BenchmarkWriterEmpty-6            	  500000	      2922 ns/op	    4096 B/op	       1 allocs/op
BenchmarkWriterEmpty-6            	  500000	      3092 ns/op	    4096 B/op	       1 allocs/op
BenchmarkWriterEmpty-6            	  500000	      3191 ns/op	    4096 B/op	       1 allocs/op
BenchmarkWriterEmpty-6            	  500000	      3015 ns/op	    4096 B/op	       1 allocs/op
BenchmarkWriterFlush-6            	100000000	        13.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkWriterFlush-6            	100000000	        14.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkWriterFlush-6            	100000000	        13.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkWriterFlush-6            	100000000	        13.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkWriterFlush-6            	100000000	        13.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkWriterFlush-6            	100000000	        13.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkWriterFlush-6            	100000000	        13.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkWriterFlush-6            	100000000	        13.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkWriterFlush-6            	100000000	        13.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkWriterFlush-6            	100000000	        13.9 ns/op	       0 B/op	       0 allocs/op
```

## コンパイラ最適化によるベンチマークのハマりポイント

下記のコードのベンチマークをとります。

```go
const m1 = 0x5555555555555555
const m2 = 0x3333333333333333
const m4 = 0x0f0f0f0f0f0f0f0f
const h01 = 0x0101010101010101

func popcnt(x uint64) uint64 {
	x -= (x >> 1) & m1
	x = (x & m2) + ((x >> 2) & m2)
	x = (x + (x >> 4)) & m4
	return (x * h01) >> 56
}

func BenchmarkPopcnt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcnt(uint64(i))
	}
}
```

```
$ go test -bench=Popcnt
goos: darwin
goarch: amd64
pkg: github.com/po3rin/performance
BenchmarkPopcnt-12    	2000000000	         0.25 ns/op
PASS
ok  	github.com/po3rin/performance	0.544s
```

この0.25秒は概ね1clock周期です。よってこの値はかなりおかしいです。

popcntはリーフ関数(他の関数は呼び出しをしない)になっています。。コンパイラはこの関数をインライン展開できます。 関数がインライン化されます。そしてpopcntは、どのグローバル変数の状態にも影響しません。したがって、呼び出し自体が排除されています。


ベンチマークを機能させるためにインライン化を無効にすることは非現実的であり、最適化をオンにしてベンチマークを撮ることが求められます。このベンチマークを修正するには、BenchmarkPopcntの本体がグローバル状態を変更しないことをコンパイラが証明できないようにする必要があります。下記は、コンパイラがループ本体を最適化できないようにするための推奨方法です。

```go
var Result uint64

func BenchmarkPopcnt(b *testing.B) {
	var r uint64
	for i := 0; i < b.N; i++ {
		r = popcnt(uint64(i))
	}
	Result = r
}
```

Resultはpublic変数なので、これをインポートしている他のパッケージがResultの値が時間とともに変化するのを見ることができないことをコンパイラは証明できません。したがって、この操作をコンパイラが最適化することはできません。
