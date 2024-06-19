package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	app "github.com/reckycless/event-checker"
	"github.com/sirupsen/logrus"
)

type RolePostgres struct {
	db *sqlx.DB
}

func NewRolePostgres(db *sqlx.DB) *RolePostgres {
	return &RolePostgres{db: db}
}

func (r *RolePostgres) GetAll() ([]app.Role, error) {
	var roles []app.Role
	query := fmt.Sprintf(`SELECT * FROM %s`, rolesTable)
	if err := r.db.Select(&roles, query); err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *RolePostgres) GetByID(roleID int) (app.Role, error) {
	var role app.Role
	query := fmt.Sprintf(`SELECT *
	FROM %s WHERE id=$1`, rolesTable)
	if err := r.db.Get(&role, query, roleID); err != nil {
		return role, err
	}

	return role, nil
}

func (r *RolePostgres) Update(roleID int, input app.UpdateRoleInput) (int64, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argID))
		args = append(args, *input.Name)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s SET %s
			WHERE id=$%d`,
		rolesTable, setQuery, argID)

	args = append(args, roleID)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	result, err := r.db.Exec(query, args...)
	rowsAffected, errRows := result.RowsAffected()
	if errRows != nil {
		return 0, errRows
	}

	return rowsAffected, err
}
