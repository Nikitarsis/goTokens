package common

import "net"

type DataIP struct {
	IP       net.IP
	Port     uint16
	UserId   UUID
	KeyId    UUID
}

func GetTestDataIP() DataIP {
	return DataIP{
		IP:     net.ParseIP("192.168.1.1"),
		Port:   8080,
		UserId: GetTestUUID(),
		KeyId:  GetTestUUID(),
	}
}