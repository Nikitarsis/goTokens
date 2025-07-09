package iptracer

import "time"

type ITracerConfig interface {
	GetWebhookURL() string
	ShouldSendWebhookMessage() bool
	GetDelay() time.Duration
	GetBufferSize() uint
}
