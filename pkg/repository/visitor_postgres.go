package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	app "github.com/reckycless/event-checker"
	"github.com/sirupsen/logrus"
)

type VisitorPostgres struct {
	db *sqlx.DB
}

func NewVisitorPostgres(db *sqlx.DB) *VisitorPostgres {
	return &VisitorPostgres{db: db}
}

// Creating the visitor/Siging up for the event (Authorized)
func (r *VisitorPostgres) CreateAsUser(userID, eventID int) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	const is_visited = false
	var visitorID int
	createVisitorQuery := fmt.Sprintf(`INSERT INTO %s (
		is_visited, 
		user_id, 
		event_id, 
		created_at)
		VALUES($1, $2, $3, $4)
		RETURNING id`, eventsVisitorsTable)
	row := tx.QueryRow(
		createVisitorQuery,
		is_visited,
		userID,
		eventID,
		time.Now().UTC())
	if err := row.Scan(&visitorID); err != nil {
		tx.Rollback()
		return 0, err
	}

	return visitorID, tx.Commit()
}

// Creating the visitor/Siging up for the event (NON-Authorized)
func (r *VisitorPostgres) Create(eventID int, visitor app.Visitor) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	const is_visited = false
	var visitorID int
	createVisitorQuery := fmt.Sprintf(`INSERT INTO %s (
		is_visited, 
		event_id,
		name,
		surname, 
		patronymic,
		birth_date,
		phone,
		email, 
		created_at)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id`, eventsVisitorsTable)
	row := tx.QueryRow(
		createVisitorQuery,
		is_visited,
		eventID,
		visitor.Name,
		visitor.Surname,
		visitor.Patronymic,
		visitor.BirthDate,
		visitor.Phone,
		visitor.Email,
		time.Now().UTC())
	if err := row.Scan(&visitorID); err != nil {
		tx.Rollback()
		return 0, err
	}

	return visitorID, tx.Commit()
}

func (r *VisitorPostgres) GetAllEventVisitors(userID, eventID int) ([]app.Visitor, error) {
	var visitors []app.Visitor
	// TODO: CHECK WITH NORMAL DATE IN USERS SOMEDAY AND CHANGE ti.birth_date To li.birth_date
	query := fmt.Sprintf(`SELECT 
	ti.name, ti.surname, ti.patronymic, ti.birth_date, ti.phone, ti.email, ti.is_visited, ti.created_at,
	li.name as user_name, li.surname as user_surname, li.patronymic as user_patronymic, 
	ti.birth_date as user_birth_date, li.email as user_email, li.phone as user_phone
	FROM %s ti INNER JOIN %s ul on ul.id = ti.event_id LEFT JOIN %s li on li.id = ti.user_id WHERE ul.user_id = $1 AND ti.event_id = $2`, eventsVisitorsTable, eventsTable, usersTable)
	if err := r.db.Select(&visitors, query, userID, eventID); err != nil {
		return nil, err
	}

	return visitors, nil
}

func (r *VisitorPostgres) GetByID(userID, visitorID int) (app.Visitor, error) {
	var visitor app.Visitor
	// TODO: CHECK WITH NORMAL DATE IN USERS SOMEDAY AND CHANGE ti.birth_date To li.birth_date
	query := fmt.Sprintf(`SELECT 
	ti.name, ti.surname, ti.patronymic, ti.birth_date, ti.phone, ti.email, ti.is_visited, ti.created_at,
	li.name as user_name, li.surname as user_surname, li.patronymic as user_patronymic, 
	ti.birth_date as user_birth_date, li.email as user_email, li.phone as user_phone
	FROM %s ti INNER JOIN %s ul on ul.id = ti.event_id LEFT JOIN %s li on li.id = ti.user_id WHERE ul.user_id = $1 AND ti.id = $2`, eventsVisitorsTable, eventsTable, usersTable)
	if err := r.db.Get(&visitor, query, userID, visitorID); err != nil {
		return visitor, err
	}

	return visitor, nil
}

func (r *VisitorPostgres) Delete(userID, visitorID int) (int64, error) {
	query := fmt.Sprintf("DELETE FROM %s ti USING %s ul WHERE ul.id = ti.event_id AND ul.user_id=$1 AND ti.id=$2",
		eventsVisitorsTable, eventsTable)

	result, err := r.db.Exec(query, userID, visitorID)
	rowsAffected, errRows := result.RowsAffected()
	if errRows != nil {
		return 0, errRows
	}

	return rowsAffected, err
}

func (r *VisitorPostgres) Update(userID, visitorID int, input app.UpdateVisitorInput) (int64, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argID))
		args = append(args, *input.Name)
		argID++
	}
	if input.Surname != nil {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argID))
		args = append(args, *input.Surname)
		argID++
	}
	if input.Patronymic != nil {
		setValues = append(setValues, fmt.Sprintf("patronymic=$%d", argID))
		args = append(args, *input.Patronymic)
		argID++
	}
	if input.BirthDate != nil {
		setValues = append(setValues, fmt.Sprintf("birth_date=$%d", argID))
		args = append(args, *input.BirthDate)
		argID++
	}
	if input.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argID))
		args = append(args, *input.Email)
		argID++
	}
	if input.Phone != nil {
		setValues = append(setValues, fmt.Sprintf("phone=$%d", argID))
		args = append(args, *input.Phone)
		argID++
	}
	if input.IsVisited != nil {
		setValues = append(setValues, fmt.Sprintf("is_visited=$%d", argID))
		args = append(args, *input.IsVisited)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul
			WHERE ti.event_id = ul.id AND ul.user_id = li.id
			AND ti.id=$%d AND ul.user_id=$%d AND ti.user_id IS NULL`,
		eventsVisitorsTable, setQuery, usersTable, eventsTable, argID, argID+1)

	args = append(args, visitorID, userID)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	result, err := r.db.Exec(query, args...)
	rowsAffected, errRows := result.RowsAffected()
	if errRows != nil {
		return 0, errRows
	}

	return rowsAffected, err
}
