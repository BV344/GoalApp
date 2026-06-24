package service

import (
	"github.com/bv344/goalapp/internal/domain"
	"github.com/bv344/goalapp/internal/repository"
)

type ScheduleService struct {
	repo repository.ScheduleRepository
}

func NewScheduleService(repo repository.ScheduleRepository) *ScheduleService {
	return &ScheduleService{repo: repo}
}

func (s *ScheduleService) GetAll() ([]domain.ScheduleSlot, error)        { return s.repo.FindAll() }
func (s *ScheduleService) Create(slot *domain.ScheduleSlot) error         { return s.repo.Create(slot) }
func (s *ScheduleService) Delete(id int64) error                          { return s.repo.Delete(id) }
