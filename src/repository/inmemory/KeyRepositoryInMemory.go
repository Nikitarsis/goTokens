package inmemory

import (
	co "github.com/Nikitarsis/goTokens/common"
)

// In-Memory репозиторий ключей
type InMemoryKeyRepository struct {
	cache co.SafeMap[co.UUID, co.Key]
}

// CreateInMemoryKeyRepository создает новый экземпляр InMemoryKeyRepository.
func CreateInMemoryKeyRepository() *InMemoryKeyRepository {
	return &InMemoryKeyRepository{
		cache: *co.CreateSafeMap[co.UUID, co.Key](),
	}
}

// SaveKey сохраняет ключ в репозитории.
func (r *InMemoryKeyRepository) SaveKey(key co.Key) {
	r.cache.Store(key.GetKid(), key)
}

// GetKey загружает ключ из репозитория.
func (r *InMemoryKeyRepository) GetKey(kid co.UUID) (co.Key, bool) {
	val, ok := r.cache.Load(kid)
	if !ok {
		return co.Key{}, false
	}
	return val, true
}

// DropKey удаляет ключ из репозитория.
func (r *InMemoryKeyRepository) DropKey(kid co.UUID) bool {
	_, ok := r.cache.Load(kid)
	if !ok {
		return false
	}
	r.cache.Delete(kid)
	return true
}
