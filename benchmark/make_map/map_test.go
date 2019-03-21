package map_test // import "make_map"

import (
	"log"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

var sl = []string{
	"a=a",
	"b=a",
	"c=a",
	"d=a",
	"e=a",
	"f=a",
	"g=a",
	"h=a",
	"i=a",
	"j=a",
	"l=a",
	"m=a",
	"n=a",
	"o=a",
	"p=a",
	"q=a",
	"r=a",
}

func AttrMapWithoutCaps(sl []string) (map[string]string, error) {
	m := map[string]string{}
	for _, v := range sl {
		parts := strings.SplitN(v, "=", 2)
		if len(parts) != 2 {
			return nil, errors.Errorf("invalid value %s", v)
		}
		m[parts[0]] = parts[1]
	}
	return m, nil
}

func AttrMapWithCaps(sl []string) (map[string]string, error) {
	m := make(map[string]string, len(sl))
	for _, v := range sl {
		parts := strings.SplitN(v, "=", 2)
		if len(parts) != 2 {
			return nil, errors.Errorf("invalid value %s", v)
		}
		m[parts[0]] = parts[1]
	}
	return m, nil
}

func BenchmarkAttrMapWithoutCaps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := AttrMapWithoutCaps(sl)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkAttrMapWithCaps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := AttrMapWithCaps(sl)
		if err != nil {
			log.Fatal(err)
		}
	}
}
