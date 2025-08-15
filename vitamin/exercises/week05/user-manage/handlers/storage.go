package handlers

import (
	"encoding/json"
	"os"
	"sync"
)

var (
	users []User
	mu    sync.RWMutex
)

const filePath = "user.json"

func initStorage() {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if _, err := os.Create("user.json"); err != nil {
			panic(err)
		}
	}
	loadFromFile()
}

func loadFromFile() {
	data, err := os.ReadFile(filePath)
	if err != nil {
		users = make([]User, 0)
		return
	}
	if len(data) == 0 {
		users = make([]User, 0)
		return
	}
	if err := json.Unmarshal(data, &users); err != nil {
		users = make([]User, 0)
		return
	}
}

func saveToFile() error {
	data, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

func init() {
	initStorage()
}
