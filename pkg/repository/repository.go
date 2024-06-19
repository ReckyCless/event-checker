package repository

import (
	"github.com/jmoiron/sqlx"
	app "github.com/reckycless/event-checker"
)

type Authorization interface {
	CreateUser(user app.CreateUserInput) (int, error)
	GetUser(login, password string) (app.User, error)
}

type Event interface {
	Create(userID int, event app.CreateEventInput) (int, error)
	GetAll() ([]app.GetEventOutput, error)
	GetAllForUser(userID int) ([]app.GetEventOutput, error)
	GetByID(eventID int) (app.GetEventOutput, error)
	Delete(userID, eventID int) (int64, error)
	Update(userID, eventID int, input app.UpdateEventInput) (int64, error)
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
	Create(organisator app.CreateOrganisatorInput) (int, error)
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

type Repository struct {
	Authorization
	Event
	Visitor
	Organisator
	EventType
	User
	Role
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Event:         NewEventPostgres(db),
		Visitor:       NewVisitorPostgres(db),
		Organisator:   NewOrganisatorPostgres(db),
		EventType:     NewEventTypePostgres(db),
		User:          NewUserPostgres(db),
		Role:          NewRolePostgres(db),
	}
}
