package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	app "github.com/reckycless/event-checker"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) UpdateUserRole(userID int, input app.UserRoleInput) (int64, error) {
	query := fmt.Sprintf(`UPDATE %s SET role_id=$1 WHERE id=$2`,
		usersTable)

	result, err := r.db.Exec(query, input.RoleID, userID)
	rowsAffected, errRows := result.RowsAffected()
	if errRows != nil {
		return 0, errRows
	}

	return rowsAffected, err
}

func (r *UserPostgres) UpdateUserOrganisator(userID int, input app.UserOrganisatorInput) (int64, error) {
	query := fmt.Sprintf(`UPDATE %s SET organisator_id=$1 WHERE id=$2`,
		usersTable)

	result, err := r.db.Exec(query, input.OrganisatorID, userID)
	rowsAffected, errRows := result.RowsAffected()
	if errRows != nil {
		return 0, errRows
	}

	return rowsAffected, err
}
