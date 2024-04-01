package repository

import (
	"github.com/jmoiron/sqlx"
	app "github.com/reckycless/event-checker"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
	GetUser(login, password string) (app.User, error)
}

type Event interface {
	Create(userID int, event app.Event) (int, error)
	GetAll() ([]app.Event, error)
	GetAllForUser(userID int) ([]app.Event, error)
	GetByID(eventID int) (app.Event, error)
	Delete(userID, eventID int) (int64, error)
	Update(userID, eventID int, input app.UpdateEventInput) (int64, error)
}

type Visitor interface {
	CreateAsUser(userID, eventID int) (int, error)
	Create(eventID int, visitor app.Visitor) (int, error)
	GetAllEventVisitors(userID, eventID int) ([]app.Visitor, error)
	GetByID(userID, visitorID int) (app.Visitor, error)
	Delete(userID, visitorID int) (int64, error)
	Update(userID, visitorID int, input app.UpdateVisitorInput) (int64, error)
}

type Organisator interface {
	Create(visitor app.Organisator) (int, error)
	GetAll() ([]app.Organisator, error)
	GetByID(organisatorID int) (app.Organisator, error)
	Delete(organisatorID int) (int64, error)
	Update(organisatorID int, input app.UpdateOrganisatorInput) (int64, error)
}

type EventType interface {
	Create(eventType app.EventType) (int, error)
	GetAll() ([]app.EventType, error)
	GetByID(eventTypeID int) (app.EventType, error)
	Delete(eventTypeID int) (int64, error)
	Update(eventTypeID int, input app.UpdateEventTypeInput) (int64, error)
}

type User interface {
	UpdateUserRole(userID int, input app.UserRoleInput) (int64, error)
	UpdateUserOrganisator(userID int, input app.UserOrganisatorInput) (int64, error)
}

type Repository struct {
	Authorization
	Event
	Visitor
	Organisator
	EventType
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Event:         NewEventPostgres(db),
		Visitor:       NewVisitorPostgres(db),
		Organisator:   NewOrganisatorPostgres(db),
		EventType:     NewEventTypePostgres(db),
		User:          NewUserPostgres(db),
	}
}
