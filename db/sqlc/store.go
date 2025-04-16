package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// ExecTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// Operations for Weapons
func (store *Store) CreateWeaponTx(ctx context.Context, arg CreateWeaponParams) (Weapons, error) {
	var weapon Weapons
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		weapon, err = q.CreateWeapon(ctx, arg)
		return err
	})
	return weapon, err
}

func (store *Store) GetWeaponTx(ctx context.Context, id int32) (Weapons, error) {
	var weapon Weapons
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		weapon, err = q.GetWeapon(ctx, id)
		return err
	})
	return weapon, err
}

func (store *Store) ListAllWeaponsTx(ctx context.Context, args ListAllWeaponsParams) ([]Weapons, error) {
	var listWeapons []Weapons

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		listWeapons, err = q.ListAllWeapons(ctx, args)
		return err
	})
	return listWeapons, err
}

func (store *Store) ListWeaponsByCategoryTx(ctx context.Context, arg ListWeaponsByCategoryParams) ([]Weapons, error) {
	var listWeapons []Weapons

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		listWeapons, err = q.ListWeaponsByCategory(ctx, arg)
		return err
	})
	return listWeapons, err
}

func (store *Store) UpdateWeaponTx(ctx context.Context, arg UpdateWeaponParams) (Weapons, error) {
	var weapon Weapons
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		weapon, err = q.UpdateWeapon(ctx, arg)
		return err
	})
	return weapon, err
}

func (store *Store) DeleteWeaponTx(ctx context.Context, id int32) error {
	return store.execTx(ctx, func(q *Queries) error {
		return q.DeleteWeapon(ctx, id)
	})
}

// Operations for Items
func (store *Store) CreateItemTx(ctx context.Context, arg CreateItemParams) (Items, error) {
	var item Items
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		item, err = q.CreateItem(ctx, arg)
		return err
	})
	return item, err
}

func (store *Store) GetItemTx(ctx context.Context, id int32) (Items, error) {
	var item Items
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		item, err = q.GetItem(ctx, id)
		return err
	})
	return item, err
}

func (store *Store) ListAllItemsTx(ctx context.Context, arg ListAllItemsParams) ([]Items, error) {
	var items []Items
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		items, err = q.ListAllItems(ctx, arg)
		return err
	})
	return items, err
}

func (store *Store) ListItemsByCategoryTx(ctx context.Context, arg ListItemsByCategoryParams) ([]Items, error) {
	var items []Items
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		items, err = q.ListItemsByCategory(ctx, arg)
		return err
	})
	return items, err
}

func (store *Store) UpdateItemTx(ctx context.Context, arg UpdateItemParams) (Items, error) {
	var item Items
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		item, err = q.UpdateItem(ctx, arg)
		return err
	})
	return item, err
}

func (store *Store) DeleteItemTx(ctx context.Context, id int32) error {
	return store.execTx(ctx, func(q *Queries) error {
		return q.DeleteItem(ctx, id)
	})
}

// Operations for Armours
func (store *Store) CreateArmourTx(ctx context.Context, arg CreateArmourParams) (Armours, error) {
	var armour Armours
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		armour, err = q.CreateArmour(ctx, arg)
		return err
	})
	return armour, err
}

func (store *Store) GetArmourTx(ctx context.Context, id int32) (Armours, error) {
	var armour Armours
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		armour, err = q.GetArmour(ctx, id)
		return err
	})
	return armour, err
}

func (store *Store) ListAllArmoursTx(ctx context.Context, arg ListAllArmoursParams) ([]Armours, error) {
	var armours []Armours
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		armours, err = q.ListAllArmours(ctx, arg)
		return err
	})

	return armours, err
}

func (store *Store) ListArmoursByCategoryTx(ctx context.Context, arg ListArmoursByCategoryParams) ([]Armours, error) {
	var armours []Armours
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		armours, err = q.ListArmoursByCategory(context.Background(), arg)
		return err
	})
	return armours, err
}

func (store *Store) UpdateArmourTx(ctx context.Context, arg UpdateArmourParams) (Armours, error) {
	var armour Armours
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		armour, err = q.UpdateArmour(ctx, arg)
		return err
	})
	return armour, err
}

func (store *Store) DeleteArmourTx(ctx context.Context, id int32) error {
	return store.execTx(ctx, func(q *Queries) error {
		return q.DeleteArmour(ctx, id)
	})
}
