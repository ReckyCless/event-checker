package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	app "github.com/reckycless/event-checker"
	"github.com/sirupsen/logrus"
)

type EventTypePostgres struct {
	db *sqlx.DB
}

func NewEventTypePostgres(db *sqlx.DB) *EventTypePostgres {
	return &EventTypePostgres{db: db}
}

func (r *EventTypePostgres) Create(eventType app.EventType) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var eventTypeID int
	createQuery := fmt.Sprintf(`INSERT INTO %s (
		name,
		created_at)
		VALUES($1, $2)
		RETURNING id`, eventsTypesTable)
	row := tx.QueryRow(
		createQuery,
		eventType.Name,
		time.Now().UTC())
	if err := row.Scan(&eventTypeID); err != nil {
		tx.Rollback()
		return 0, err
	}

	return eventTypeID, tx.Commit()
}

func (r *EventTypePostgres) GetAll() ([]app.EventType, error) {
	var eventTypes []app.EventType
	query := fmt.Sprintf(`SELECT * FROM %s`, eventsTypesTable)
	if err := r.db.Select(&eventTypes, query); err != nil {
		return nil, err
	}

	return eventTypes, nil
}

func (r *EventTypePostgres) GetByID(eventTypeID int) (app.EventType, error) {
	var eventType app.EventType
	query := fmt.Sprintf(`SELECT *
	FROM %s WHERE id=$1`, eventsTypesTable)
	if err := r.db.Get(&eventType, query, eventTypeID); err != nil {
		return eventType, err
	}

	return eventType, nil
}

func (r *EventTypePostgres) Delete(eventTypeID int) (int64, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1",
		eventsTypesTable)

	result, err := r.db.Exec(query, eventTypeID)
	rowsAffected, errRows := result.RowsAffected()
	if errRows != nil {
		return 0, errRows
	}

	return rowsAffected, err
}

func (r *EventTypePostgres) Update(eventTypeID int, input app.UpdateEventTypeInput) (int64, error) {
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
		eventsTypesTable, setQuery, argID)

	args = append(args, eventTypeID)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	result, err := r.db.Exec(query, args...)
	rowsAffected, errRows := result.RowsAffected()
	if errRows != nil {
		return 0, errRows
	}

	return rowsAffected, err
}
