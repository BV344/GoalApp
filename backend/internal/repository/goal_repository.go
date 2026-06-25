package repository

import "github.com/bv344/goalapp/internal/domain"

type GoalRepository interface {
	FindAll() ([]domain.Goal, error)
	FindByID(id int64) (*domain.Goal, error)
	Create(goal *domain.Goal) error
	Update(goal *domain.Goal) error
	Delete(id int64) error
}
