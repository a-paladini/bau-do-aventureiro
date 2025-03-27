package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createTestItem(t *testing.T) []Items {
	sheets, err := ReadExcelSheets("testdata/data.xlsx")
	if err != nil {
		t.Fatalf("Error reading Excel file: %v", err)
	}

	excel, err := ProcessExcelDataItems(sheets)
	if err != nil {
		t.Fatalf("Error processing data: %v", err)
	}

	if len(excel) == 0 {
		t.Fatalf("No register was found in Items sheet")
	}

	var listItems []Items

	for _, i := range excel {
		arg := CreateItemParams{
			Name:        i.Name,
			Description: i.Description,
			Price:       i.Price,
			Slot:        i.Slot,
			Origin:      i.Origin,
		}

		item, err := testQueries.CreateItem(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, item)

		require.Equal(t, arg.Name, item.Name)
		require.Equal(t, arg.Description, item.Description)
		require.Equal(t, arg.Price, item.Price)
		require.Equal(t, arg.Slot, item.Slot)

		listItems = append(listItems, item)
	}

	return listItems
}

func TestCreateItem(t *testing.T) {
	createTestItem(t)
}

func TestGetItem(t *testing.T) {
	items := createTestItem(t)

	for _, item := range items {
		i, err := testQueries.GetItem(context.Background(), item.ID)
		require.NoError(t, err)
		require.NotEmpty(t, i)

		require.Equal(t, item.Name, i.Name)
		require.Equal(t, item.Description, i.Description)
		require.Equal(t, item.Price, i.Price)
		require.Equal(t, item.Slot, i.Slot)
	}
}

func TestListItems(t *testing.T) {
	_ = createTestItem(t)

	items, err := testQueries.ListItems(context.Background())
	require.NoError(t, err)

	for _, item := range items {
		require.NotEmpty(t, item)
	}
}

func TestUpdateItem(t *testing.T) {
	items := createTestItem(t)

	for _, item := range items {
		arg := UpdateItemParams{
			ID:          item.ID,
			Name:        item.Name + "_updated",
			Description: item.Description + "_updated",
			Price:       item.Price + 100,
			Slot:        item.Slot + 1,
		}

		i, err := testQueries.UpdateItem(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, i)

		require.Equal(t, arg.Name, i.Name)
		require.Equal(t, arg.Description, i.Description)
		require.Equal(t, arg.Price, i.Price)
		require.Equal(t, arg.Slot, i.Slot)
	}
}

func TestDeleteItem(t *testing.T) {
	items := createTestItem(t)

	for _, item := range items {
		err := testQueries.DeleteItem(context.Background(), item.ID)
		require.NoError(t, err)

		i, err := testQueries.GetItem(context.Background(), item.ID)
		require.Error(t, err)
		require.Empty(t, i)
	}
}
