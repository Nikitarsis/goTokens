package common

// IDefaultConfig - интерфейс для получения конфигурации по умолчанию, которая используется в разных модулях
type IDefaultConfig interface {
	// IsDebugMode - возвращает проверку режима отладки
	IsDebugMode() bool
}