package service

import (
	app "github.com/reckycless/event-checker"
	"github.com/reckycless/event-checker/pkg/repository"
)

type VisitorService struct {
	repo      repository.Visitor
	eventRepo repository.Event
}

func NewVisitorService(repo repository.Visitor, eventRepo repository.Event) *VisitorService {
	return &VisitorService{repo: repo, eventRepo: eventRepo}
}

func (s *VisitorService) CreateAsUser(userID, eventID int) (int, error) {
	_, err := s.eventRepo.GetByID(eventID)
	if err != nil {
		//event does not exsists
		return 0, err
	}

	return s.repo.CreateAsUser(userID, eventID)
}

func (s *VisitorService) Create(eventID int, visitor app.Visitor) (int, error) {
	_, err := s.eventRepo.GetByID(eventID)
	if err != nil {
		//event does not exsists
		return 0, err
	}

	return s.repo.Create(eventID, visitor)
}

func (s *VisitorService) GetAllEventVisitors(userID, eventID int) ([]app.Visitor, error) {
	return s.repo.GetAllEventVisitors(userID, eventID)
}

func (s *VisitorService) GetByID(userID, visitorID int) (app.Visitor, error) {
	return s.repo.GetByID(userID, visitorID)
}

func (s *VisitorService) Delete(userID, visitorID int) (int64, error) {
	return s.repo.Delete(userID, visitorID)
}

func (s *VisitorService) Update(userID, visitorID int, input app.UpdateVisitorInput) (int64, error) {
	return s.repo.Update(userID, visitorID, input)
}
