package tools

import (
	"errors"
	"fmt"
	"sync"
)

type MockDB struct {
	BestTimes []BestTime
	Users     map[string]User
	mu        sync.Mutex
}

func (m *MockDB) Create(bt BestTime) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.BestTimes = append(m.BestTimes, bt)
	return nil
}

func (m *MockDB) GetAll() ([]BestTime, error) {
	if len(m.BestTimes) == 0 {
		return nil, fmt.Errorf("no records in table")
	}
	return m.BestTimes, nil
}

func (m *MockDB) CreateUser(user User) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.Users[user.Username]; exists {
		return errors.New("user already exists")
	}
	m.Users[user.Username] = user
	return nil
}

func (m *MockDB) GetUser(username string) (User, error) {
	user, exists := m.Users[username]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

func InitMockDB() *MockDB {
	return &MockDB{
		BestTimes: make([]BestTime, 0),
		Users:     make(map[string]User),
	}
}
