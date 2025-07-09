package repository

import (
	"sync"

	co "github.com/Nikitarsis/goTokens/common"
	pg "github.com/Nikitarsis/goTokens/repository/database/postgres"
	in "github.com/Nikitarsis/goTokens/repository/interfaces"
)

type KeyRepository struct {
	db       in.IKeyRepository
	hotCacheToSave *co.SafeMap[co.UUID, co.Key]
	hotCacheToRemove *co.SafeSet[co.UUID]
	mutexMap *co.SafeMap[co.UUID, *sync.Mutex]
	localMutex *sync.Mutex
}

func NewKeyRepository(config in.IRepositoryConfig) in.IKeyRepository {
	db := pg.CreatePostgresKeyRepository(config)
	hotCacheToSave := &co.SafeMap[co.UUID, co.Key]{}
	hotCacheToRemove := &co.SafeSet[co.UUID]{}
	mutexMap := &co.SafeMap[co.UUID, *sync.Mutex]{}
	return &KeyRepository{
		db:              db,
		hotCacheToSave:  hotCacheToSave,
		hotCacheToRemove: hotCacheToRemove,
		mutexMap:        mutexMap,
		localMutex:     &sync.Mutex{},
	}
}

func (kr *KeyRepository) getMutex(kid co.UUID) *sync.Mutex {
	// Загрузка мьютекса
	mtx, ok := kr.mutexMap.Load(kid)
	// Если мьютекс уже существует, возвращаем его
	if ok {
		return mtx
	}
	// Если мьютекс не существует, блокируем горутину
	kr.localMutex.Lock()
	defer kr.localMutex.Unlock()
	// Потом проверяем, что мьютекса всё ещё нет
	ret, anotherCheck := kr.mutexMap.Load(kid)
	// Если мьютекс не существует, создаём новый
	if !anotherCheck {
		ret = &sync.Mutex{}
		kr.mutexMap.Store(kid, ret)
	}
	// Возвращаем мьютекс
	return ret
}

func (kr *KeyRepository) SaveKey(key co.Key) {
	// Сохраняем ключ в горячий кэш сохранения и назначаем удаление
	kr.hotCacheToSave.Store(key.GetKid(), key)
	defer kr.hotCacheToSave.Delete(key.GetKid())
	// Получаем мьютекс для ключа
	mtx := kr.getMutex(key.GetKid())
	mtx.Lock()
	defer mtx.Unlock()
	// Начинаем сохранение ключа
	kr.db.SaveKey(key)
}

func (kr *KeyRepository) GetKey(kid co.UUID) (co.Key, bool) {
	// Проверяем горячий кэш сохранения на наличие ключа
	toSave, ok := kr.hotCacheToSave.Load(kid)
	// В случае наличия, возвращаем его с true
	if ok {
		return toSave, true
	}
	// Проверяем горячий кэш удаления на наличие ключа
	check := kr.hotCacheToRemove.Load(kid)
	// В случае наличия, возвращаем его с false
	if check {
		return co.Key{}, false
	}
	// Получаем мьютекс для ключа
	mtx := kr.getMutex(kid)
	mtx.Lock()
	defer mtx.Unlock()
	// Начинаем получение ключа
	return kr.db.GetKey(kid)
}

func (kr *KeyRepository) DropKey(kid co.UUID) bool {
	// Сохраняем ключ в горячий кэш сохранения и назначаем удаление
	kr.hotCacheToRemove.Store(kid)
	defer kr.hotCacheToRemove.Delete(kid)
	// Получаем мьютекс для ключа
	mtx := kr.getMutex(kid)
	mtx.Lock()
	defer mtx.Unlock()
	// Начинаем удаление ключа
	return kr.db.DropKey(kid)
}