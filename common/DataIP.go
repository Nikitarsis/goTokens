package common

import "net"

// DataIP - структура для хранения информации о IP-адресе и его запросе
type DataIP struct {
	// IP - адрес клиента
	IP       net.IP
	// Port - порт клиента
	Port     uint16
	// UserId - идентификатор клиента
	UserId   UUID
	// KeyId - идентификатор ключа
	KeyId    UUID
}

// GetTestDataIP - возвращает тестовые данные IP
func GetTestDataIP() DataIP {
	return DataIP{
		IP:     net.ParseIP("192.168.1.1"),
		Port:   8080,
		UserId: GetTestUUID(),
		KeyId:  GetTestUUID(),
	}
}