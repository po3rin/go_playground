# create original slice

```go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```