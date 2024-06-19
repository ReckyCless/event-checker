package repository

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	app "github.com/reckycless/event-checker"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user app.CreateUserInput) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (password_hash, name, surname, patronymic, birth_date, sex, phone, email, image_path, role_id, created_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id", usersTable)

	const role = 3
	row := r.db.QueryRow(query, user.Password, user.Name, user.Surname, user.Patronymic, user.BirthDate, user.Sex, user.Phone, user.Email, user.ImagePath, role, time.Now().UTC())
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(login, password string) (app.User, error) {
	var user app.User
	query := fmt.Sprintf("SELECT id, role_id FROM %s WHERE (email=$1 OR phone=$1) AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, login, password)

	return user, err
}
