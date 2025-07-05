package tokenProducer

import (
	"testing"
)

func getTestComponentSupplier() *simpleComponentSupplier {
	return NewSimpleComponentSupplier(5, 5)
}

func TestNew(t *testing.T) {
	componentSupplier := getTestComponentSupplier()
	componentSupplier.NewId()
	componentSupplier.NewKey()
}

func TestNewKeys(t *testing.T) {
	componentSupplier := getTestComponentSupplier()
	key1 := componentSupplier.NewKey()
	key2 := componentSupplier.NewKey()
	if key1.ToString() == key2.ToString() {
		t.Error("Expected unique keys to be generated")
	}
}

func TestNewIds(t *testing.T) {
	componentSupplier := getTestComponentSupplier()
	id1 := componentSupplier.NewId()
	id2 := componentSupplier.NewId()
	if id1.ToString() == id2.ToString() {
		t.Error("Expected unique IDs to be generated")
	}
}
