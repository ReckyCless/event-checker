package app

import (
	"errors"
	"time"
)

type Event struct {
	ID          int        `json:"-" db:"id"`
	Title       string     `json:"title" binding:"required" db:"title"`
	Description string     `json:"description" binding:"required" db:"description"`
	Location    string     `json:"location" binding:"required" db:"location"`
	StartTime   time.Time  `json:"start_time" binding:"required" db:"start_time"`
	EndTime     *time.Time `json:"end_time" db:"end_time"`
	ImagePath   *string    `json:"image_path" db:"image_path"`
	CreatorID   int        `json:"user_id" db:"user_id"`
	TypeID      int        `json:"type_id" binding:"required" db:"type_id"`
	CreatedAt   *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
}

type CreateEventInput struct {
	Title       string     `json:"title" binding:"required" db:"title"`
	Description string     `json:"description" binding:"required" db:"description"`
	Location    string     `json:"location" binding:"required" db:"location"`
	StartTime   time.Time  `json:"start_time" binding:"required" db:"start_time"`
	EndTime     *time.Time `json:"end_time" db:"end_time"`
	ImagePath   *string    `json:"image_path" db:"image_path"`
	CreatorID   int        `json:"creator_id" db:"creator_id"`
	TypeID      int        `json:"type_id" binding:"required" db:"type_id"`
}

type GetEventOutput struct {
	Title               string     `json:"title" db:"title"`
	Description         string     `json:"description" db:"description"`
	TypeName            string     `json:"type_name" db:"type_name"`
	Location            string     `json:"location" db:"location"`
	StartTime           time.Time  `json:"start_time" db:"start_time"`
	EndTime             *time.Time `json:"end_time" db:"end_time"`
	ImagePath           *string    `json:"image_path" db:"image_path"`
	CreatorOrganisation *string    `json:"creator_organisation" db:"creator_organisation"`
	CreatorName         *string    `json:"creator_name" db:"creator_name"`
	CreatorSurname      *string    `json:"creator_surname" db:"creator_surname"`
	CreatorPatronymic   *string    `json:"creator_patronymic" db:"creator_patronymic"`
	CreatorEmail        *string    `json:"creator_email" db:"creator_email"`
	CreatorPhone        *string    `json:"creator_phone" db:"creator_phone"`
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
