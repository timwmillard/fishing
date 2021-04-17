// Code generated by sqlc. DO NOT EDIT.
// source: competitors.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createCompetitor = `-- name: CreateCompetitor :one
INSERT INTO competitors (
    id, competitor_no, firstname, lastname, email, address1, address2, suburb, state, postcode, phone, mobile
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
)
RETURNING id, competitor_no, firstname, lastname, email, address1, address2, suburb, state, postcode, phone, mobile
`

type CreateCompetitorParams struct {
	ID           uuid.UUID
	CompetitorNo string
	Firstname    string
	Lastname     string
	Email        string
	Address1     string
	Address2     string
	Suburb       string
	State        string
	Postcode     string
	Phone        string
	Mobile       string
}

func (q *Queries) CreateCompetitor(ctx context.Context, arg CreateCompetitorParams) (Competitor, error) {
	row := q.db.QueryRowContext(ctx, createCompetitor,
		arg.ID,
		arg.CompetitorNo,
		arg.Firstname,
		arg.Lastname,
		arg.Email,
		arg.Address1,
		arg.Address2,
		arg.Suburb,
		arg.State,
		arg.Postcode,
		arg.Phone,
		arg.Mobile,
	)
	var i Competitor
	err := row.Scan(
		&i.ID,
		&i.CompetitorNo,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.Address1,
		&i.Address2,
		&i.Suburb,
		&i.State,
		&i.Postcode,
		&i.Phone,
		&i.Mobile,
	)
	return i, err
}

const deleteCompetitor = `-- name: DeleteCompetitor :execrows
DELETE FROM competitors
WHERE id = $1
`

func (q *Queries) DeleteCompetitor(ctx context.Context, id uuid.UUID) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteCompetitor, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const getCompetitor = `-- name: GetCompetitor :one
SELECT id, competitor_no, firstname, lastname, email, address1, address2, suburb, state, postcode, phone, mobile FROM competitors
WHERE id = $1
`

func (q *Queries) GetCompetitor(ctx context.Context, id uuid.UUID) (Competitor, error) {
	row := q.db.QueryRowContext(ctx, getCompetitor, id)
	var i Competitor
	err := row.Scan(
		&i.ID,
		&i.CompetitorNo,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.Address1,
		&i.Address2,
		&i.Suburb,
		&i.State,
		&i.Postcode,
		&i.Phone,
		&i.Mobile,
	)
	return i, err
}

const listCompetitors = `-- name: ListCompetitors :many
SELECT id, competitor_no, firstname, lastname, email, address1, address2, suburb, state, postcode, phone, mobile FROM competitors
ORDER BY competitor_no, lastname, firstname ASC
`

func (q *Queries) ListCompetitors(ctx context.Context) ([]Competitor, error) {
	rows, err := q.db.QueryContext(ctx, listCompetitors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Competitor
	for rows.Next() {
		var i Competitor
		if err := rows.Scan(
			&i.ID,
			&i.CompetitorNo,
			&i.Firstname,
			&i.Lastname,
			&i.Email,
			&i.Address1,
			&i.Address2,
			&i.Suburb,
			&i.State,
			&i.Postcode,
			&i.Phone,
			&i.Mobile,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCompetitor = `-- name: UpdateCompetitor :one
UPDATE competitors
SET competitor_no = $2,
    firstname = $3,
    lastname = $4,
    email = $5, 
    address1 = $6,
    address2 = $7,
    suburb = $8,
    state = $9,
    postcode = $10,
    phone = $11,
    mobile = $12
WHERE id = $1
RETURNING id, competitor_no, firstname, lastname, email, address1, address2, suburb, state, postcode, phone, mobile
`

type UpdateCompetitorParams struct {
	ID           uuid.UUID
	CompetitorNo string
	Firstname    string
	Lastname     string
	Email        string
	Address1     string
	Address2     string
	Suburb       string
	State        string
	Postcode     string
	Phone        string
	Mobile       string
}

func (q *Queries) UpdateCompetitor(ctx context.Context, arg UpdateCompetitorParams) (Competitor, error) {
	row := q.db.QueryRowContext(ctx, updateCompetitor,
		arg.ID,
		arg.CompetitorNo,
		arg.Firstname,
		arg.Lastname,
		arg.Email,
		arg.Address1,
		arg.Address2,
		arg.Suburb,
		arg.State,
		arg.Postcode,
		arg.Phone,
		arg.Mobile,
	)
	var i Competitor
	err := row.Scan(
		&i.ID,
		&i.CompetitorNo,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.Address1,
		&i.Address2,
		&i.Suburb,
		&i.State,
		&i.Postcode,
		&i.Phone,
		&i.Mobile,
	)
	return i, err
}
