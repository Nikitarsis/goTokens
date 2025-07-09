package interfaces

type IRepositoryConfig interface {
	IPostgresConfig
	TracePorts() bool
}
