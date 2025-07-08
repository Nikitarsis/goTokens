package main

import (
	repo "github.com/Nikitarsis/goTokens/repository"
	tok "github.com/Nikitarsis/goTokens/tokenProducer"
	con "github.com/Nikitarsis/goTokens/controller"
)

func main() {
	// Создание репозиториев долгосрочного хранения
	keyRepository := repo.CreateKeyRepository()
	userRepository := repo.CreateUserRepository()
	ipRepository := repo.CreateIPRepository()
	// Создание компонента обработки токенов
	tokenComponent := tok.NewTokenComponentDefault(keyRepository, NewTestConfig())
	// Создание обработчика HTTP
	handlerBuilder := con.InitHttpServerBuilder()
	// Возвращает новую пару токенов
	handlerBuilder.AddHandler(
		"/token/new",
		con.NewTokensPairGetter(
			tokenComponent.CreateTokens, 
			userRepository, 
			ipRepository,
		),
	)
	// Обновляет существующую пару токенов, если refresh токен верный
	handlerBuilder.AddHandler(
		"/token/refresh", 
		con.NewTokensRefresher(
			tokenComponent.CreateTokens, 
			tokenComponent.ParseToken, 
			userRepository, 
			ipRepository, 
			keyRepository.DropKey,
		),
	)
	// Возвращает uid, если access токен верный
	handlerBuilder.AddHandler(
		"/id",
		con.NewTokensIdGetter(
			tokenComponent.ParseToken,
		),
	)
	// Удаляет токены
	handlerBuilder.AddHandler(
		"/unauthorize",
		con.NewTokensUnauthorizer(
			tokenComponent.ParseToken,
			keyRepository.DropKey,
		),
	)

	server := handlerBuilder.Build()
	server.ListenAndServe()
}
