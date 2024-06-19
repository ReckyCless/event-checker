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

func (r *EventPostgres) Create(userID int, event app.CreateEventInput) (int, error) {
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
		creator_id,
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

func (r *EventPostgres) GetAll() ([]app.GetEventOutput, error) {
	var events []app.GetEventOutput
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.location, ti.type_id,
		ti.start_time, ti.end_time, ti.image_path,
		li.name as creator_name, li.surname as creator_surname, li.patronymic as creator_patronymic, li.email as creator_email, li.phone as creator_phone,
		mt.title as organisation_title, mt.logo_path as organisation_logo, mt.site_url as organisation_site_url,
		ul.name as type_name,
		ti.created_at, ti.updated_at
		FROM %s ti INNER JOIN %s li on ti.creator_id = li.id INNER JOIN %s ul on ti.type_id = ul.id LEFT JOIN %s mt on li.organisator_id = mt.id`,
		eventsTable, usersTable, eventsTypesTable, organisatorsTable)

	err := r.db.Select(&events, query)

	return events, err
}

func (r *EventPostgres) GetAllForUser(userID int) ([]app.GetEventOutput, error) {
	var events []app.GetEventOutput
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.location, ti.type_id,
		ti.start_time, ti.end_time, ti.image_path,
		li.name as creator_name, li.surname as creator_surname, li.patronymic as creator_patronymic, li.email as creator_email, li.phone as creator_phone,
		mt.title as organisation_title, mt.logo_path as organisation_logo, mt.site_url as organisation_site_url,
		ul.name as type_name,
		ti.created_at, ti.updated_at
		FROM %s ti INNER JOIN %s li on ti.creator_id = li.id INNER JOIN %s ul on ti.type_id = ul.id LEFT JOIN %s mt on li.organisator_id = mt.id
		WHERE ti.creator_id = $1`,
		eventsTable, usersTable, eventsTypesTable, organisatorsTable)

	err := r.db.Select(&events, query, userID)

	return events, err
}

func (r *EventPostgres) GetByID(eventID int) (app.GetEventOutput, error) {
	var event app.GetEventOutput
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.location, ti.type_id,
		ti.start_time, ti.end_time, ti.image_path,
		li.name as creator_name, li.surname as creator_surname, li.patronymic as creator_patronymic, li.email as creator_email, li.phone as creator_phone,
		mt.title as organisation_title, mt.logo_path as organisation_logo, mt.site_url as organisation_site_url,
		ul.name as type_name,
		ti.created_at, ti.updated_at
		FROM %s ti INNER JOIN %s li on ti.creator_id = li.id INNER JOIN %s ul on ti.type_id = ul.id LEFT JOIN %s mt on li.organisator_id = mt.id
		WHERE ti.id = $1`, eventsTable, usersTable, eventsTypesTable, organisatorsTable)

	err := r.db.Get(&event, query, eventID)

	return event, err
}

func (r *EventPostgres) Delete(userID, eventID int) (int64, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE creator_id=$1 AND id=$2",
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
	// Set UpdatedAt Time
	setValues = append(setValues, fmt.Sprintf("updated_at=$%d", argID))
	args = append(args, time.Now().UTC())
	argID++

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d AND creator_id=$%d",
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
