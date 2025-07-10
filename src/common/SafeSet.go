package common

import (
	"sync"
)

// SafeSet - DIY потокобезопасный сет
type SafeSet[K comparable] struct {
	mutex    sync.Mutex
	innerMap map[K]bool
}

// CreateSafeSet - создает новый безопасный сет
func CreateSafeSet[K comparable]() *SafeSet[K] {
	return &SafeSet[K]{
		mutex:    sync.Mutex{},
		innerMap: make(map[K]bool),
	}
}

// Load - загружает значение из мапы
func (s *SafeSet[K]) Load(key K) (bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	value, ok := s.innerMap[key]
	return value && ok
}

// Store - сохраняет значение в мапу
func (s *SafeSet[K]) Store(key K) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, ok := s.innerMap[key]
	if ok {
		return
	}
	s.innerMap[key] = true
}

// Delete - удаляет значение из мапы
func (s *SafeSet[K]) Delete(key K) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.innerMap, key)
}