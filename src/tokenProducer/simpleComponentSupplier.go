package tokenProducer

import (
	"crypto/rand"
	"log"

	co "github.com/Nikitarsis/goTokens/common"
	goid "github.com/google/uuid"
)

// Создание нового ключа
func createKey() (co.Key, error) {
	id, err := createId()
	if err != nil {
		return co.Key{}, err
	}
	key := make([]byte, 64)
	rand.Read(key) // Случайная генерация ключа
	ret := co.CreateKeyFromBytes(id, key)
	return ret, nil
}

// Создание нового идентификатора
func createId() (co.UUID, error) {
	// Использование Google UUID для генерации
	bytesId, _ := goid.New().MarshalBinary()
	id, err := co.GetUUIDFromBytes(bytesId)
	if err != nil {
		return co.UUID{}, err
	}
	return id, nil
}

// Цикл, производящий новые ключи
func keyRoutine(keyChan chan co.Key) {
	for {
		key, err := createKey()
		if err != nil {
			log.Println(err)
			continue
		}
		keyChan <- key
	}
}

// Цикл, производящий новые идентификаторы
func idRoutine(idChan chan co.UUID) {
	for {
		jti, err := createId()
		if err != nil {
			log.Println(err)
			continue
		}
		idChan <- jti
	}
}

// simpleComponentSupplier - простая реализация поставщика компонентов
type simpleComponentSupplier struct {
	keyChan chan co.Key
	idChan  chan co.UUID
}

// Функция создает новый экземпляр simpleComponentSupplier
func newSimpleComponentSupplier(keyChanSize, idChanSize uint) *simpleComponentSupplier {
	keyChan := make(chan co.Key, keyChanSize)
	idChan := make(chan co.UUID, idChanSize)
	go keyRoutine(keyChan)
	go idRoutine(idChan)
	return &simpleComponentSupplier{
		keyChan: keyChan,
		idChan:  idChan,
	}
}

// NewKey создает новый ключ
func (s *simpleComponentSupplier) NewKey() co.Key {
	return <-s.keyChan
}

// NewId создает новый идентификатор
func (s *simpleComponentSupplier) NewId() co.UUID {
	return <-s.idChan
}
