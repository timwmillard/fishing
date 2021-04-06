
-- name: ListCompetitors :many
SELECT * FROM competitors
ORDER BY competitor_no, lastname, firstname ASC;

-- name: GetCompetitors :one
SELECT * FROM competitors
WHERE id = $1;