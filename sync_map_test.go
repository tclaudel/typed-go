package typed_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"typed"
)

func TestNewSyncMap(t *testing.T) {
	t.Parallel()

	m := typed.NewSyncMap[int]()

	if m == nil {
		t.Error("Expected a non-nil map")
	}
}

func TestSyncMap_StoreLoad(t *testing.T) {
	t.Parallel()

	m := typed.NewSyncMap[int]()

	m.Store("key", 1)

	v, ok := m.Load("key")

	if !ok {
		t.Error("Expected to load the value")
	}

	if v != 1 {
		t.Errorf("Expected 1, got %d", v)
	}
}

func TestSyncMap_CompareAndDelete(t *testing.T) {
	t.Parallel()

	m := typed.NewSyncMap[int]()

	m.Store("key", 1)

	assert.True(t, m.CompareAndDelete("key", 1))
	assert.False(t, m.CompareAndDelete("key", 1))
}

func TestSyncMap_CompareAndSwap(t *testing.T) {
	t.Parallel()

	m := typed.NewSyncMap[int]()

	m.Store("key", 1)

	assert.True(t, m.CompareAndSwap("key", 1, 2))
	assert.False(t, m.CompareAndSwap("key", 1, 2))
}

func TestSyncMap_Delete(t *testing.T) {
	t.Parallel()

	m := typed.NewSyncMap[int]()

	m.Store("key", 1)

	m.Delete("key")

	_, ok := m.Load("key")

	assert.False(t, ok, "Expected the key to be deleted")
}

func TestSyncMap_LoadAndDelete(t *testing.T) {
	t.Parallel()

	m := typed.NewSyncMap[int]()

	m.Store("key", 1)

	v, ok := m.LoadAndDelete("key")

	if !ok {
		t.Error("Expected to load the value")
	}

	if v != 1 {
		t.Errorf("Expected 1, got %d", v)
	}

	_, ok = m.Load("key")

	assert.False(t, ok, "Expected the key to be deleted")
}

func TestSyncMap_LoadOrStore(t *testing.T) {
	t.Parallel()

	m := typed.NewSyncMap[int]()

	v, ok := m.LoadOrStore("key", 1)

	if ok {
		t.Error("Expected to store the value")
	}

	if v != 1 {
		t.Errorf("Expected 1, got %d", v)
	}

	v, ok = m.LoadOrStore("key", 2)

	if !ok {
		t.Error("Expected to load the value")
	}

	if v != 1 {
		t.Errorf("Expected 1, got %d", v)
	}
}

func TestSyncMap_Range(t *testing.T) {
	t.Parallel()

	m := typed.NewSyncMap[int]()

	m.Store("key1", 1)
	m.Store("key2", 2)

	count := 0

	m.Range(func(key string, value int) bool {
		count++

		return true
	})

	if count != 2 {
		t.Errorf("Expected 2, got %d", count)
	}
}

func TestSyncMap_Swap(t *testing.T) {
	t.Parallel()

	m := typed.NewSyncMap[int]()

	m.Store("key", 1)

	v, ok := m.Swap("key", 2)

	if !ok {
		t.Error("Expected to swap the value")
	}

	if v != 1 {
		t.Errorf("Expected 1, got %d", v)
	}

	v, ok = m.Load("key")

	if !ok {
		t.Error("Expected to load the value")
	}

	if v != 2 {
		t.Errorf("Expected 2, got %d", v)
	}
}
