package controller

import (
	"net/http"
	"strings"
	"testing"

	co "github.com/Nikitarsis/goTokens/common"
)

func TestRequestGet(t *testing.T) {
	uid, err := co.GetUUIDFromString("AAAAAAAB41151824")
	reader := strings.NewReader("nothing");
	if err != nil {
		t.Fatalf("Failed to get UUID from string: %v", err)
	}
	url := "/tokens/?uid=" + uid.ToString()
	request, err := http.NewRequest(http.MethodGet, url, reader)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	controller := NewTokensPairGetter(func(userId co.UUID) (co.TokensPair, error) {
		if uid.ToString() != userId.ToString() {
			t.Fatal("User ID mismatch")
		}
		return co.TokensPair{
			Access:  "access_token",
			Refresh: "refresh_token",
		}, nil
	})
	ret, err := controller.parseRequestGet(request)
	if err != nil {
		t.Fatalf("Failed to parse request: %v", err)
	}
	if ret != uid {
		t.Fatalf("Expected user ID %v, got %v", uid, ret)
	}
}
