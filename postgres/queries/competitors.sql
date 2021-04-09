
-- name: ListCompetitors :many
SELECT id, competitor_no, firstname, lastname, email, address1, address2, suburb, state, postcode, phone, mobile FROM competitors
ORDER BY competitor_no, lastname, firstname ASC;

-- name: GetCompetitor :one
SELECT id, competitor_no, firstname, lastname, email, address1, address2, suburb, state, postcode, phone, mobile FROM competitors
WHERE id = $1;

-- name: CreateCompetitor :one
INSERT INTO competitors (
    id, competitor_no, firstname, lastname, email, address1, address2, suburb, state, postcode, phone, mobile
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
)
RETURNING *;

-- name: UpdateCompetitor :one
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
RETURNING *;

-- name: DeleteCompetitor :execrows
DELETE FROM competitors
WHERE id = $1;