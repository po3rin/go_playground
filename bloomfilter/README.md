# bloomfilter

bloomfilter impliments bloom filter using polynomial rolling hash function.

https://en.wikipedia.org/wiki/Bloom_filter

```go
package main

import (
	"log"

	"github.com/po3rin/bloomfilter"
)

func main() {
	s := "reg"
	strList := []string{
		"reg",
		"fw",
		"eb",
	}

	strNum := len(strList)
	filter, _ := bloomfilter.NewBloomFilter(1000, strNum)

	for _, s := range strList {
		filter.Add(s)
	}

	ffp := filter.EstimateFPP()
	log.Printf("ffp: %e", ffp)

	if filter.MightContain(s) {
		log.Print("maybe contain!!")
		return
	}
	log.Print("never contain!!")
}
```