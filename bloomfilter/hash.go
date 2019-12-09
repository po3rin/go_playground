package bloomfilter

import "hash"

func hashStr(sep string, prime int, m int) int {
	hash := uint32(0)
	p := uint32(prime)
	for i := 0; i < len(sep); i++ {
		hash = (hash*p + uint32(sep[i])) % uint32(m)
	}

	return int(hash)
}

var h hash.Hash
