package mocks

import "github.com/bv344/goalapp/internal/domain"

// GoalRepositoryMock implements repository.GoalRepository for unit tests.
type GoalRepositoryMock struct {
	Goals []domain.Goal
	Err   error
}

func (m *GoalRepositoryMock) FindAll() ([]domain.Goal, error)         { return m.Goals, m.Err }
func (m *GoalRepositoryMock) FindByID(id int64) (*domain.Goal, error) { return nil, m.Err }
func (m *GoalRepositoryMock) Create(goal *domain.Goal) error           { return m.Err }
func (m *GoalRepositoryMock) Update(goal *domain.Goal) error           { return m.Err }
func (m *GoalRepositoryMock) Delete(id int64) error                    { return m.Err }
