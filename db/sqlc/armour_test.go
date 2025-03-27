package db

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func createTestArmour(t *testing.T) []Armours {
	sheets, err := ReadExcelSheets("testdata/data.xlsx")
	if err != nil {
		log.Fatalf("Error reading Excel file: %v", err)
	}

	excel, err := ProcessExcelDataArmours(sheets)
	if err != nil {
		log.Fatalf("Error processing data: %v", err)
	}

	if len(excel) == 0 {
		log.Fatalf("No register was found in Armours sheet")
	}

	var listArmours []Armours

	for _, a := range excel {
		arg := CreateArmourParams{
			Name:        a.Name,
			Description: a.Description,
			Category:    a.Category,
			Price:       a.Price,
			Slot:        a.Slot,
			Origin:      a.Origin,
			CaBonus:     a.CaBonus,
			Penality:    a.Penality,
		}

		armour, err := testQueries.CreateArmour(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, armour)

		require.Equal(t, arg.Name, armour.Name)
		require.Equal(t, arg.Description, armour.Description)
		require.Equal(t, arg.Category, armour.Category)
		require.Equal(t, arg.Price, armour.Price)
		require.Equal(t, arg.Slot, armour.Slot)
		require.Equal(t, arg.Origin, armour.Origin)
		require.Equal(t, arg.CaBonus, armour.CaBonus)
		require.Equal(t, arg.Penality, armour.Penality)

		listArmours = append(listArmours, armour)
	}

	return listArmours
}

func TestCreateArmour(t *testing.T) {
	createTestArmour(t)
}

func TestGetArmour(t *testing.T) {
	listArmours := createTestArmour(t)

	for _, a := range listArmours {
		armour, err := testQueries.GetArmour(context.Background(), a.ID)
		require.NoError(t, err)
		require.NotEmpty(t, armour)

		require.Equal(t, a.ID, armour.ID)
		require.Equal(t, a.Name, armour.Name)
		require.Equal(t, a.Description, armour.Description)
		require.Equal(t, a.Category, armour.Category)
		require.Equal(t, a.Price, armour.Price)
		require.Equal(t, a.Slot, armour.Slot)
		require.Equal(t, a.Origin, armour.Origin)
		require.Equal(t, a.CaBonus, armour.CaBonus)
		require.Equal(t, a.Penality, armour.Penality)
	}
}

func TestListArmours(t *testing.T) {
	_ = createTestArmour(t)

	armours, err := testQueries.ListArmours(context.Background())
	require.NoError(t, err)

	for _, a := range armours {
		require.NotEmpty(t, a)
	}
}

func TestUpdateArmour(t *testing.T) {
	listArmours := createTestArmour(t)

	for _, a := range listArmours {
		arg := UpdateArmourParams{
			ID:          a.ID,
			Name:        a.Name + "_updated",
			Description: a.Description + "_updated",
			Category:    a.Category + "_updated",
			Price:       a.Price + 100,
			Slot:        a.Slot + 1,
			Origin:      a.Origin + "_updated",
			CaBonus:     a.CaBonus + 1,
			Penality:    a.Penality + 1,
		}

		armour, err := testQueries.UpdateArmour(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, armour)

		require.Equal(t, arg.ID, armour.ID)
		require.Equal(t, arg.Name, armour.Name)
		require.Equal(t, arg.Description, armour.Description)
		require.Equal(t, arg.Category, armour.Category)
		require.Equal(t, arg.Price, armour.Price)
		require.Equal(t, arg.Slot, armour.Slot)
		require.Equal(t, arg.Origin, armour.Origin)
		require.Equal(t, arg.CaBonus, armour.CaBonus)
		require.Equal(t, arg.Penality, armour.Penality)
	}
}

func TestDeleteArmour(t *testing.T) {
	listArmours := createTestArmour(t)

	for _, a := range listArmours {
		err := testQueries.DeleteArmour(context.Background(), a.ID)
		require.NoError(t, err)

		armour, err := testQueries.GetArmour(context.Background(), a.ID)
		require.Error(t, err)
		require.Empty(t, armour)
	}
}
