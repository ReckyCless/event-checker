package app

import (
	"errors"
	"time"
)

type Visitor struct {
	ID             int        `json:"-" db:"id"`
	EventID        int        `json:"event_id" db:"event_id"`
	Name           *string    `json:"name" db:"name" binding:"required"`
	Surname        *string    `json:"surname" db:"surname" binding:"required"`
	Patronymic     *string    `json:"patronymic" db:"patronymic" binding:"required"`
	BirthDate      *time.Time `json:"birth_date" db:"birth_date" binding:"required"`
	Email          *string    `json:"email" db:"email"`
	Phone          *string    `json:"phone" db:"phone"`
	IsVisited      bool       `json:"is_visited" db:"is_visited"`
	UserName       *string    `json:"user_name" db:"user_name"`
	UserSurname    *string    `json:"user_surname" db:"user_surname"`
	UserPatronymic *string    `json:"user_patronymic" db:"user_patronymic"`
	UserBirthDate  *time.Time `json:"user_birth_date" db:"user_birth_date"`
	UserEmail      *string    `json:"user_email" db:"user_email"`
	UserPhone      *string    `json:"user_phone" db:"user_phone"`
	CreatedAt      *time.Time `json:"created_at" db:"created_at"`
}

type UpdateVisitorInput struct {
	Name       *string    `json:"name" db:"name"`
	Surname    *string    `json:"surname" db:"surname"`
	Patronymic *string    `json:"patronymic" db:"patronymic"`
	BirthDate  *time.Time `json:"birth_date" db:"birth_date"`
	Email      *string    `json:"email" db:"email"`
	Phone      *string    `json:"phone" db:"phone"`
	IsVisited  *bool      `json:"is_visited" db:"is_visited"`
}

func (i UpdateVisitorInput) Validate() error {
	if i.Name == nil && i.Surname == nil && i.Patronymic == nil && i.BirthDate == nil && i.Email == nil && i.Phone == nil && i.IsVisited == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
