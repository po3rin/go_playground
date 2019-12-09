package main

import (
	"fmt"
	"math"
)

const primeRK = 6

func main() {
	c := "karp"
	fmt.Println(float64(c[0])*math.Pow(primeRK, 3) + float64(c[1])*math.Pow(primeRK, 2) + float64(c[2])*math.Pow(primeRK, 1) + float64(c[3])*math.Pow(primeRK, 0))
	// fmt.Println(int(float64(c[0])*math.Pow(primeRK, 1) + float64(c[1])*math.Pow(primeRK, 0)))
	h, p := hashStr(c)
	fmt.Println(h)
	fmt.Println(p)
}

func hashStr(sep string) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint32(sep[i])
	}
	var pow, sq uint32 = 1, primeRK
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 {
			fmt.Println("in if")
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}
