package inmemory

import (
	co "github.com/Nikitarsis/goTokens/common"
	"testing"
)

// TestInMemoryKeyRepositorySave проверяет сохранение ключа в репозитории.
func TestInMemoryKeyRepositorySave(t *testing.T) {
	repo := CreateInMemoryKeyRepository()
	key := co.CreateTestKey()
	repo.SaveKey(key)
	ret, ok := repo.cache.Load(key.GetKid())
	if !ok {
		t.Errorf("Expected to find key %v", key.GetKid())
	}
	if ret.ToString() != key.ToString() {
		t.Errorf("Expected to find key %v, got %v", key.GetKid(), ret)
	}
	if ret.GetKid().ToString() != key.GetKid().ToString() {
		t.Errorf("Expected key ID %v, got %v", key.GetKid().ToString(), ret.GetKid().ToString())
	}
}

// TestInMemoryKeyRepositoryLoad проверяет загрузку ключа из репозитория.
func TestInMemoryKeyRepositoryLoad(t *testing.T) {
	repo := CreateInMemoryKeyRepository()
	key := co.CreateTestKey()
	repo.SaveKey(key)
	ret, ok := repo.cache.Load(key.GetKid())
	if !ok {
		t.Errorf("Expected to find key %v", key.GetKid())
	}
	if ret.ToString() != key.ToString() {
		t.Errorf("Expected to find key %v, got %v", key.GetKid(), ret)
	}
	if ret.GetKid().ToString() != key.GetKid().ToString() {
		t.Errorf("Expected key ID %v, got %v", key.GetKid().ToString(), ret.GetKid().ToString())
	}
}

// TestInMemoryKeyRepositoryDrop проверяет удаление ключа из репозитория.
func TestInMemoryKeyRepositoryDrop(t *testing.T) {
	repo := CreateInMemoryKeyRepository()
	key := co.CreateTestKey()
	repo.SaveKey(key)
	if !repo.DropKey(key.GetKid()) {
		t.Errorf("Expected to drop key %v", key.GetKid())
	}
	if _, ok := repo.cache.Load(key.GetKid()); ok {
		t.Errorf("Expected key %v to be deleted", key.GetKid())
	}
}
