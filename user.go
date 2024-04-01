package app

import (
	"time"
)

type User struct {
	ID            int        `json:"-" db:"id"`
	Password      string     `json:"password" binding:"required" db:"password"`
	Name          string     `json:"name" binding:"required" db:"name"`
	Surname       string     `json:"surname" binding:"required" db:"surname"`
	Patronymic    string     `json:"patronymic" db:"patronymic"`
	BirthDate     time.Time  `json:"birth_date" db:"birth_date"`
	Phone         *string    `json:"phone" db:"phone"`
	Email         string     `json:"email" db:"email"`
	ImagePath     *string    `json:"image_path"  db:"image_path"`
	RoleID        int        `json:"role_id" binding:"required" db:"role_id"`
	OrganisatorID *int       `json:"organisator_id" db:"organisator_id"`
	CreatedAt     *time.Time `json:"created_at" db:"created_at"`
}

type UserPersonalInput struct {
	Name       *string    `json:"name" db:"name"`
	Surname    *string    `json:"surname" db:"surname"`
	Patronymic *string    `json:"patronymic" db:"patronymic"`
	BirthDate  *time.Time `json:"birth_date" db:"birth_date"`
	Phone      *string    `json:"phone" db:"phone"`
	ImagePath  *string    `json:"image_path"  db:"image_path"`
}

type UserPasswordInput struct {
	OldPassword string `json:"old_password" binding:"required"`
	Password    string `json:"password" binding:"required" db:"password"`
}

type UserRoleInput struct {
	RoleID int `json:"role_id" binding:"required" db:"role_id"`
}

type UserOrganisatorInput struct {
	OrganisatorID *int `json:"organisator_id" db:"organisator_id"`
}

type UserEmailInput struct {
	Password string `json:"password" binding:"required" db:"password"`
	Email    string `json:"email" binding:"required" db:"email"`
}

type Role struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" binding:"required" db:"name"`
}
