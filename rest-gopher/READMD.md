# Go言語でロジックの差し替えをインターフェースにするか関数にするか

## 基本的な判断基準
- interface を満たすメソッドが、型が持つ feild に依存する場合は interface
- それ以外は関数

```go
package main

import "fmt"

func main() {
	Do(NewHelloFunc())
}
func Do(f DoFunc) {
    // any process ...
	f()
    // any process ...
}

type DoFunc func()

func NewHelloFunc() DoFunc {
	f := func() {
		fmt.Println("hello")
	}
	return DoFunc(f)
}
```

```go
package main

import "fmt"

func main() {
    h := &Hello{}
    Do(h)
}

func Do(f DoFunction) {
    // any process
    f.Exec()
    // any process
}

type DoFunction interface {
    Exec()
}

type Hello struct {}

func (h *Hello) Exec() {
    fmt.Println("hello")
}
```

## 複数の振る舞い定義

### interface の場合

```go
type DBFunc interface {
    Begin()
    Commit()
    Rollback()
}
```

### 関数の場合

```go
type DBFunc struct {
    Begin Begin
    Commit Commit
    Rollback Rollback
}
type Begin func()
type Commit func()
type Rollback func()
```