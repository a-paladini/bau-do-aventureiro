-- name: CreateWeapon :one
INSERT INTO weapons 
(name, description, category, price, slot, 
origin, damage, critical, range, type_damage, 
property, proficiency, special)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) 
RETURNING *;

-- name: GetWeapon :one
SELECT * FROM weapons
WHERE id = $1;

-- name: ListWeapons :many
SELECT * FROM weapons
ORDER BY id OFFSET 5;

-- name: UpdateWeapon :one
UPDATE weapons
SET name = $2, description = $3, category = $4, price = $5, slot = $6,
origin = $7, damage = $8, critical = $9, range = $10, type_damage = $11,
property = $12, proficiency = $13, special = $14
WHERE id = $1
RETURNING *;

-- name: DeleteWeapon :exec
DELETE FROM weapons
WHERE id = $1;