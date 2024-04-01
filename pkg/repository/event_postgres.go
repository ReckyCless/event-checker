package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	app "github.com/reckycless/event-checker"
	"github.com/sirupsen/logrus"
)

type EventPostgres struct {
	db *sqlx.DB
}

func NewEventPostgres(db *sqlx.DB) *EventPostgres {
	return &EventPostgres{db: db}
}

func (r *EventPostgres) Create(userID int, event app.Event) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createEventQuery := fmt.Sprintf(`INSERT INTO %s (
		title, 
		description, 
		location, 
		image_path, 
		start_time, 
		end_time,
		user_id,
		type_id,
		created_at)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id`, eventsTable)
	row := tx.QueryRow(
		createEventQuery,
		event.Title,
		event.Description,
		event.Location,
		event.ImagePath,
		event.StartTime,
		event.EndTime,
		userID,
		event.TypeID,
		time.Now().UTC())
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *EventPostgres) GetAll() ([]app.Event, error) {
	var events []app.Event
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.location, 
		ti.start_time, ti.end_time, ti.image_path,
		li.name as user_name, li.surname as user_surname, li.patronymic as user_patronymic, li.email as user_email, li.phone as user_phone,
		ul.name as type_name,
		ti.user_id, ti.type_id, ti.created_at
		FROM %s ti INNER JOIN %s li on ti.user_id = li.id INNER JOIN %s ul on ti.type_id = ul.id`, eventsTable, usersTable, eventsTypesTable)

	err := r.db.Select(&events, query)

	return events, err
}

func (r *EventPostgres) GetAllForUser(userID int) ([]app.Event, error) {
	var events []app.Event
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.location, 
			ti.start_time, ti.end_time, ti.image_path,
			li.name as user_name, li.surname as user_surname, li.patronymic as user_patronymic, li.email as user_email, li.phone as user_phone,
			ul.name as type_name,
			ti.user_id, ti.type_id, ti.created_at
		  	FROM %s ti INNER JOIN %s li on ti.user_id = li.id INNER JOIN %s ul on ti.type_id = ul.id
			WHERE ti.user_id = $1`, eventsTable, usersTable, eventsTypesTable)

	err := r.db.Select(&events, query, userID)

	return events, err
}

func (r *EventPostgres) GetByID(eventID int) (app.Event, error) {
	var event app.Event

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.location, 
		ti.start_time, ti.end_time, ti.image_path,
		li.name as user_name, li.surname as user_surname, li.patronymic as user_patronymic, li.email as user_email, li.phone as user_phone,
		ul.name as type_name,
		ti.user_id, ti.type_id, ti.created_at
		FROM %s ti INNER JOIN %s li on ti.user_id = li.id INNER JOIN %s ul on ti.type_id = ul.id
		WHERE ti.id = $1`, eventsTable, usersTable, eventsTypesTable)

	err := r.db.Get(&event, query, eventID)

	return event, err
}

func (r *EventPostgres) Delete(userID, eventID int) (int64, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 AND id=$2",
		eventsTable)

	result, err := r.db.Exec(query, userID, eventID)
	rowsAffected, errRows := result.RowsAffected()
	if errRows != nil {
		return 0, errRows
	}

	return rowsAffected, err
}

func (r *EventPostgres) Update(userID, eventID int, input app.UpdateEventInput) (int64, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, *input.Title)
		argID++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argID))
		args = append(args, *input.Description)
		argID++
	}
	if input.Location != nil {
		setValues = append(setValues, fmt.Sprintf("location=$%d", argID))
		args = append(args, *input.Location)
		argID++
	}
	if input.ImagePath != nil {
		setValues = append(setValues, fmt.Sprintf("image_path=$%d", argID))
		args = append(args, *input.ImagePath)
		argID++
	}
	if input.StartTime != nil {
		setValues = append(setValues, fmt.Sprintf("start_time=$%d", argID))
		args = append(args, *input.StartTime)
		argID++
	}
	if input.EndTime != nil {
		setValues = append(setValues, fmt.Sprintf("end_time=$%d", argID))
		args = append(args, *input.EndTime)
		argID++
	}
	if input.TypeID != nil {
		setValues = append(setValues, fmt.Sprintf("type_id=$%d", argID))
		args = append(args, *input.TypeID)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d AND user_id=$%d",
		eventsTable, setQuery, argID, argID+1)

	args = append(args, eventID, userID)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	result, err := r.db.Exec(query, args...)
	rowsAffected, errRows := result.RowsAffected()
	if errRows != nil {
		return 0, errRows
	}

	return rowsAffected, err
}
