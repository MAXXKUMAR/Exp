package cache

import (
	"sync"
)

var modeCountCache = make(map[string]map[string]int)
var mu sync.Mutex

func GetFromCache(areaCode string) (map[string]int, bool) {
	mu.Lock()
	defer mu.Unlock()
	data, ok := modeCountCache[areaCode]
	return data, ok
}

func StoreInCache(areaCode string, data map[string]int) {
	mu.Lock()
	defer mu.Unlock()
	modeCountCache[areaCode] = data
}
