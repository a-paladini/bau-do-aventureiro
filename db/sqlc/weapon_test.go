package db

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func createTestWeapon(t *testing.T) Weapons {
	arg := CreateWeaponParams{
		Name:        "Adaga",
		Description: "Descrição longa de exemplo da Adaga",
		Price:       2,
		Slot:        1,
		Origin:      "Livro T20 - Base",
		Damage:      "1d4",
		Critical:    "19",
		Range:       "Curto",
		Category:    "Perfuração",
		Property:    "Leve",
		Proficiency: "Simples",
		Special:     sql.NullString{String: "Discreta", Valid: true},
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
	require.Equal(t, arg.Category, weapon.Category)
	require.Equal(t, arg.Property, weapon.Property)
	require.Equal(t, arg.Proficiency, weapon.Proficiency)
	require.Equal(t, arg.Special, weapon.Special)

	require.NotZero(t, weapon.ID)

	require.NotEmpty(t, weapon.Name)
	require.NotEmpty(t, weapon.Description)
	require.NotEmpty(t, weapon.Origin)
	require.NotEmpty(t, weapon.Damage)
	require.NotEmpty(t, weapon.Critical)
	require.NotEmpty(t, weapon.Range)
	require.NotEmpty(t, weapon.Category)
	require.NotEmpty(t, weapon.Property)
	require.NotEmpty(t, weapon.Proficiency)

	return weapon
}

func createTestWeapons(t *testing.T) []Weapons {
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
			Category:    w.Category,
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
		require.Equal(t, arg.Category, weapon.Category)
		require.Equal(t, arg.Property, weapon.Property)
		require.Equal(t, arg.Proficiency, weapon.Proficiency)
		require.Equal(t, arg.Special, weapon.Special)

		require.NotZero(t, weapon.ID)

		require.NotEmpty(t, weapon.Name)
		require.NotEmpty(t, weapon.Description)
		require.NotEmpty(t, weapon.Origin)
		require.NotEmpty(t, weapon.Damage)
		require.NotEmpty(t, weapon.Critical)
		require.NotEmpty(t, weapon.Range)
		require.NotEmpty(t, weapon.Category)
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
	weapons := createTestWeapons(t)

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
		require.Equal(t, w.Category, weapon.Category)
		require.Equal(t, w.Property, weapon.Property)
		require.Equal(t, w.Proficiency, weapon.Proficiency)
		require.Equal(t, w.Special, weapon.Special)
	}
}

func TestDeleteWeapon(t *testing.T) {
	weapon := createTestWeapon(t)

	err := testQueries.DeleteWeapon(context.Background(), weapon.ID)
	require.NoError(t, err)

	weapon2, err := testQueries.GetWeapon(context.Background(), weapon.ID)
	require.Error(t, err)
	require.Empty(t, weapon2)
}

func TestListWeapons(t *testing.T) {
	_ = createTestWeapon(t)

	args := ListAllWeaponsParams{
		Limit:  10,
		Offset: 5,
	}

	listWeapons, err := testQueries.ListAllWeapons(context.Background(), args)
	require.NoError(t, err)

	for _, w := range listWeapons {
		require.NotEmpty(t, w)
	}
}

func TestUpdateWeapon(t *testing.T) {
	weapon := createTestWeapon(t)

	arg := UpdateWeaponParams{
		ID:          weapon.ID,
		Name:        weapon.Name + "_updated",
		Description: weapon.Description + "_updated",
		Price:       weapon.Price + 100,
		Slot:        weapon.Slot + 1,
		Origin:      weapon.Origin + "_updated",
		Damage:      weapon.Damage + "_updated",
		Critical:    weapon.Critical + "_updated",
		Range:       weapon.Range + "_updated",
		Category:    weapon.Category + "_updated",
		Property:    weapon.Property + "_updated",
		Proficiency: weapon.Proficiency + "_updated",
		Special:     sql.NullString{String: weapon.Special.String + "_updated", Valid: true},
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
	require.Equal(t, arg.Category, weapon.Category)
	require.Equal(t, arg.Property, weapon.Property)
	require.Equal(t, arg.Proficiency, weapon.Proficiency)
	require.Equal(t, arg.Special, weapon.Special)
}
