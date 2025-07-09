package repository

import (
	co "github.com/Nikitarsis/goTokens/common"
	inter "github.com/Nikitarsis/goTokens/repository/interfaces"
)

// CreateKeyRepository создает новый экземпляр IKeyRepository.
func CreateKeyRepository(config inter.IRepositoryConfig) inter.IKeyRepository {
	return NewKeyRepository(config)
}

// CreateUserRepository создает новый экземпляр IUserAgentRepository.
func CreateUserRepository(config inter.IRepositoryConfig) co.IUserAgentRepository {
	return NewUserAgentRepository(config)
}

// CreateIPRepository создает новый экземпляр IIpRepository.
func CreateIPRepository(config inter.IRepositoryConfig) inter.IIpRepository {
	return NewIpRepository(config)
}
