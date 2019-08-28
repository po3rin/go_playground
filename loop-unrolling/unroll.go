package unroll

func ContainsUint64(ids []uint64, id uint64) bool {
	for i := 0; i < len(ids); i++ {
		if ids[i] == id {
			return true
		}
	}
	return false
}

func ContainsUint64WithMap(ids []uint64, id uint64) bool {
	m := make(map[uint64]struct{}, len(ids))
	for _, n := range ids {
		m[n] = struct{}{}
	}

	if _, ok := m[id]; !ok {
		return false
	}
	return true
}

func ContainsUint64Unroll2(ids []uint64, id uint64) bool {
	var i int
	for i = len(ids) - 1; i >= 1; i -= 2 {
		if ids[i] == id || ids[i-1] == id {
			return true
		}
	}
	for ; i >= 0; i-- {
		if ids[i] == id {
			return true
		}
	}
	return false
}

func ContainsUint64Unroll4(ids []uint64, id uint64) bool {
	var i int
	for i = len(ids) - 1; i >= 3; i -= 4 {
		if ids[i] == id ||
			ids[i-1] == id ||
			ids[i-2] == id ||
			ids[i-3] == id {
			return true
		}
	}
	for ; i >= 0; i-- {
		if ids[i] == id {
			return true
		}
	}
	return false
}

func ContainsUint64Unroll8(ids []uint64, id uint64) bool {
	var i int
	for i = len(ids) - 1; i >= 7; i -= 8 {
		if ids[i] == id ||
			ids[i-1] == id ||
			ids[i-2] == id ||
			ids[i-3] == id ||
			ids[i-4] == id ||
			ids[i-5] == id ||
			ids[i-6] == id ||
			ids[i-7] == id {
			return true
		}
	}
	for ; i >= 0; i-- {
		if ids[i] == id {
			return true
		}
	}
	return false
}

func ContainsUint64Unroll8WithBoundsCheck(ids []uint64, id uint64) bool {
	var i int
	for ; i < len(ids); i += 8 {
		if ids[i] == id ||
			ids[i+1] == id ||
			ids[i+2] == id ||
			ids[i+3] == id ||
			ids[i+4] == id ||
			ids[i+5] == id ||
			ids[i+6] == id ||
			ids[i+7] == id {
			return true
		}
	}
	for ; i < len(ids); i++ {
		if ids[i] == id {
			return true
		}
	}
	return false
}
