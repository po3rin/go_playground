# zap

from https://github.com/uber-go/zap

Blazing fast, structured, leveled logging in Go.

Log a message and 10 fields:

| Package | Time | Objects Allocated |
| :--- | :---: | :---: |
| :zap: zap | 3131 ns/op | 5 allocs/op |
| :zap: zap (sugared) | 4173 ns/op | 21 allocs/op |
| zerolog | 16154 ns/op | 90 allocs/op |
| lion | 16341 ns/op | 111 allocs/op |
| go-kit | 17049 ns/op | 126 allocs/op |
| logrus | 23662 ns/op | 142 allocs/op |
| log15 | 36351 ns/op | 149 allocs/op |
| apex/log | 42530 ns/op | 126 allocs/op |

# How to adopt log level to stackdriver

https://tomokazu-kozuma.com/how-to-make-golangs-logger-zap-correspond-to-stackdriver/

# explain zap in ja
https://qiita.com/emonuh/items/28dbee9bf2fe51d28153