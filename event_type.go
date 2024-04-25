package app

import (
	"errors"
	"time"
)

type EventType struct {
	ID        int        `json:"id" db:"id"`
	Name      string     `json:"name" binding:"required" db:"name"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"created_at"`
}

type CreateEventTypeInput struct {
	Name string `json:"name" binding:"required" db:"name"`
}

type GetEventTypeOutput struct {
	Name      string     `json:"name" db:"name"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"created_at"`
}

type UpdateEventTypeInput struct {
	Name *string `json:"title" db:"title"`
}

func (i UpdateEventTypeInput) Validate() error {
	if i.Name == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
