package set

import "advanced/ch00/undo"

type Set[T comparable] struct {
	data map[T]struct{}
	mgr  *undo.UndoManager
	err  error
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		data: make(map[T]struct{}),
		mgr:  undo.NewUndoManager(10),
	}
}

func (s *Set[T]) Add(x T) {
	if s.Contains(x) {
		s.mgr.Push(nil)
		return
	}

	s.mgr.Push(&AddCommand[T]{set: s, x: x})
	s.add(x)
}

func (s *Set[T]) Delete(x T) {
	if !s.Contains(x) {
		s.mgr.Push(nil)
		return
	}

	s.mgr.Push(&DeleteCommand[T]{set: s, x: x})
	s.delete(x)
}

func (s *Set[T]) add(x T) {
	s.data[x] = struct{}{}
}

func (s *Set[T]) delete(x T) {
	delete(s.data, x)
}

func (s *Set[T]) Contains(x T) bool {
	_, ok := s.data[x]
	return ok
}

func (s *Set[T]) Undo() {
	if err := s.mgr.Undo(); err != nil {
		s.err = err
	}
}

func (s *Set[T]) Size() int {
	return len(s.data)
}

func (s *Set[T]) Err() error {
	return s.err
}
