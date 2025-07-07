package common

import "net"

type DataIP struct {
	IP       net.IP
	Port     uint16
	UserId   UUID
	KeyId    UUID
}