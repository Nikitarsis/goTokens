package repository

import (
	inmem "github.com/Nikitarsis/goTokens/repository/inmemory"
	co "github.com/Nikitarsis/goTokens/common"
)

func CreateKeyRepository() IKeyRepository {
	return inmem.CreateInMemoryKeyRepository()
}

func CreateUserRepository() co.IUserAgentRepository {
	return inmem.CreateInMemoryUserRepository()
}

func CreateIPRepository() co.IIpRepository {
	return inmem.CreateInMemoryIPRepository()
}