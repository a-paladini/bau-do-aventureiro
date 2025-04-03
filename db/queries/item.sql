-- name: CreateItem :one
INSERT INTO items
(name, description, category, price, slot, origin)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetItem :one
SELECT * FROM items
WHERE id = $1;

-- name: ListAllItems :many
SELECT * FROM items
ORDER BY category, price OFFSET 5;

-- name: ListItemsByCategory :many
SELECT * FROM items
WHERE UPPER(category) = $1
ORDER BY price, name;

-- name: UpdateItem :one
UPDATE items
SET name = $2, description = $3, category = $4, price = $5, slot = $6,
origin = $7
WHERE id = $1
RETURNING *;

-- name: DeleteItem :exec
DELETE FROM items
WHERE id = $1;