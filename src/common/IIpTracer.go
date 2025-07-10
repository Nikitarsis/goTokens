package common

// IIpTracer - интерфейс для работы с репозиторием IP-адресов
type IIpTracer interface {
	// TraceIp - трассирует IP-адрес
	TraceIp(data DataIP)
}
