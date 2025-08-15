package model

type Manageable interface {
	Borrow() bool
	Return() bool
	GetInfo() string
}
