package typed_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tclaudel/typed-go"
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

	assert.Equal(t, 1, v)
	assert.True(t, ok)

	v, ok = m.Load("Not Found")

	assert.Equal(t, 0, v)
	assert.False(t, ok)

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

	v, ok := m.Load("key")

	assert.Equal(t, 0, v)
	assert.False(t, ok, "Expected the key to be deleted")
}

func TestSyncMap_LoadAndDelete(t *testing.T) {
	t.Parallel()

	m := typed.NewSyncMap[int]()

	m.Store("key", 1)

	v, ok := m.LoadAndDelete("key")

	assert.Equal(t, 1, v)
	assert.True(t, ok)

	v, ok = m.Load("key")

	assert.Equal(t, 0, v)
	assert.False(t, ok, "Expected the key to be deleted")
}

func TestSyncMap_LoadAndDeleteNOK(t *testing.T) {
	t.Parallel()

	m := typed.NewSyncMap[int]()

	v, ok := m.LoadAndDelete("key")

	assert.Equal(t, 0, v)
	assert.False(t, ok)
}

func TestSyncMap_LoadOrStore(t *testing.T) {
	t.Parallel()

	m := typed.NewSyncMap[int]()

	m.Store("key", 1)

	v, ok := m.LoadOrStore("key", 2)

	assert.Equal(t, 1, v)
	assert.True(t, ok)

	v, ok = m.LoadOrStore("new key", 1)

	assert.Equal(t, 0, v)
	assert.False(t, ok)

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

	assert.Equal(t, 2, count)
}

func TestSyncMap_Swap(t *testing.T) {
	t.Parallel()

	m := typed.NewSyncMap[int]()

	m.Store("key", 1)

	v, ok := m.Swap("key", 2)

	assert.Equal(t, 1, v)
	assert.True(t, ok)

	v, ok = m.Swap("not found", 1)

	assert.Equal(t, 0, v)
	assert.False(t, ok)
}
