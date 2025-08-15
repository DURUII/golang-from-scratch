package model

import "fmt"

type Magazine struct {
	ID          int
	Title       string
	Issue       int
	IsAvailable bool
}

func (m *Magazine) Borrow() bool {
	if m.IsAvailable {
		m.IsAvailable = false
		return true
	}
	return false
}

func (m *Magazine) Return() bool {
	if !m.IsAvailable {
		m.IsAvailable = true
		return true
	}
	return false
}

func (m *Magazine) GetInfo() string {
	return fmt.Sprintf("ID: %d, Title: %s, Issue: %d, Available: %v", m.ID, m.Title, m.Issue, m.IsAvailable)
}
