package map_test // import "make_map"

import (
	"testing"
)

var tt = map[string]interface{}{
	"a": 23,
	"b": "wwwww",
	"c": "wewewewew",
	"d": 3344,
	"e": 23,
	"f": "wwwww",
	"g": "wewewewew",
	"h": 3344,
	"i": 23,
	"j": "wwwww",
	"k": "wewewewew",
	"l": 3344,
	"m": 23,
	"n": "wwwww",
	"o": "wewewewew",
	"p": 3344,
	"q": 23,
	"r": "wwwww",
	"s": "wewewewew",
	"t": 3344,
}

func BenchmarkMapLen0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doc := make(map[string]interface{})
		for key, v := range tt {
			doc[key] = v
		}
	}
}

func BenchmarkMapWithLen10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doc := make(map[string]interface{}, 10)
		for key, v := range tt {
			doc[key] = v
		}
	}
}

func BenchmarkMapWithLen20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doc := make(map[string]interface{}, 20)
		for key, v := range tt {
			doc[key] = v
		}
	}
}

func BenchmarkMapWithLen30(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doc := make(map[string]interface{}, 30)
		for key, v := range tt {
			doc[key] = v
		}
	}
}
