package bloomfilter

import (
	"fmt"
	"math"
)

// BloomFilter is a slice of boolen.
type BloomFilter struct {
	BitFiled []bool
	K        int
	M        int
	N        int
}

func NewBloomFilter(size int, strNum int) (*BloomFilter, error) {
	if strNum == 0 {
		return nil, fmt.Errorf("strNum requires 1 or more")
	}
	return &BloomFilter{
		BitFiled: make([]bool, size),
		K:        int(math.Log(2) * float64(size) / float64(strNum)),
		M:        size,
	}, nil
}

func (b *BloomFilter) Add(substr string) {
	hs := getHash(substr, b.M, b.K)
	b.N++
	for _, v := range hs {
		b.BitFiled[v] = true
	}
}

func (b *BloomFilter) MightContain(s string) bool {
	hs := getHash(s, b.M, b.K)
	for _, v := range hs {
		if !b.BitFiled[v] {
			return false
		}
	}
	return true
}

func (b *BloomFilter) EstimateFPP() float64 {
	x := -float64(b.K) * float64(b.N) / float64(b.M)
	y := 1 - math.Exp(x)
	return math.Pow(y, float64(b.K))
}

const (
	prime1 = 53
	prime2 = 31
)

func getHash(s string, m int, k int) []int {
	h1 := hashStr(s, prime1, m)
	h2 := hashStr(s, prime2, m)

	list := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Println((h1 + i*h2) % m)
		list[i] = (h1 + i*h2) % m
	}
	return list
}
