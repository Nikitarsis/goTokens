package tokenProducer

import (
	co "github.com/Nikitarsis/goTokens/common"
)

type ITokenComponent interface{
	CreateTokens(uid co.UUID) (map[string]co.TokenData, error)
	ParseToken(token co.Token) (co.TokenData, error)
} 

// TokenComponent - компонент для работы с данным модулем
type TokenComponent struct {
	producer          *tokenProducer
	parser           *tokenParser
	keyRepository     co.IKeyKeepingRepository
	componentSupplier *simpleComponentSupplier
}

// NewTokenComponentDefault создает новый экземпляр TokenComponent с конфигурацией
func NewTokenComponentDefault(keyRepository co.IKeyKeepingRepository, config ITokenComponentConfig) ITokenComponent {
	componentSupplier := newSimpleComponentSupplier(config.GetKeyChannelSize(), config.GetIdChannelSize())
	idSupply := componentSupplier.NewId
	idFactory := keyRepository.GetKey
	return &TokenComponent{
		producer:          NewTokenProducer(co.NewIssuer(config.GetIssuer()), idSupply),
		parser:           NewTokenParser(idFactory),
		keyRepository:     keyRepository,
		componentSupplier: componentSupplier,
	}
}

// CreateTokens создает пару токенов (access и refresh) для данного пользователя
func (tc *TokenComponent) CreateTokens(uid co.UUID) (map[string]co.TokenData, error) {
	key := tc.componentSupplier.NewKey()
	// Создание access токена
	access, err := tc.producer.CreateAccessToken(key, uid)
	if err != nil {
		return nil, err
	}
	// Создание refresh токена
	refresh, err := tc.producer.CreateRefreshToken(key, uid)
	if err != nil {
		return nil, err
	}
	// Асинхронное сохранение ключа
	go func() {
		tc.keyRepository.SaveKey(key)
	}()
	return map[string]co.TokenData{
		"access":  access,
		"refresh": refresh,
	}, nil
}

// ParseToken парсит токен и возвращает его данные
// Если токен невалиден, возвращает ошибку без значения co.TokenData
func (tc *TokenComponent) ParseToken(token co.Token) (co.TokenData, error) {
	ret, err := tc.parser.GetTokenData(token)
	if err != nil {
		return co.TokenData{}, err
	}
	return ret, nil
}
