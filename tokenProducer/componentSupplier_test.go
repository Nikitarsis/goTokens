package tokenProducer

import (
	"testing"
)

// Тестовый поставщик контента
func getTestComponentSupplier() *simpleComponentSupplier {
	return newSimpleComponentSupplier(5, 5)
}

// Тест на создание новых компонентов без ошибок
func TestNew(t *testing.T) {
	componentSupplier := getTestComponentSupplier()
	componentSupplier.NewId()
	componentSupplier.NewKey()
}

// Тест на корректную работу поставщика компонентов
func TestWorking(t *testing.T) {
	componentSupplier := getTestComponentSupplier()
	for i := 0; i < 500; i++ {
		componentSupplier.NewId()
		componentSupplier.NewKey()
	}
}

// Тест уникальности ключей
func TestNewKeys(t *testing.T) {
	componentSupplier := getTestComponentSupplier()
	key1 := componentSupplier.NewKey()
	key2 := componentSupplier.NewKey()
	if key1.ToString() == key2.ToString() {
		t.Error("Expected unique keys to be generated")
	}
}

// Тест уникальности идентификаторов
func TestNewIds(t *testing.T) {
	componentSupplier := getTestComponentSupplier()
	id1 := componentSupplier.NewId()
	id2 := componentSupplier.NewId()
	if id1.ToString() == id2.ToString() {
		t.Error("Expected unique IDs to be generated")
	}
}
