package typed_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tclaudel/typed-go"
)

func TestNewList(t *testing.T) {
	t.Parallel()

	l := typed.NewList[int]()

	assert.NotNil(t, l)
}

func TestList_Init(t *testing.T) {
	t.Parallel()

	l := typed.NewList[int]()

	l.Init()

	assert.Equal(t, 0, l.Len())
}

func TestList_PushFront(t *testing.T) {
	t.Parallel()

	l := typed.NewList[int]()

	l.PushFront(1)
	e := l.Front()

	assert.Equal(t, 1, e.Value())
}

func TestList_PushBack(t *testing.T) {
	t.Parallel()

	l := typed.NewList[int]()

	l.PushBack(1)
	e := l.Back()

	assert.Equal(t, 1, e.Value())
}

func TestList_Len(t *testing.T) {
	t.Parallel()

	l := typed.NewList[int]()

	l.PushBack(1)
	l.PushBack(2)

	assert.Equal(t, 2, l.Len())
}

func TestList_Remove(t *testing.T) {
	t.Parallel()

	l := typed.NewList[int]()

	l.PushBack(1)
	e := l.Front()

	assert.Equal(t, 1, l.Remove(e))
}

func TestElement_InsertBefore(t *testing.T) {
	t.Parallel()

	l := typed.NewList[int]()

	l.PushBack(1)
	e := l.Front()

	l.InsertBefore(2, e)
	e = l.Front()

	assert.Equal(t, 2, e.Value())
}

func TestElement_InsertAfter(t *testing.T) {
	t.Parallel()

	l := typed.NewList[int]()

	l.PushBack(1)
	e := l.Front()

	l.InsertAfter(2, e)
	e = l.Back()

	assert.Equal(t, 2, e.Value())
}

func TestElement_MoveToFront(t *testing.T) {
	t.Parallel()

	l := typed.NewList[int]()

	l.PushBack(1)
	l.PushBack(2)
	e := l.Back()

	l.MoveToFront(e)
	e = l.Front()

	assert.Equal(t, 2, e.Value())
}

func TestElement_MoveToBack(t *testing.T) {
	t.Parallel()

	l := typed.NewList[int]()

	l.PushBack(1)
	l.PushBack(2)
	e := l.Front()

	l.MoveToBack(e)
	e = l.Back()

	assert.Equal(t, 1, e.Value())
}

func TestElement_MoveBefore(t *testing.T) {
	t.Parallel()

	l := typed.NewList[int]()

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	e := l.Back()
	mark := l.Front()

	l.MoveBefore(e, mark)

	e = l.Front()
	assert.Equal(t, 3, e.Value())
}

func TestElement_MoveAfter(t *testing.T) {
	t.Parallel()

	l := typed.NewList[int]()

	l.PushBack(1)
	l.PushBack(2)
	e := l.Front()

	l.MoveAfter(e, l.Back())
	e = l.Back()

	assert.Equal(t, 1, e.Value())
}

func TestList_PushFrontList(t *testing.T) {
	t.Parallel()

	l := typed.NewList[int]()
	other := typed.NewList[int]()

	other.PushBack(1)
	other.PushBack(2)

	l.PushFrontList(other)

	assert.Equal(t, 2, l.Len())
}

func TestList_PushBackList(t *testing.T) {
	t.Parallel()

	l := typed.NewList[int]()
	other := typed.NewList[int]()

	other.PushBack(1)
	other.PushBack(2)

	l.PushBackList(other)

	assert.Equal(t, 2, l.Len())
}

func TestElement_Next(t *testing.T) {
	t.Parallel()

	l := typed.NewList[int]()

	l.PushBack(1)
	l.PushBack(2)
	e := l.Front()

	assert.Equal(t, 2, e.Next().Value())
}

func TestElement_Prev(t *testing.T) {
	t.Parallel()

	l := typed.NewList[int]()

	l.PushBack(1)
	l.PushBack(2)
	e := l.Back()

	assert.Equal(t, 1, e.Prev().Value())
}

func TestElement_Value(t *testing.T) {
	t.Parallel()

	l := typed.NewList[int]()

	l.PushBack(1)
	e := l.Front()

	assert.Equal(t, 1, e.Value())
}
