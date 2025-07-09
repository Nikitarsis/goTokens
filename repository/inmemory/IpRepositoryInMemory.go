package inmemory

import (
	"fmt"
	"sync"

	co "github.com/Nikitarsis/goTokens/common"
	inter "github.com/Nikitarsis/goTokens/repository/interfaces"
)

// In-Memory репозиторий IP, отправляющий данные в консоль.
//
// Да, это не репозиторий на самом деле, но это простая реализация для тестирования.
type IpRepositoryInMemory struct {
	savePorts bool
	innerMap  *co.SafeMap[co.UUID, *co.SafeSet[string]]
	mutex     sync.Mutex
}

// CreateInMemoryIPRepository создает новый экземпляр IpRepositoryInMemory.
func CreateInMemoryIPRepository(savePorts bool) inter.IIpRepository {
	return &IpRepositoryInMemory{
		savePorts: savePorts,
		innerMap:  co.CreateSafeMap[co.UUID, *co.SafeSet[string]](),
		mutex:     sync.Mutex{},
	}
}

func (ir *IpRepositoryInMemory)	SaveIp(ip co.DataIP) error {
	ir.mutex.Lock()
	defer ir.mutex.Unlock()

	data, ok := ir.innerMap.Load(ip.KeyId)
	if !ok {
		data = co.CreateSafeSet[string]()
	}
	if (ir.savePorts) {
		data.Store(fmt.Sprintf("%s:%d", ip.IP.String(), ip.Port))
	}
	data.Store(ip.IP.String())
	ir.innerMap.Store(ip.KeyId, data)
	return nil
}

func (ir *IpRepositoryInMemory)	CheckIp(ip co.DataIP) bool {
	ir.mutex.Lock()
	defer ir.mutex.Unlock()

	storedIp, ok := ir.innerMap.Load(ip.KeyId)
	if !ok {
		return false
	}
	if ir.savePorts {
		check := storedIp.Load(fmt.Sprintf("%s:%d", ip.IP.String(), ip.Port))
		return check
	}
	check := storedIp.Load(ip.IP.String())
	return check
}
