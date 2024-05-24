package tools

import (
	"fmt"
	"sync"
)

type MockDB struct {
	BestTimes []BestTime
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

func InitMockDB() *MockDB {
	return &MockDB{
		BestTimes: make([]BestTime, 0),
	}
}
