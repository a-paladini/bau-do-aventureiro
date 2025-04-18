// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: armour.sql

package db

import (
	"context"
)

const createArmour = `-- name: CreateArmour :one
INSERT INTO armours
(name, description, price, slot,
origin, ca_bonus, penality, category)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, name, description, category, price, slot, origin, ca_bonus, penality
`

type CreateArmourParams struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Slot        float64 `json:"slot"`
	Origin      string  `json:"origin"`
	CaBonus     int32   `json:"ca_bonus"`
	Penality    int32   `json:"penality"`
	Category    string  `json:"category"`
}

func (q *Queries) CreateArmour(ctx context.Context, arg CreateArmourParams) (Armours, error) {
	row := q.db.QueryRowContext(ctx, createArmour,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Slot,
		arg.Origin,
		arg.CaBonus,
		arg.Penality,
		arg.Category,
	)
	var i Armours
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Category,
		&i.Price,
		&i.Slot,
		&i.Origin,
		&i.CaBonus,
		&i.Penality,
	)
	return i, err
}

const deleteArmour = `-- name: DeleteArmour :exec
DELETE FROM armours
WHERE id = $1
`

func (q *Queries) DeleteArmour(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteArmour, id)
	return err
}

const getArmour = `-- name: GetArmour :one
SELECT id, name, description, category, price, slot, origin, ca_bonus, penality FROM armours
WHERE id = $1
`

func (q *Queries) GetArmour(ctx context.Context, id int32) (Armours, error) {
	row := q.db.QueryRowContext(ctx, getArmour, id)
	var i Armours
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Category,
		&i.Price,
		&i.Slot,
		&i.Origin,
		&i.CaBonus,
		&i.Penality,
	)
	return i, err
}

const listAllArmours = `-- name: ListAllArmours :many
SELECT id, name, description, category, price, slot, origin, ca_bonus, penality FROM armours
ORDER BY category, ca_bonus, price LIMIT $1 OFFSET $2
`

type ListAllArmoursParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAllArmours(ctx context.Context, arg ListAllArmoursParams) ([]Armours, error) {
	rows, err := q.db.QueryContext(ctx, listAllArmours, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Armours{}
	for rows.Next() {
		var i Armours
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Category,
			&i.Price,
			&i.Slot,
			&i.Origin,
			&i.CaBonus,
			&i.Penality,
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

const listArmoursByCategory = `-- name: ListArmoursByCategory :many
SELECT id, name, description, category, price, slot, origin, ca_bonus, penality FROM armours
WHERE category = $1
ORDER BY ca_bonus, price LIMIT $2 OFFSET $3
`

type ListArmoursByCategoryParams struct {
	Category string `json:"category"`
	Limit    int32  `json:"limit"`
	Offset   int32  `json:"offset"`
}

func (q *Queries) ListArmoursByCategory(ctx context.Context, arg ListArmoursByCategoryParams) ([]Armours, error) {
	rows, err := q.db.QueryContext(ctx, listArmoursByCategory, arg.Category, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Armours{}
	for rows.Next() {
		var i Armours
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Category,
			&i.Price,
			&i.Slot,
			&i.Origin,
			&i.CaBonus,
			&i.Penality,
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

const updateArmour = `-- name: UpdateArmour :one
UPDATE armours
SET name = $2, description = $3, price = $4, slot = $5,
origin = $6, ca_bonus = $7, penality = $8, category = $9
WHERE id = $1
RETURNING id, name, description, category, price, slot, origin, ca_bonus, penality
`

type UpdateArmourParams struct {
	ID          int32   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Slot        float64 `json:"slot"`
	Origin      string  `json:"origin"`
	CaBonus     int32   `json:"ca_bonus"`
	Penality    int32   `json:"penality"`
	Category    string  `json:"category"`
}

func (q *Queries) UpdateArmour(ctx context.Context, arg UpdateArmourParams) (Armours, error) {
	row := q.db.QueryRowContext(ctx, updateArmour,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Slot,
		arg.Origin,
		arg.CaBonus,
		arg.Penality,
		arg.Category,
	)
	var i Armours
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Category,
		&i.Price,
		&i.Slot,
		&i.Origin,
		&i.CaBonus,
		&i.Penality,
	)
	return i, err
}
