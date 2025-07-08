package inmemory

import (
	"fmt"

	co "github.com/Nikitarsis/goTokens/common"
)

// In-Memory репозиторий IP, отправляющий данные в консоль.
//
// Да, это не репозиторий на самом деле, но это простая реализация для тестирования.
type IpRepositoryInMemory struct {
}

// CreateInMemoryIPRepository создает новый экземпляр IpRepositoryInMemory.
func CreateInMemoryIPRepository() co.IIpRepository {
	return &IpRepositoryInMemory{}
}

// TraceIp отслеживает IP-адреса.
func (r *IpRepositoryInMemory) TraceIp(data co.DataIP) {
	fmt.Printf("KeyId: %s, UserId: %s, FromPort: %d, IP: %s\n", data.KeyId.ToString(), data.UserId.ToString(), data.Port, data.IP.String())
}
