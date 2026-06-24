package repository

import "github.com/bv344/goalapp/internal/domain"

type ScheduleRepository interface {
	FindAll() ([]domain.ScheduleSlot, error)
	Create(slot *domain.ScheduleSlot) error
	Delete(id int64) error
}
