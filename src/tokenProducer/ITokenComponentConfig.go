package tokenProducer

// Интерфейс для конфигурации компонентов токенов
type ITokenComponentConfig interface {
	// Размер канала для ключей
	GetKeyChannelSize() uint
	// Размер канала для идентификаторов
	GetIdChannelSize() uint
	// Эмиттер токенов
	GetIssuer() string
}
