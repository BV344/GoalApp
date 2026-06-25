package service

import (
	"github.com/bv344/goalapp/internal/domain"
	"github.com/bv344/goalapp/internal/repository"
)

type GoalService struct {
	repo repository.GoalRepository
}

func NewGoalService(repo repository.GoalRepository) *GoalService {
	return &GoalService{repo: repo}
}

func (s *GoalService) GetAll() ([]domain.Goal, error)          { return s.repo.FindAll() }
func (s *GoalService) GetByID(id int64) (*domain.Goal, error)  { return s.repo.FindByID(id) }
func (s *GoalService) Create(goal *domain.Goal) error           { return s.repo.Create(goal) }
func (s *GoalService) Update(goal *domain.Goal) error           { return s.repo.Update(goal) }
func (s *GoalService) Delete(id int64) error                    { return s.repo.Delete(id) }
