package safemap

import "sync"

type SafeMap struct {
	mu sync.Mutex
	m  map[string]interface{}
}

func (sm *SafeMap) Get(key string) (interface{}, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	val, ok := sm.m[key]
	return val, ok
}

func (sm *SafeMap) Set(key string, val interface{}) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.m[key] = val
}

func NewSyncMap() *SafeMap {
	return &SafeMap{m: make(map[string]interface{}, 5)}
}
