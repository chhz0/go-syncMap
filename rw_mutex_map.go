package gosyncmap

import "sync"

// RWMutexMap 使用读写锁实现的线程安全map
type RWMutexMap struct {
	mu *sync.RWMutex
	mp map[string]any
}

// Clear implements SyncMap.
func (m *RWMutexMap) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.mp = make(map[string]any)
}

// Delete implements SyncMap.
func (m *RWMutexMap) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.mp, key)
}

// Load implements SyncMap.
func (m *RWMutexMap) Load(key string) (value any, ok bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, ok = m.mp[key]
	return value, ok
}

// Store implements SyncMap.
func (m *RWMutexMap) Store(key string, value any) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.mp[key] = value
}

func NewRWMutexMap() *RWMutexMap {
	return &RWMutexMap{
		mu: new(sync.RWMutex),
		mp: make(map[string]any),
	}
}

var _ SyncMap = (*RWMutexMap)(nil)
