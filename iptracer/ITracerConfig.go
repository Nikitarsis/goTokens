package iptracer

import "time"

// ITracerConfig определяет интерфейс конфигурации трассировки
type ITracerConfig interface {
	GetWebhookURL() string
	ShouldSendWebhookMessage() bool
	GetDelay() time.Duration
	GetBufferSize() uint
}
