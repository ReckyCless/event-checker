package service

import (
	app "github.com/reckycless/event-checker"
	"github.com/reckycless/event-checker/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user app.CreateUserInput) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, int, error)
}

type Event interface {
	Create(userID int, event app.CreateEventInput) (int, error)
	GetAll() ([]app.GetEventOutput, error)
	GetAllForUser(userID int) ([]app.GetEventOutput, error)
	GetByID(eventID int) (app.GetEventOutput, error)
	Update(userID, eventID int, input app.UpdateEventInput) (int64, error)
	Delete(userID, eventID int) (int64, error)
}

type Visitor interface {
	CreateAsUser(userID, eventID int) (int, error)
	Create(eventID int, visitor app.CreateVisitorInput) (int, error)
	GetAllEventVisitors(userID, eventID int) ([]app.GetVisitorOutput, error)
	GetByID(userID, visitorID int) (app.GetVisitorOutput, error)
	Delete(userID, visitorID int) (int64, error)
	Update(userID, visitorID int, input app.UpdateVisitorInput) (int64, error)
}

type Organisator interface {
	Create(visitor app.CreateOrganisatorInput) (int, error)
	GetAll() ([]app.GetOrganisatorOutput, error)
	GetByID(organisatorID int) (app.GetOrganisatorOutput, error)
	Delete(organisatorID int) (int64, error)
	Update(organisatorID int, input app.UpdateOrganisatorInput) (int64, error)
}

type EventType interface {
	Create(eventType app.CreateEventTypeInput) (int, error)
	GetAll() ([]app.GetEventTypeOutput, error)
	GetByID(eventTypeID int) (app.GetEventTypeOutput, error)
	Delete(eventTypeID int) (int64, error)
	Update(eventTypeID int, input app.UpdateEventTypeInput) (int64, error)
}

type User interface {
	UpdateUserRole(userID int, input app.UserRoleInput) (int64, error)
	UpdateUserOrganisator(userID int, input app.UserOrganisatorInput) (int64, error)
}

type Role interface {
	GetAll() ([]app.Role, error)
	GetByID(roleID int) (app.Role, error)
	Update(roleID int, input app.UpdateRoleInput) (int64, error)
}

type Service struct {
	Authorization
	Event
	Visitor
	Organisator
	EventType
	User
	Role
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Event:         NewEventService(repos.Event),
		Visitor:       NewVisitorService(repos.Visitor, repos.Event),
		Organisator:   NewOrganisatorService(repos.Organisator),
		EventType:     NewEventTypeService(repos.EventType),
		User:          NewUserService(repos.User),
		Role:          NewRoleService(repos.Role),
	}
}
