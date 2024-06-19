package app

import (
	"errors"
	"time"
)

type Organisator struct {
	ID        int        `json:"-" db:"id"`
	Title     string     `json:"title" binding:"required" db:"title"`
	LogoPath  *string    `json:"logo_path" db:"logo_path"`
	SiteUrl   *string    `json:"site_url" db:"site_url"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type CreateOrganisatorInput struct {
	Title    string  `json:"title" binding:"required" db:"title"`
	LogoPath *string `json:"logo_path" db:"logo_path"`
	SiteUrl  *string `json:"site_url" db:"site_url"`
}

type GetOrganisatorOutput struct {
	ID        int        `json:"id" db:"id"`
	Title     string     `json:"title" db:"title"`
	LogoPath  *string    `json:"logo_path" db:"logo_path"`
	SiteUrl   *string    `json:"site_url" db:"site_url"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type UpdateOrganisatorInput struct {
	Title    *string `json:"title" db:"title"`
	LogoPath *string `json:"logo_path" db:"logo_path"`
	SiteUrl  *string `json:"site_url" db:"site_url"`
}

func (i UpdateOrganisatorInput) Validate() error {
	if i.Title == nil && i.LogoPath == nil && i.SiteUrl == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
