package main

import (
	"fmt"

	co "github.com/Nikitarsis/goTokens/common"
	con "github.com/Nikitarsis/goTokens/controller"
	repo "github.com/Nikitarsis/goTokens/repository"
	tok "github.com/Nikitarsis/goTokens/tokenProducer"
)

var keyRepository repo.IKeyRepository
var userRepository co.IUserAgentRepository
var ipRepository co.IIpRepository
var config co.IDefaultConfig
var tokenComponent tok.ITokenComponent

func main() {
	// Создание репозиториев долгосрочного хранения
	fmt.Println("Create repositories")
	keyRepository = repo.CreateKeyRepository()
	userRepository = repo.CreateUserRepository()
	ipRepository = repo.CreateIPRepository()
	// Создание компонента обработки токенов
	fmt.Println("Create token component")
	tokenComponent = tok.NewTokenComponentDefault(keyRepository, NewTestConfig())
	config = NewTestConfig()
	// Создание обработчика HTTP
	fmt.Println("Create HTTP handler")
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
	//Тестовый обработчик
	handlerBuilder.AddHandler(
		"/test/",
		GetTestHandler(config.IsDebugMode()),
	)
	// Удаляет токены
	handlerBuilder.AddHandler(
		"/unauthorize",
		con.NewTokensUnauthorizer(
			tokenComponent.ParseToken,
			keyRepository.DropKey,
		),
	)
	handlerBuilder.SetServerAddr(":10000")

	fmt.Println("Start HTTP server")
	server := handlerBuilder.Build()
	fmt.Printf("Started server on %s\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
