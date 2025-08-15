package model

import "fmt"

type Book struct {
	ID          int
	Title       string
	Author      string
	IsAvailable bool
}

func (b *Book) Borrow() bool {
	if b.IsAvailable {
		b.IsAvailable = false
		return true
	}
	return false
}

func (b *Book) Return() bool {
	if !b.IsAvailable {
		b.IsAvailable = true
		return true
	}
	return false
}

func (b *Book) GetInfo() string {
	return fmt.Sprintf("ID: %d, Title: %s, Author: %s, Available: %v", b.ID, b.Title, b.Author, b.IsAvailable)
}
