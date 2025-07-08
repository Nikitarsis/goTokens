package repository

import (
	co "github.com/Nikitarsis/goTokens/common"
	inmem "github.com/Nikitarsis/goTokens/repository/inmemory"
)

// CreateKeyRepository создает новый экземпляр IKeyRepository.
func CreateKeyRepository() IKeyRepository {
	return inmem.CreateInMemoryKeyRepository()
}

// CreateUserRepository создает новый экземпляр IUserAgentRepository.
func CreateUserRepository() co.IUserAgentRepository {
	return inmem.CreateInMemoryUserRepository()
}

// CreateIPRepository создает новый экземпляр IIpRepository.
func CreateIPRepository() co.IIpRepository {
	return inmem.CreateInMemoryIPRepository()
}
