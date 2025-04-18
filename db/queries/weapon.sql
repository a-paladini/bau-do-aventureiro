-- name: CreateWeapon :one
INSERT INTO weapons 
(name, description, price, slot, 
origin, damage, critical, range, category, 
property, proficiency, special)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) 
RETURNING *;

-- name: GetWeapon :one
SELECT * FROM weapons
WHERE id = $1;

-- name: ListAllWeapons :many
SELECT * FROM weapons
ORDER BY name LIMIT $1 OFFSET $2;

-- name: ListWeaponsByCategory :many
SELECT * FROM weapons
WHERE LOWER(category) = $1
ORDER BY price LIMIT $2 OFFSET $3;

-- name: UpdateWeapon :one
UPDATE weapons
SET name = $2, description = $3, price = $4, slot = $5,
origin = $6, damage = $7, critical = $8, range = $9, category = $10,
property = $11, proficiency = $12, special = $13
WHERE id = $1
RETURNING *;

-- name: DeleteWeapon :exec
DELETE FROM weapons
WHERE id = $1;