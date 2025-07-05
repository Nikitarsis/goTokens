package tokenProducer

import (
	"crypto/rand"
	"log"

	co "github.com/Nikitarsis/goTokens/common"
	goid "github.com/google/uuid"
)

func createKey() (co.Key, error) {
	id, err := createId()
	if err != nil {
		return co.Key{}, err
	}
	key := make([]byte, 64)
	rand.Read(key)
	ret := co.CreateKeyFromBytes(id, key)
	return ret, nil
}

func createId() (co.UUID, error) {
	bytesId, _ := goid.New().MarshalBinary()
	id, err := co.GetUUIDFromBytes(bytesId)
	if err != nil {
		return co.UUID{}, err
	}
	return id, nil
}

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

func jtiRoutine(jtiChan chan co.UUID) {
	for {
		jti, err := createId()
		if err != nil {
			log.Println(err)
			continue
		}
		jtiChan <- jti
	}
}

type SimpleComponentSupplier struct {
	keyChan chan co.Key
	jtiChan chan co.UUID
}

func NewSimpleComponentSupplier(keyChanSize, jtiChanSize uint) *SimpleComponentSupplier {
	keyChan := make(chan co.Key, keyChanSize)
	jtiChan := make(chan co.UUID, jtiChanSize)
	go keyRoutine(keyChan)
	go jtiRoutine(jtiChan)
	return &SimpleComponentSupplier{
		keyChan: keyChan,
		jtiChan: jtiChan,
	}
}

func (s *SimpleComponentSupplier) NewKey() co.Key {
	return <-s.keyChan
}

func (s *SimpleComponentSupplier) NewId() co.UUID {
	return <-s.jtiChan
}
