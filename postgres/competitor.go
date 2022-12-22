// Postgres implementation of CompetitorRepo.
package postgres

import (
	"context"

	"github.com/timwmillard/fishing"
)

var _ fishing.CompetitorRepo = (*CompetitorRepo)(nil)

// CompetitorRepo is a repository of competitors.
type CompetitorRepo struct {
	DB DBTX
}

const listCompetitors = `-- name: ListCompetitors :many
SELECT id, competitor_no, first_name, last_name, email, address1, address2, suburb, state, postcode, phone, mobile
FROM fishing.competitor
ORDER BY competitor_no, last_name, first_name ASC
`

// List returns a list of all competitors.
func (r *CompetitorRepo) List(ctx context.Context) ([]fishing.Competitor, error) {
	rows, err := r.DB.QueryContext(ctx, listCompetitors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []fishing.Competitor
	for rows.Next() {
		var i fishing.Competitor
		if err := rows.Scan(
			&i.ID,
			&i.CompetitorNo,
			&i.FirstName,
			&i.LastName,
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

const getCompetitor = `-- name: GetCompetitor :one
SELECT id, competitor_no, first_name, last_name, email, address1, address2, suburb, state, postcode, phone, mobile
FROM fishing.competitor
WHERE id = $1
`

// Get's a single competitor by id.
func (r *CompetitorRepo) Get(ctx context.Context, id fishing.HashID) (fishing.Competitor, error) {
	row := r.DB.QueryRowContext(ctx, getCompetitor, id)
	var i fishing.Competitor
	err := row.Scan(
		&i.ID,
		&i.CompetitorNo,
		&i.FirstName,
		&i.LastName,
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

const createCompetitor = `-- name: CreateCompetitor :one
INSERT INTO fishing.competitor (
	competitor_no, first_name, last_name, email, address1, address2, suburb, state, postcode, phone, mobile
) VALUES (
	$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING id, competitor_no, first_name, last_name, email, address1, address2, suburb, state, postcode, phone, mobile
`

// Create's a new competitor.
func (r *CompetitorRepo) Create(ctx context.Context, arg fishing.CompetitorParams) (fishing.Competitor, error) {
	row := r.DB.QueryRowContext(ctx, createCompetitor,
		arg.CompetitorNo,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Address1,
		arg.Address2,
		arg.Suburb,
		arg.State,
		arg.Postcode,
		arg.Phone,
		arg.Mobile,
	)
	var i fishing.Competitor
	err := row.Scan(
		&i.ID,
		&i.CompetitorNo,
		&i.FirstName,
		&i.LastName,
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

const updateCompetitor = `-- name: UpdateCompetitor :one
UPDATE fishing.competitor
SET competitor_no = $2,
    first_name = $3,
    last_name = $4,
    email = $5, 
    address1 = $6,
    address2 = $7,
    suburb = $8,
    state = $9,
    postcode = $10,
    phone = $11,
    mobile = $12
WHERE id = $1
RETURNING id, competitor_no, first_name, last_name, email, address1, address2, suburb, state, postcode, phone, mobile
`

// Update's an existing competitor.  Returns the updated competitor.
func (r *CompetitorRepo) Update(ctx context.Context, id fishing.HashID, arg fishing.CompetitorParams) (fishing.Competitor, error) {
	row := r.DB.QueryRowContext(ctx, updateCompetitor,
		id,
		arg.CompetitorNo,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Address1,
		arg.Address2,
		arg.Suburb,
		arg.State,
		arg.Postcode,
		arg.Phone,
		arg.Mobile,
	)
	var i fishing.Competitor
	err := row.Scan(
		&i.ID,
		&i.CompetitorNo,
		&i.FirstName,
		&i.LastName,
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
DELETE
FROM fishing.competitor
WHERE id = $1
`

// Delete's a competitor by id.
func (r *CompetitorRepo) Delete(ctx context.Context, id fishing.HashID) error {
	_, err := r.DB.ExecContext(ctx, deleteCompetitor, id)
	if err != nil {
		return err
	}
	return nil
}
