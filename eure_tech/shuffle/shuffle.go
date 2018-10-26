package shuffle

import (
	"math/rand"
)

type Interface interface {
	Seed() int64
	Len() int
	Swap(int, int)
}

func Shuffle(data Interface) {
	rand.Seed(data.Seed())
	n := data.Len()
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data.Swap(i, j)
	}
}
