
-- name: ListCompetitors :many
SELECT * FROM competitors
ORDER BY competitor_no, lastname, firstname ASC;

-- name: GetCompetitor :one
SELECT * FROM competitors
WHERE id = $1;

-- name: CreateCompetitor :one
INSERT INTO competitors (
    id,
    firstname,
    lastname
) VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: UpdateCompetitor :one
UPDATE competitors
SET competitor_no = $2,
    firstname = $3,
    lastname = $4
WHERE id = $1
RETURNING *;

-- name: DeleteCompetitor :execrows
DELETE FROM competitors
WHERE id = $1;