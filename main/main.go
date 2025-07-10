package main

import (
	"fmt"

	co "github.com/Nikitarsis/goTokens/common"
	con "github.com/Nikitarsis/goTokens/controller"
	it "github.com/Nikitarsis/goTokens/iptracer"
	repo "github.com/Nikitarsis/goTokens/repository"
	ri "github.com/Nikitarsis/goTokens/repository/interfaces"
	tok "github.com/Nikitarsis/goTokens/tokenProducer"
)

var keyRepository ri.IKeyRepository
var userRepository co.IUserAgentRepository
var ipRepository ri.IIpRepository
var config IConfig
var tokenComponent tok.ITokenComponent
var ipTracer co.IIpTracer

func main() {
	// Создание репозиториев долгосрочного хранения
	config = NewTestConfig()
	fmt.Println("Create repositories")
	keyRepository = repo.CreateKeyRepository(config)
	userRepository = repo.CreateUserRepository(config)
	ipRepository = repo.CreateIPRepository(config)
	// Создание компонента трассировки IP
	fmt.Println("Create IP tracer")
	ipTracer = it.CreateDefaultTracer(config, ipRepository.SaveIp, ipRepository.CheckIp)
	// Создание компонента обработки токенов
	fmt.Println("Create token component")
	tokenComponent = tok.NewTokenComponentDefault(keyRepository, config)
	// Создание обработчика HTTP
	fmt.Println("Create HTTP handler")
	handlerBuilder := con.InitHttpServerBuilder()
	// Возвращает новую пару токенов
	handlerBuilder.AddHandler(
		"/token/new",
		con.NewTokensPairGetter(
			tokenComponent.CreateTokens,
			userRepository,
			ipTracer,
		),
	)
	// Обновляет существующую пару токенов, если refresh токен верный
	handlerBuilder.AddHandler(
		"/token/refresh",
		con.NewTokensRefresher(
			tokenComponent.CreateTokens,
			tokenComponent.ParseToken,
			userRepository,
			ipTracer,
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
