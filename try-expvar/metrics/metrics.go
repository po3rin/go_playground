package metrics

import (
	"expvar"
)

var (
	Map = expvar.NewMap("myapp") // expvar.NewXxx() は直接公開するメトリクスを作成する。

	Conns = new(expvar.Int) // Map に入れる値などは expvar.NewInt() ではなく new(expvar.Int) で作る

	MessageRecv = new(expvar.Int)
	MessageSent = new(expvar.Int)
)

func init() {
	// init() で Map に 値を登録していく
	Map.Set("conns", Conns)
	Map.Set("msg_recv", MessageRecv)
	Map.Set("msg_sent", MessageSent)
}
