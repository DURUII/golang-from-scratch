package undo

import "errors"

var ErrNoUndo = errors.New("no functions to undo")

type Command interface {
	Undo()
}

// 命令管理器
type UndoManager struct {
	stack []Command
}

func NewUndoManager(cap int) *UndoManager {
	return &UndoManager{
		stack: make([]Command, 0, cap),
	}
}

func (m *UndoManager) Add(cmd Command) {
	m.stack = append(m.stack, cmd)
}

func (m *UndoManager) Push(cmd Command) {
	m.stack = append(m.stack, cmd)
}

func (m *UndoManager) Undo() error {
	if len(m.stack) == 0 {
		return ErrNoUndo
	}
	idx := len(m.stack) - 1
	cmd := m.stack[idx]
	if cmd != nil {
		cmd.Undo()
	}
	m.stack = m.stack[:idx]
	return nil
}
