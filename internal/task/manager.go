package task

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type Manager struct {
	mu    sync.RWMutex
	tasks map[string]*Task
}

func NewManager() *Manager {
	return &Manager{
		tasks: make(map[string]*Task),
	}
}

func (m *Manager) CreateTask() *Task {
	id := uuid.New().String()
	task := &Task{
		ID:        id,
		CreatedAt: time.Now(),
		Status:    StatusPending,
	}

	m.mu.Lock()
	m.tasks[id] = task
	m.mu.Unlock()

	go m.runTask(task)
	return task
}

func (m *Manager) runTask(t *Task) {
	now := time.Now()
	m.mu.Lock()
	t.Status = StatusRunning
	t.StartedAt = &now
	m.mu.Unlock()

	time.Sleep(time.Duration(3) * time.Minute)

	done := time.Now()
	m.mu.Lock()
	t.Status = StatusCompleted
	t.FinishedAt = &done
	t.Result = "task completed successfully"
	m.mu.Unlock()
}

func (m *Manager) GetTask(id string) (*Task, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	task, exists := m.tasks[id]
	return task, exists
}

func (m *Manager) DeleteTask(id string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.tasks[id]; exists {
		delete(m.tasks, id)
		return true
	}
	return false
}
