package tokenProducer

import (
	"testing"

	co "github.com/Nikitarsis/goTokens/common"
)

// Тестовый производитель токенов
func getTestTokenProducer(componentSupplier *simpleComponentSupplier) *tokenProducer {
	return NewTokenProducer(co.NewIssuer("test-issuer"), componentSupplier.NewId)
}

// TestCreateClaims тестирует создание заявок
func TestCreateClaims(t *testing.T) {
	componentSupplier := getTestComponentSupplier()
	producer := getTestTokenProducer(componentSupplier)
	uid := componentSupplier.NewId()
	kid := componentSupplier.NewId()
	claims := producer.createClaims(co.AccessToken, uid, kid)
	if claims == nil {
		t.Error("Expected non-nil claims")
	}
}

// TestCreateToken тестирует создание токена
func TestCreateToken(t *testing.T) {
	componentSupplier := getTestComponentSupplier()
	producer := getTestTokenProducer(componentSupplier)
	uid := componentSupplier.NewId()
	key := componentSupplier.NewKey()
	tokenData, err := producer.createToken(key, uid, co.AccessToken)
	if err != nil {
		t.Fatal(err)
	}
	if tokenData.Token.Value == "" {
		t.Error("Expected non-empty token")
	}
	if tokenData.TokenId.ToString() == "" {
		t.Error("Expected non-empty JTI")
	}
	if tokenData.KeyId.ToString() == "" {
		t.Error("Expected non-empty Key ID")
	}
	if tokenData.UserId.ToString() == "" {
		t.Error("Expected non-empty User ID")
	}
	if tokenData.Type != co.AccessToken {
		t.Error("Expected AccessToken type")
	}
}
