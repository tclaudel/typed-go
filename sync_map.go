package typed

import "sync"

type SyncMap[T any] struct {
	innerMap     sync.Map
	defaultValue T
}

func NewSyncMap[T any]() *SyncMap[T] {
	return new(SyncMap[T])
}

// CompareAndDelete deletes the key if the value is equal to the value in the map.
func (m *SyncMap[T]) CompareAndDelete(key string, value T) bool {
	return m.innerMap.CompareAndDelete(key, value)
}

// CompareAndSwap compares the value for the key with the old value and if it is equal, swaps the value with the new value.
func (m *SyncMap[T]) CompareAndSwap(key string, oldValue, newValue T) bool {
	return m.innerMap.CompareAndSwap(key, oldValue, newValue)
}

// Delete deletes the key from the map.
func (m *SyncMap[T]) Delete(key string) {
	m.innerMap.Delete(key)
}

// Load returns the value for the key from the map.
func (m *SyncMap[T]) Load(key string) (T, bool) {
	v, ok := m.innerMap.Load(key)
	if !ok {
		return m.defaultValue, false
	}

	return v.(T), ok
}

// LoadAndDelete loads the value for the key and deletes the key from the map.
func (m *SyncMap[T]) LoadAndDelete(key string) (T, bool) {
	v, ok := m.innerMap.LoadAndDelete(key)
	if !ok {
		return m.defaultValue, false
	}

	return v.(T), true
}

// LoadOrStore loads the value for the key and stores the value if the key is not present in the map.
func (m *SyncMap[T]) LoadOrStore(key string, value T) (actual T, loaded bool) {
	v, ok := m.innerMap.LoadOrStore(key, value)
	if !ok {
		return m.defaultValue, false
	}

	return v.(T), true
}

// Range calls the function for each key, value pair in the map.
func (m *SyncMap[T]) Range(f func(key string, value T) bool) {
	m.innerMap.Range(func(key, value interface{}) bool {
		return f(key.(string), value.(T))
	})
}

func (m *SyncMap[T]) Store(key string, value T) {
	m.innerMap.Store(key, value)
}

func (m *SyncMap[T]) Swap(key string, value T) (T, bool) {
	v, ok := m.innerMap.Swap(key, value)
	if !ok {
		return m.defaultValue, false
	}

	return v.(T), true
}
