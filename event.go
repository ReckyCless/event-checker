package app

import (
	"errors"
	"time"
)

type Event struct {
	ID             int        `json:"-" db:"id"`
	Title          string     `json:"title" binding:"required" db:"title"`
	Description    string     `json:"description" binding:"required" db:"description"`
	Location       string     `json:"location" binding:"required" db:"location"`
	StartTime      time.Time  `json:"start_time" binding:"required" db:"start_time"`
	EndTime        *time.Time `json:"end_time" db:"end_time"`
	ImagePath      *string    `json:"image_path" db:"image_path"`
	UserID         int        `json:"user_id" db:"user_id"`
	UserName       *string    `json:"user_name" db:"user_name"`
	UserSurname    *string    `json:"user_surname" db:"user_surname"`
	UserPatronymic *string    `json:"user_patronymic" db:"user_patronymic"`
	UserEmail      *string    `json:"user_email" db:"user_email"`
	UserPhone      *string    `json:"user_phone" db:"user_phone"`
	TypeID         int        `json:"type_id" binding:"required" db:"type_id"`
	TypeName       *string    `json:"type_name" db:"type_name"`
	CreatedAt      *time.Time `json:"created_at" db:"created_at"`
}

type UpdateEventInput struct {
	Title       *string    `json:"title" db:"title"`
	Description *string    `json:"description" db:"description"`
	Location    *string    `json:"location" db:"location"`
	StartTime   *time.Time `json:"start_time" db:"start_time"`
	EndTime     *time.Time `json:"end_time" db:"end_time"`
	ImagePath   *string    `json:"image_path" db:"image_path"`
	TypeID      *int       `json:"type_id" db:"type_id"`
}

func (i UpdateEventInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Location == nil && i.ImagePath == nil && i.StartTime == nil && i.EndTime == nil && i.TypeID == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
