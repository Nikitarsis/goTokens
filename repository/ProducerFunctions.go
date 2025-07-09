package repository

import (
	co "github.com/Nikitarsis/goTokens/common"
	inmem "github.com/Nikitarsis/goTokens/repository/inmemory"
	inter "github.com/Nikitarsis/goTokens/repository/interfaces"
)

// CreateKeyRepository создает новый экземпляр IKeyRepository.
func CreateKeyRepository() inter.IKeyRepository {
	return inmem.CreateInMemoryKeyRepository()
}

// CreateUserRepository создает новый экземпляр IUserAgentRepository.
func CreateUserRepository() co.IUserAgentRepository {
	return inmem.CreateInMemoryUserRepository()
}

// CreateIPRepository создает новый экземпляр IIpRepository.
func CreateIPRepository(config inter.IRepositoryConfig) inter.IIpRepository {
	return inmem.CreateInMemoryIPRepository(config.TracePorts()	)
}
