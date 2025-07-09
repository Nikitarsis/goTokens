package interfaces

// IRepositoryConfig определяет интерфейс конфигурации репозитория
type IRepositoryConfig interface {
	IPostgresConfig
	TracePorts() bool
}
