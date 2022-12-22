
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
SET competitor_no = $2,
    first_name = $3,
    last_name = $4,
    email = $5, 
    address1 = $6,
    address2 = $7,
    suburb = $8,
    state = $9,
    postcode = $10,
    mobile = $11
WHERE id = $1
RETURNING *;

-- name: DeleteCompetitor :execrows
DELETE
FROM fishing.competitor
WHERE id = $1;