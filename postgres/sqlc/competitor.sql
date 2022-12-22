
-- name: ListCompetitor :many
SELECT id, competitor_no, first_name, last_name, email, address1, address2, suburb, state, postcode, mobile
FROM fishing.competitor
ORDER BY competitor_no, last_name, first_name ASC;

-- name: GetCompetitor :one
SELECT id, competitor_no, first_name, last_name, email, address1, address2, suburb, state, postcode, mobile
FROM fishing.competitor
WHERE id = $1;

-- name: CreateCompetitor :one
INSERT INTO fishing.competitor (
    id, competitor_no, first_name, last_name, email, address1, address2, suburb, state, postcode, mobile
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING *;

-- name: UpdateCompetitor :one
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
RETURNING *;

-- name: DeleteCompetitor :execrows
DELETE
FROM fishing.competitor
WHERE id = $1;