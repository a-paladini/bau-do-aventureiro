package db

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func createTestWeapon(t *testing.T) []Weapons {
	sheets, err := ReadExcelSheets("testdata/data.xlsx")
	if err != nil {
		log.Fatalf("Error reading Excel file: %v", err)
	}

	excel, err := ProcessExcelDataWeapons(sheets)
	if err != nil {
		log.Fatalf("Error processing data: %v", err)
	}

	if len(excel) == 0 {
		log.Fatalf("No register was found in Weapons sheet")
	}

	var listWeapons []Weapons

	for _, w := range excel {
		arg := CreateWeaponParams{
			Name:        w.Name,
			Description: w.Description,
			Price:       w.Price,
			Slot:        w.Slot,
			Origin:      w.Origin,
			Damage:      w.Damage,
			Critical:    w.Critical,
			Range:       w.Range,
			TypeDamage:  w.TypeDamage,
			Property:    w.Property,
			Proficiency: w.Proficiency,
			Special:     w.Special,
		}

		weapon, err := testQueries.CreateWeapon(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, weapon)

		require.Equal(t, arg.Name, weapon.Name)
		require.Equal(t, arg.Description, weapon.Description)
		require.Equal(t, arg.Price, weapon.Price)
		require.Equal(t, arg.Slot, weapon.Slot)
		require.Equal(t, arg.Origin, weapon.Origin)
		require.Equal(t, arg.Damage, weapon.Damage)
		require.Equal(t, arg.Critical, weapon.Critical)
		require.Equal(t, arg.Range, weapon.Range)
		require.Equal(t, arg.TypeDamage, weapon.TypeDamage)
		require.Equal(t, arg.Property, weapon.Property)
		require.Equal(t, arg.Proficiency, weapon.Proficiency)
		require.Equal(t, arg.Special, weapon.Special)

		require.NotZero(t, weapon.ID)
		require.NotZero(t, weapon.Price)
		require.NotZero(t, weapon.Slot)

		require.NotEmpty(t, weapon.Name)
		require.NotEmpty(t, weapon.Description)
		require.NotEmpty(t, weapon.Origin)
		require.NotEmpty(t, weapon.Damage)
		require.NotEmpty(t, weapon.Critical)
		require.NotEmpty(t, weapon.Range)
		require.NotEmpty(t, weapon.TypeDamage)
		require.NotEmpty(t, weapon.Property)
		require.NotEmpty(t, weapon.Proficiency)

		listWeapons = append(listWeapons, weapon)
	}

	return listWeapons
}

func TestCreateWeapon(t *testing.T) {
	createTestWeapon(t)
}

func TestGetWeapon(t *testing.T) {
	weapons := createTestWeapon(t)

	for _, w := range weapons {
		weapon, err := testQueries.GetWeapon(context.Background(), w.ID)
		require.NoError(t, err)
		require.NotEmpty(t, weapon)

		require.Equal(t, w.ID, weapon.ID)
		require.Equal(t, w.Name, weapon.Name)
		require.Equal(t, w.Description, weapon.Description)
		require.Equal(t, w.Price, weapon.Price)
		require.Equal(t, w.Slot, weapon.Slot)
		require.Equal(t, w.Origin, weapon.Origin)
		require.Equal(t, w.Damage, weapon.Damage)
		require.Equal(t, w.Critical, weapon.Critical)
		require.Equal(t, w.Range, weapon.Range)
		require.Equal(t, w.TypeDamage, weapon.TypeDamage)
		require.Equal(t, w.Property, weapon.Property)
		require.Equal(t, w.Proficiency, weapon.Proficiency)
		require.Equal(t, w.Special, weapon.Special)
	}
}

func TestDeleteWeapon(t *testing.T) {
	weapons := createTestWeapon(t)

	for _, w := range weapons {
		err := testQueries.DeleteWeapon(context.Background(), w.ID)
		require.NoError(t, err)

		weapon, err := testQueries.GetWeapon(context.Background(), w.ID)
		require.Error(t, err)
		require.Empty(t, weapon)
	}
}

func TestListWeapons(t *testing.T) {
	_ = createTestWeapon(t)

	listWeapons, err := testQueries.ListWeapons(context.Background())
	require.NoError(t, err)

	for _, w := range listWeapons {
		require.NotEmpty(t, w)
	}
}

func TestUpdateWeapon(t *testing.T) {
	weapons := createTestWeapon(t)

	for _, w := range weapons {
		arg := UpdateWeaponParams{
			ID:          w.ID,
			Name:        w.Name + "_updated",
			Description: w.Description + "_updated",
			Price:       w.Price + 100,
			Slot:        w.Slot + 1,
			Origin:      w.Origin + "_updated",
			Damage:      w.Damage + "_updated",
			Critical:    w.Critical + "_updated",
			Range:       w.Range + "_updated",
			TypeDamage:  w.TypeDamage + "_updated",
			Property:    w.Property + "_updated",
			Proficiency: w.Proficiency + "_updated",
			Special:     sql.NullString{String: w.Special.String + "_updated", Valid: true},
		}

		weapon, err := testQueries.UpdateWeapon(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, weapon)

		require.Equal(t, arg.ID, weapon.ID)
		require.Equal(t, arg.Name, weapon.Name)
		require.Equal(t, arg.Description, weapon.Description)
		require.Equal(t, arg.Price, weapon.Price)
		require.Equal(t, arg.Slot, weapon.Slot)
		require.Equal(t, arg.Origin, weapon.Origin)
		require.Equal(t, arg.Damage, weapon.Damage)
		require.Equal(t, arg.Critical, weapon.Critical)
		require.Equal(t, arg.Range, weapon.Range)
		require.Equal(t, arg.TypeDamage, weapon.TypeDamage)
		require.Equal(t, arg.Property, weapon.Property)
		require.Equal(t, arg.Proficiency, weapon.Proficiency)
		require.Equal(t, arg.Special, weapon.Special)
	}
}
