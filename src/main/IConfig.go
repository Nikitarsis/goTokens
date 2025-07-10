package main

import (
	"time"

	co "github.com/Nikitarsis/goTokens/common"
	it "github.com/Nikitarsis/goTokens/iptracer"
	ri "github.com/Nikitarsis/goTokens/repository/interfaces"
	prod "github.com/Nikitarsis/goTokens/tokenProducer"
)

// IConfig определяет интерфейс для работы с конфигурацией
type IConfig interface {
	co.IDefaultConfig
	prod.ITokenComponentConfig
	ri.IRepositoryConfig
	it.ITracerConfig
}

// NewTestConfig создает новый экземпляр TestConfig
func NewTestConfig() IConfig {
	return &TestConfig{}
}

// TestConfig реализует интерфейс IConfig
type TestConfig struct{}

// GetBufferSize возвращает размер буфера
func (tc TestConfig) GetBufferSize() uint {
	return 100
}

// GetDelay возвращает задержку
func (tc TestConfig) GetDelay() time.Duration {
	return 5 * time.Second
}

// GetWebhookURL возвращает URL вебхука
func (tc TestConfig) GetWebhookURL() string {
	return "http://localhost:8080/webhook"
}

// ShouldSendWebhookMessage возвращает, нужно ли отправлять сообщение вебхука
func (tc TestConfig) ShouldSendWebhookMessage() bool {
	return false
}

// GetConnectionString возвращает строку подключения
func (tc TestConfig) GetConnectionString() string {
	return "user=postgres password=postgres host=db dbname=appdb sslmode=disable"
}

// GetKeyChannelSize возвращает размер канала ключей
func (tc TestConfig) GetKeyChannelSize() uint {
	return 100
}

// GetIdChannelSize возвращает размер канала идентификаторов
func (tc TestConfig) GetIdChannelSize() uint {
	return 200
}

// GetIssuer возвращает издателя
func (tc TestConfig) GetIssuer() string {
	return "test-issuer"
}

// IsDebugMode возвращает, включен ли режим отладки
func (tc TestConfig) IsDebugMode() bool {
	return true
}

// TracePorts возвращает, нужно ли отслеживать порты
func (tc TestConfig) TracePorts() bool {
	return true
}
