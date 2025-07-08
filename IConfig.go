package main

import (
	prod "github.com/Nikitarsis/goTokens/tokenProducer"
)

type IConfig interface {
	prod.ITokenComponentConfig
}

func NewTestConfig() IConfig {
	return &TestConfig{}
}

type TestConfig struct{}

func (tc TestConfig) GetKeyChannelSize() uint {
	return 100
}

func (tc TestConfig) GetIdChannelSize() uint {
	return 200
}

func (tc TestConfig) GetIssuer() string {
	return "test-issuer"
}
