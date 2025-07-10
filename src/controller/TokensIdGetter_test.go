package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	co "github.com/Nikitarsis/goTokens/common"
)

// TestGetTokenId - тест для функции GetTokenId	
func TestGetTokenId(t *testing.T) {
	// Тестовое окружение
	userId := co.GetTestUUID()
	token := co.GetTestToken()
	tokenData := co.TokenData{
		Token:   token,
		TokenId: co.GetTestUUID(),
		UserId:  userId,
		KeyId:   co.GetTestUUID(),
		Type:    co.RefreshToken,
	}
	parseToken := func(tkn co.Token) (co.TokenData, error) {
		if tkn.ToString() != token.ToString() {
			t.Errorf("Invalid token %s != %s", tkn.ToString(), token.ToString())
		}
		return tokenData, nil
	}
	//Создание тестовых классов
	refresher := NewTokensIdGetter(parseToken)
	userToken := UserToken{
		UID: userId.ToString(),
		Token: token.ToString(),
	}
	body, err := json.Marshal(userToken)
	if err != nil {
		t.Fatalf("Failed to marshal user token: %v", err)
	}
	reader := bytes.NewReader(body)
	request, err := http.NewRequest(http.MethodPost, "/refresh-tokens", reader)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	response := refresher.GetTokenId(request)
	time.Sleep(10 * time.Millisecond)
	// Проверка ответа
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %v, got %v", http.StatusOK, response.StatusCode)
	}
}