package common

type IIpRepository interface {
	TraceIp(kid UUID, ip string)
}