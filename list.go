package typed

import "container/list"

// Element is an element of a linked list.
type Element[T any] struct {
	innerElement *list.Element
}

// Next returns the next list element or nil.
func (e *Element[T]) Next() *Element[T] {
	return &Element[T]{innerElement: e.innerElement.Next()}
}

// Prev returns the previous list element or nil.
func (e *Element[T]) Prev() *Element[T] {
	return &Element[T]{innerElement: e.innerElement.Prev()}
}

// Value returns the value of the element.
func (e *Element[T]) Value() T {
	return e.innerElement.Value.(T)
}

// List represents a doubly linked list.
type List[T any] struct {
	innerList *list.List
}

// Init initializes or clears list l.
func (l *List[T]) Init() *List[T] {
	l.innerList.Init()
	return l
}

// NewList returns an initialized list.
func NewList[T any]() *List[T] {
	return &List[T]{
		innerList: list.New(),
	}
}

// Len returns the number of elements of list l.
func (l *List[T]) Len() int {
	return l.innerList.Len()
}

// Front returns the first element of list l or nil if the list is empty.
func (l *List[T]) Front() *Element[T] {
	return &Element[T]{innerElement: l.innerList.Front()}
}

// Back returns the last element of list l or nil if the list is empty.
func (l *List[T]) Back() *Element[T] {
	return &Element[T]{innerElement: l.innerList.Back()}
}

// Remove removes e from l if e is an element of list l.
func (l *List[T]) Remove(e *Element[T]) T {
	return l.innerList.Remove(e.innerElement).(T)
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *List[T]) PushFront(v T) *Element[T] {
	return &Element[T]{innerElement: l.innerList.PushFront(v)}
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *List[T]) PushBack(v T) *Element[T] {
	return &Element[T]{innerElement: l.innerList.PushBack(v)}
}

// InsertBefore inserts a new element e with value v immediately before mark and returns e.
func (l *List[T]) InsertBefore(v T, mark *Element[T]) *Element[T] {
	return &Element[T]{innerElement: l.innerList.InsertBefore(v, mark.innerElement)}
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
func (l *List[T]) InsertAfter(v T, mark *Element[T]) *Element[T] {
	return &Element[T]{innerElement: l.innerList.InsertAfter(v, mark.innerElement)}
}

// MoveToFront moves element e to the front of list l.
func (l *List[T]) MoveToFront(e *Element[T]) {
	l.innerList.MoveToFront(e.innerElement)
}

// MoveToBack moves element e to the back of list l.
func (l *List[T]) MoveToBack(e *Element[T]) {
	l.innerList.MoveToBack(e.innerElement)
}

// MoveBefore moves element e to its new position before mark.
func (l *List[T]) MoveBefore(e, mark *Element[T]) {
	l.innerList.MoveBefore(e.innerElement, mark.innerElement)
}

// MoveAfter moves element e to its new position after mark.
func (l *List[T]) MoveAfter(e, mark *Element[T]) {
	l.innerList.MoveAfter(e.innerElement, mark.innerElement)
}

// PushBackList inserts a copy of an other list at the back of list l.
func (l *List[T]) PushBackList(other *List[T]) {
	l.innerList.PushBackList(other.innerList)
}

// PushFrontList inserts a copy of an other list at the front of list l.
func (l *List[T]) PushFrontList(other *List[T]) {
	l.innerList.PushFrontList(other.innerList)
}
