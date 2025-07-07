package inmemory

import (
	"fmt"

	co "github.com/Nikitarsis/goTokens/common"
)

type IpRepositoryInMemory struct {
}

func NewIpRepositoryInMemory() co.IIpRepository {
	return &IpRepositoryInMemory{}
}

func (r *IpRepositoryInMemory) TraceIp(data co.DataIP) {
	fmt.Printf("KeyId: %s, IP: %s\n", data.KeyId, data.IP.String())
}
