package common

import (
	"sync"
)

// SafeMap - DIY потокобезопасная мапа
type SafeMap[K comparable, V any] struct {
	mutex    sync.Mutex
	innerMap map[K]V
}

// CreateSafeMap - создает новую безопасную мапу
func CreateSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		mutex:    sync.Mutex{},
		innerMap: make(map[K]V),
	}
}

// Load - загружает значение из мапы
func (s *SafeMap[K, V]) Load(key K) (V, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	val, ok := s.innerMap[key]
	return val, ok
}

// Store - сохраняет значение в мапу
func (s *SafeMap[K, V]) Store(key K, value V) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.innerMap[key] = value
}

// Delete - удаляет значение из мапы
func (s *SafeMap[K, V]) Delete(key K) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.innerMap, key)
}
