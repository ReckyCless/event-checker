package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	app "github.com/reckycless/event-checker"
	"github.com/sirupsen/logrus"
)

type OrganisatorPostgres struct {
	db *sqlx.DB
}

func NewOrganisatorPostgres(db *sqlx.DB) *OrganisatorPostgres {
	return &OrganisatorPostgres{db: db}
}

func (r *OrganisatorPostgres) Create(organisator app.CreateOrganisatorInput) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var organisatorID int
	createOrganisatorQuery := fmt.Sprintf(`INSERT INTO %s (
		title,
		logo_path,
		site_url
		created_at)
		VALUES($1, $2, $3, $4)
		RETURNING id`, organisatorsTable)
	row := tx.QueryRow(
		createOrganisatorQuery,
		organisator.Title,
		organisator.LogoPath,
		organisator.SiteUrl,
		time.Now().UTC())
	if err := row.Scan(&organisatorID); err != nil {
		tx.Rollback()
		return 0, err
	}

	return organisatorID, tx.Commit()
}

func (r *OrganisatorPostgres) GetAll() ([]app.GetOrganisatorOutput, error) {
	var organisators []app.GetOrganisatorOutput
	query := fmt.Sprintf(`SELECT *
	FROM %s`, organisatorsTable)
	if err := r.db.Select(&organisators, query); err != nil {
		return nil, err
	}

	return organisators, nil
}

func (r *OrganisatorPostgres) GetByID(organisatorID int) (app.GetOrganisatorOutput, error) {
	var organisator app.GetOrganisatorOutput
	query := fmt.Sprintf(`SELECT *
	FROM %s WHERE id=$1`, organisatorsTable)
	if err := r.db.Get(&organisator, query, organisatorID); err != nil {
		return organisator, err
	}

	return organisator, nil
}

func (r *OrganisatorPostgres) Delete(organisatorID int) (int64, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1",
		organisatorsTable)

	result, err := r.db.Exec(query, organisatorID)
	rowsAffected, errRows := result.RowsAffected()
	if errRows != nil {
		return 0, errRows
	}

	return rowsAffected, err
}

func (r *OrganisatorPostgres) Update(organisatorID int, input app.UpdateOrganisatorInput) (int64, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, *input.Title)
		argID++
	}
	if input.LogoPath != nil {
		setValues = append(setValues, fmt.Sprintf("logo_path=$%d", argID))
		args = append(args, *input.LogoPath)
		argID++
	}
	if input.SiteUrl != nil {
		setValues = append(setValues, fmt.Sprintf("site_url=$%d", argID))
		args = append(args, *input.SiteUrl)
		argID++
	}
	// Set UpdatedAt Time
	setValues = append(setValues, fmt.Sprintf("updated_at=$%d", argID))
	args = append(args, time.Now().UTC())
	argID++

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s SET %s
			WHERE id=$%d`,
		organisatorsTable, setQuery, argID)

	args = append(args, organisatorID)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	result, err := r.db.Exec(query, args...)
	rowsAffected, errRows := result.RowsAffected()
	if errRows != nil {
		return 0, errRows
	}

	return rowsAffected, err
}
