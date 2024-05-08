package goplay

import (
	"math/rand"
	"sync"
)

// https://goplay.tools/snippet/TwHjL0lawUw
type RegularMap struct {
	sync.RWMutex
	internal map[string]int
}

func NewRegularMap() *RegularMap {
	return &RegularMap{
		internal: make(map[string]int),
	}
}

func (rm *RegularMap) Load(key string) (value int, ok bool) {
	rm.RLock()
	defer rm.RUnlock()
	result, ok := rm.internal[key]
	return result, ok
}

func (rm *RegularMap) Delete(key string) {
	rm.Lock()
	defer rm.Unlock()
	delete(rm.internal, key)
}

func (rm *RegularMap) Store(key string, value int) {
	rm.Lock()
	defer rm.Unlock()
	rm.internal[key] = value
}

func useWrappedMap() {
	key := "key"
	var wg sync.WaitGroup
	commonMap := NewRegularMap()
	mapSetter := func(wg *sync.WaitGroup) {
		defer wg.Done()
		commonMap.Store(key, rand.Intn(100))
	}
	// mapReader := func() {
	// 	fmt.Println(commonMap[key])
	// }

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go mapSetter(&wg)
		// go mapReader()
	}
	wg.Wait()
}
