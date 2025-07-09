package iptracer

import (
	co "github.com/Nikitarsis/goTokens/common"
)

type ITracerIp interface {
	TraceIP(ip co.DataIP) error
}
