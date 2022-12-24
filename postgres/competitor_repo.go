// Postgres implementation of CompetitorRepo.
package postgres

import (
	"context"
	"database/sql"

	"github.com/timwmillard/fishing"
)

var _ fishing.CompetitorRepo = (*CompetitorRepo)(nil)

// CompetitorRepo is a repository of competitors.
type CompetitorRepo struct {
	DB DBTX
}

const listCompetitors = `-- name: ListCompetitors :many
SELECT id, competitor_no, first_name, last_name, email, address1, address2, suburb, state, postcode, mobile
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
	items = make([]fishing.Competitor, 0)
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
SELECT id, competitor_no, first_name, last_name, email, address1, address2, suburb, state, postcode, mobile
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
		&i.Mobile,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return i, fishing.ErrCompetitorNotFound
		}
		return i, err
	}
	return i, nil
}

const createCompetitor = `-- name: CreateCompetitor :one
INSERT INTO fishing.competitor (
	competitor_no, first_name, last_name, email, address1, address2, suburb, state, postcode, mobile
) VALUES (
	$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING id, competitor_no, first_name, last_name, email, address1, address2, suburb, state, postcode, mobile
`

// Create's a new competitor.
func (r *CompetitorRepo) Create(ctx context.Context, arg fishing.CreateCompetitorParams) (fishing.Competitor, error) {
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
		&i.Mobile,
	)
	return i, err
}

const updateCompetitor = `-- name: UpdateCompetitor :one
UPDATE fishing.competitor
SET competitor_no = COALESCE($2, competitor_no),
    first_name = COALESCE($3, first_name),
    last_name = COALESCE($4, last_name),
    email = COALESCE($5, email),
    address1 = COALESCE($6, address1),
    address2 = COALESCE($7, address2),
    suburb = COALESCE($8, suburb),
    state = COALESCE($9, state),
    postcode = COALESCE($10, postcode),
    mobile = COALESCE($11, mobile)
WHERE id = $1
RETURNING id, competitor_no, first_name, last_name, email, address1, address2, suburb, state, postcode, mobile
`

type updateCompetitorParams struct {
	CompetitorNo sql.NullString
	FirstName    sql.NullString
	LastName     sql.NullString
	Email        sql.NullString
	Address1     sql.NullString
	Address2     sql.NullString
	Suburb       sql.NullString
	State        sql.NullString
	Postcode     sql.NullString
	Mobile       sql.NullString
}

func newUpdateCompetitorParams(param fishing.UpdateCompetitorParams) updateCompetitorParams {
	return updateCompetitorParams{
		CompetitorNo: toNullString(param.CompetitorNo),
		FirstName:    toNullString(param.FirstName),
		LastName:     toNullString(param.LastName),
		Email:        toNullString(param.Email),
		Address1:     toNullString(param.Address1),
		Address2:     toNullString(param.Address2),
		Suburb:       toNullString(param.Suburb),
		State:        toNullString(param.State),
		Postcode:     toNullString(param.Postcode),
		Mobile:       toNullString(param.Mobile),
	}
}

// Update's an existing competitor.  Returns the updated competitor.
func (r *CompetitorRepo) Update(ctx context.Context, id fishing.HashID, arg fishing.CreateCompetitorParams) (fishing.Competitor, error) {
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
		&i.Mobile,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return i, fishing.ErrCompetitorNotFound
		}
		return i, err
	}
	return i, nil
}

// Partially update's an existing competitor.  Returns the updated competitor.
func (r *CompetitorRepo) UpdatePartial(ctx context.Context, id fishing.HashID, params fishing.UpdateCompetitorParams) (fishing.Competitor, error) {
	arg := newUpdateCompetitorParams(params)
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
		&i.Mobile,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return i, fishing.ErrCompetitorNotFound
		}
		return i, err
	}
	return i, nil
}

const deleteCompetitor = `-- name: DeleteCompetitor :execrows
DELETE
FROM fishing.competitor
WHERE id = $1
`

// Delete's a competitor by id.
func (r *CompetitorRepo) Delete(ctx context.Context, id fishing.HashID) error {
	result, err := r.DB.ExecContext(ctx, deleteCompetitor, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows < 1 {
		return fishing.ErrCompetitorNotFound
	}
	return nil
}
