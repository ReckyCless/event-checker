package service

import (
	app "github.com/reckycless/event-checker"
	"github.com/reckycless/event-checker/pkg/repository"
)

type EventService struct {
	repo repository.Event
}

func NewEventService(repo repository.Event) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) Create(userID int, event app.CreateEventInput) (int, error) {
	return s.repo.Create(userID, event)
}

func (s *EventService) GetAll() ([]app.GetEventOutput, error) {
	return s.repo.GetAll()
}

func (s *EventService) GetAllForUser(userID int) ([]app.GetEventOutput, error) {
	return s.repo.GetAllForUser(userID)
}

func (s *EventService) GetByID(eventID int) (app.GetEventOutput, error) {
	return s.repo.GetByID(eventID)
}

func (s *EventService) Update(userID, eventID int, input app.UpdateEventInput) (int64, error) {
	if err := input.Validate(); err != nil {
		return 0, err
	}
	return s.repo.Update(userID, eventID, input)
}

func (s *EventService) Delete(userID, eventID int) (int64, error) {
	return s.repo.Delete(userID, eventID)
}
