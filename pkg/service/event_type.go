package service

import (
	app "github.com/reckycless/event-checker"
	"github.com/reckycless/event-checker/pkg/repository"
)

type EventTypeService struct {
	repo repository.EventType
}

func NewEventTypeService(repo repository.EventType) *EventTypeService {
	return &EventTypeService{repo: repo}
}

func (s *EventTypeService) Create(eventType app.EventType) (int, error) {
	return s.repo.Create(eventType)
}

func (s *EventTypeService) GetAll() ([]app.EventType, error) {
	return s.repo.GetAll()
}

func (s *EventTypeService) GetByID(eventTypeID int) (app.EventType, error) {
	return s.repo.GetByID(eventTypeID)
}

func (s *EventTypeService) Delete(eventTypeID int) (int64, error) {
	return s.repo.Delete(eventTypeID)
}

func (s *EventTypeService) Update(eventTypeID int, input app.UpdateEventTypeInput) (int64, error) {
	return s.repo.Update(eventTypeID, input)
}
