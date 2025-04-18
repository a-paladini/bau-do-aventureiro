-- name: CreateArmour :one
INSERT INTO armours
(name, description, price, slot,
origin, ca_bonus, penality, category)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetArmour :one
SELECT * FROM armours
WHERE id = $1;

-- name: ListAllArmours :many
SELECT * FROM armours
ORDER BY category, ca_bonus, price LIMIT $1 OFFSET $2;

-- name: ListArmoursByCategory :many
SELECT * FROM armours
WHERE LOWER(category) = $1
ORDER BY ca_bonus, price LIMIT $2 OFFSET $3;

-- name: UpdateArmour :one
UPDATE armours
SET name = $2, description = $3, price = $4, slot = $5,
origin = $6, ca_bonus = $7, penality = $8, category = $9
WHERE id = $1
RETURNING *;

-- name: DeleteArmour :exec
DELETE FROM armours
WHERE id = $1;