package handlers

import "errors"

var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
)

func getAllUsers() []User {
	mu.RLock()
	defer mu.RUnlock()
	return append([]User(nil), users...) // 返回副本
}

func findIndexByEmail(email string) int {
	for i, u := range users {
		if u.Email == email {
			return i
		}
	}
	return -1
}

func addUser(u User) error {
	mu.Lock()
	defer mu.Unlock()
	if findIndexByEmail(u.Email) != -1 {
		return ErrUserExists
	}
	users = append(users, u)
	return saveToFile()
}

func updateUserByEmail(u User) error {
	mu.Lock()
	defer mu.Unlock()
	i := findIndexByEmail(u.Email)
	if i == -1 {
		return ErrUserNotFound
	}
	if u.Name != "" {
		users[i].Name = u.Name
	}
	if u.Age != 0 {
		users[i].Age = u.Age
	}
	return saveToFile()
}

func deleteUserByEmail(email string) error {
	mu.Lock()
	defer mu.Unlock()
	i := findIndexByEmail(email)
	if i == -1 {
		return ErrUserNotFound
	}
	users = append(users[:i], users[i+1:]...)
	return saveToFile()
}
