package cas

import (
	"fmt"
	"sync"
)

var m = sync.Mutex{}
var store = map[string]string{}

func compareAndSwap(key, nextVal, currentVal string) (bool, error) {
	m.Lock()
	defer m.Unlock()

	_, containsKey := store[key]
	if !containsKey {
		if len(currentVal) == 0 {
			store[key] = nextVal
			return true, nil
		}
		return false, fmt.Errorf("expectedval %s for key %s, but found empty", currentVal, key)
	}

	if store[key] == currentVal {
		store[key] = nextVal
		return true, nil
	}

	return false, nil
}
