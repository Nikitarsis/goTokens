package repository

import (
	"sync"

	co "github.com/Nikitarsis/goTokens/common"
	pg "github.com/Nikitarsis/goTokens/repository/database/postgres"
	in "github.com/Nikitarsis/goTokens/repository/interfaces"
)

type UserAgentRepository struct {
	db       co.IUserAgentRepository
	hotCacheToSave *co.SafeMap[co.UUID, co.UserAgentData]
	mutexMap *co.SafeMap[co.UUID, *sync.Mutex]
	localMutex *sync.Mutex
}

func NewUserAgentRepository(config in.IRepositoryConfig) co.IUserAgentRepository {
	db := pg.CreatePostgresUserAgentRepository(config)
	hotCacheToSave := &co.SafeMap[co.UUID, co.UserAgentData]{}
	mutexMap := &co.SafeMap[co.UUID, *sync.Mutex]{}
	return &UserAgentRepository{
		db:              db,
		hotCacheToSave:  hotCacheToSave,
		mutexMap:        mutexMap,
		localMutex:     &sync.Mutex{},
	}
}

func (kr *UserAgentRepository) getMutex(kid co.UUID) *sync.Mutex {
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

func (kr *UserAgentRepository) SaveUserAgent(kid co.UUID, userAgent co.UserAgentData) error {
	// Сохраняем ключ в горячий кэш сохранения и назначаем удаление
	kr.hotCacheToSave.Store(kid, userAgent)
	defer kr.hotCacheToSave.Delete(kid)
	// Получаем мьютекс для ключа
	mtx := kr.getMutex(kid)
	mtx.Lock()
	defer mtx.Unlock()
	// Начинаем сохранение User-Agent
	return kr.db.SaveUserAgent(kid, userAgent)
}

func (kr *UserAgentRepository) CheckUserAgent(kid co.UUID, userAgent co.UserAgentData) bool {
	// Проверяем горячий кэш сохранения на наличие User-Agent
	toSave, ok := kr.hotCacheToSave.Load(kid)
	// В случае наличия, возвращаем его с true
	if ok {
		return toSave == userAgent
	}
	// Получаем мьютекс для ключа
	mtx := kr.getMutex(kid)
	mtx.Lock()
	defer mtx.Unlock()
	// Начинаем проверка User-Agent
	return kr.db.CheckUserAgent(kid, userAgent)
}