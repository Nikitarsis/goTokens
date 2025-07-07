package inmemory

import (
	co "github.com/Nikitarsis/goTokens/common"
)

type InMemoryKeyRepository struct {
	cache co.SafeMap[co.UUID, co.Key]
}

func CreateInMemoryKeyRepository() *InMemoryKeyRepository {
	return &InMemoryKeyRepository{
		cache: *co.CreateSafeMap[co.UUID, co.Key](),
	}
}

func (r *InMemoryKeyRepository) SaveKey(key co.Key) {
	r.cache.Store(key.GetKid(), key)
}

func (r *InMemoryKeyRepository) GetKey(kid co.UUID) (co.Key, bool) {
	val, ok := r.cache.Load(kid)
	if !ok {
		return co.Key{}, false
	}
	return val, true
}

func (r *InMemoryKeyRepository) DropKey(kid co.UUID) bool {
	_, ok := r.cache.Load(kid)
	if !ok {
		return false
	}
	r.cache.Delete(kid)
	return true
}
