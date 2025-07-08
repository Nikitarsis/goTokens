package common

// IIpRepository - интерфейс для работы с репозиторием IP-адресов
type IIpRepository interface {
	// TraceIp - трассирует IP-адрес
	TraceIp(data DataIP)
}
