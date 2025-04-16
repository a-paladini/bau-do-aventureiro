package db

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomWeapon() CreateWeaponParams {
	return CreateWeaponParams{
		Name:        "Adaga",
		Description: "Descrição longa de exemplo da Adaga",
		Price:       2,
		Slot:        1,
		Origin:      "Livro T20 - Base",
		Damage:      "1d4",
		Critical:    "19",
		Range:       "Curto",
		TypeDamage:  "Perfuração",
		Property:    "Leve",
		Proficiency: "Simples",
		Special:     sql.NullString{String: "Discreta", Valid: true},
	}
}

func createManyWeapons(t *testing.T) []Weapons {
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

	store := NewStore(testDB)
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

		weapon, err := store.CreateWeaponTx(context.Background(), arg)
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

func TestCreateWeaponTx(t *testing.T) {
	store := NewStore(testDB)

	arg := createRandomWeapon()

	weapon, err := store.CreateWeaponTx(context.Background(), arg)
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
}

func TestGetWeaponTx(t *testing.T) {
	store := NewStore(testDB)

	arg := createRandomWeapon()
	weapon, err := store.CreateWeaponTx(context.Background(), arg)
	require.NoError(t, err)

	fetchedWeapon, err := store.GetWeaponTx(context.Background(), weapon.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedWeapon)

	require.Equal(t, weapon.ID, fetchedWeapon.ID)
	require.Equal(t, weapon.Name, fetchedWeapon.Name)
	require.Equal(t, weapon.Description, fetchedWeapon.Description)
	require.Equal(t, weapon.Price, fetchedWeapon.Price)
	require.Equal(t, weapon.Slot, fetchedWeapon.Slot)
	require.Equal(t, weapon.Origin, fetchedWeapon.Origin)
	require.Equal(t, weapon.Damage, fetchedWeapon.Damage)
	require.Equal(t, weapon.Critical, fetchedWeapon.Critical)
	require.Equal(t, weapon.Range, fetchedWeapon.Range)
	require.Equal(t, weapon.TypeDamage, fetchedWeapon.TypeDamage)
	require.Equal(t, weapon.Property, fetchedWeapon.Property)
	require.Equal(t, weapon.Proficiency, fetchedWeapon.Proficiency)
	require.Equal(t, weapon.Special, fetchedWeapon.Special)
}

func TestListAllWeaponTx(t *testing.T) {
	store := NewStore(testDB)

	args := ListAllWeaponsParams{
		Limit:  10,
		Offset: 5,
	}

	_ = createManyWeapons(t)
	listWeapons, err := store.ListAllWeapons(context.Background(), args)
	require.NoError(t, err)

	for _, weapon := range listWeapons {
		require.NotEmpty(t, weapon)
	}
}

func TestListWeaponsByCategory(t *testing.T) {
	store := NewStore(testDB)

	aux := createManyWeapons(t)

	args := ListWeaponsByCategoryParams{
		Limit:      10,
		Offset:     5,
		TypeDamage: aux[0].TypeDamage,
	}

	listWeapons, err := store.ListWeaponsByCategoryTx(context.Background(), args)
	require.NoError(t, err)

	for _, weapon := range listWeapons {
		require.NotEmpty(t, weapon)
		require.Equal(t, aux[0].TypeDamage, weapon.TypeDamage)
	}
}

func TestUpdateWeaponTx(t *testing.T) {
	store := NewStore(testDB)

	arg := createRandomWeapon()
	weapon, err := store.CreateWeaponTx(context.Background(), arg)
	require.NoError(t, err)

	updateArg := UpdateWeaponParams{
		ID:          weapon.ID,
		Name:        weapon.Name + "_updated",
		Description: weapon.Description + "_updated",
		Price:       weapon.Price + 50,
		Slot:        weapon.Slot + 1,
		Origin:      weapon.Origin + "_updated",
		Damage:      weapon.Damage + "_updated",
		Critical:    weapon.Critical + "_updated",
		Range:       weapon.Range + "_updated",
		TypeDamage:  weapon.TypeDamage + "_updated",
		Property:    weapon.Property + "_updated",
		Proficiency: weapon.Proficiency + "_updated",
		Special:     sql.NullString{String: weapon.Special.String + "_updated", Valid: true},
	}

	updatedWeapon, err := store.UpdateWeaponTx(context.Background(), updateArg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedWeapon)

	require.Equal(t, updateArg.ID, updatedWeapon.ID)
	require.Equal(t, updateArg.Name, updatedWeapon.Name)
	require.Equal(t, updateArg.Description, updatedWeapon.Description)
	require.Equal(t, updateArg.Price, updatedWeapon.Price)
	require.Equal(t, updateArg.Slot, updatedWeapon.Slot)
	require.Equal(t, updateArg.Origin, updatedWeapon.Origin)
	require.Equal(t, updateArg.Damage, updatedWeapon.Damage)
	require.Equal(t, updateArg.Critical, updatedWeapon.Critical)
	require.Equal(t, updateArg.Range, updatedWeapon.Range)
	require.Equal(t, updateArg.TypeDamage, updatedWeapon.TypeDamage)
	require.Equal(t, updateArg.Property, updatedWeapon.Property)
	require.Equal(t, updateArg.Proficiency, updatedWeapon.Proficiency)
	require.Equal(t, updateArg.Special, updatedWeapon.Special)
}

func TestDeleteWeaponTx(t *testing.T) {
	store := NewStore(testDB)

	arg := createRandomWeapon()
	weapon, err := store.CreateWeaponTx(context.Background(), arg)
	require.NoError(t, err)

	err = store.DeleteWeaponTx(context.Background(), weapon.ID)
	require.NoError(t, err)

	deletedWeapon, err := store.GetWeaponTx(context.Background(), weapon.ID)
	require.Error(t, err)
	require.Empty(t, deletedWeapon)
}

func createRandomArmour() CreateArmourParams {
	return CreateArmourParams{
		Name:        "Armadura de Couro",
		Description: "Descrição de teste exemplo da Armadura de Couro",
		Price:       20,
		Slot:        2,
		Origin:      "Livro T20 - Base",
		CaBonus:     2,
		Penality:    0,
		Category:    "Leve",
	}
}

func createManyArmours(t *testing.T) []Armours {
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

	store := NewStore(testDB)
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

		armour, err := store.CreateArmour(context.Background(), arg)
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

		require.NotZero(t, armour.ID)

		listArmours = append(listArmours, armour)
	}

	return listArmours
}

func TestCreateArmourTx(t *testing.T) {
	store := NewStore(testDB)

	arg := createRandomArmour()

	armour, err := store.CreateArmourTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, armour)

	require.Equal(t, arg.Name, armour.Name)
	require.Equal(t, arg.Description, armour.Description)
	require.Equal(t, arg.Price, armour.Price)
	require.Equal(t, arg.Slot, armour.Slot)
	require.Equal(t, arg.Origin, armour.Origin)
	require.Equal(t, arg.CaBonus, armour.CaBonus)
	require.Equal(t, arg.Penality, armour.Penality)
	require.Equal(t, arg.Category, armour.Category)

	require.NotZero(t, armour.ID)
}

func TestGetArmourTx(t *testing.T) {
	store := NewStore(testDB)

	arg := createRandomArmour()
	armour, err := store.CreateArmourTx(context.Background(), arg)
	require.NoError(t, err)

	fetchedArmour, err := store.GetArmourTx(context.Background(), armour.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedArmour)

	require.Equal(t, armour.ID, fetchedArmour.ID)
	require.Equal(t, armour.Name, fetchedArmour.Name)
	require.Equal(t, armour.Description, fetchedArmour.Description)
	require.Equal(t, armour.Price, fetchedArmour.Price)
	require.Equal(t, armour.Slot, fetchedArmour.Slot)
	require.Equal(t, armour.Origin, fetchedArmour.Origin)
	require.Equal(t, armour.CaBonus, fetchedArmour.CaBonus)
	require.Equal(t, armour.Penality, fetchedArmour.Penality)
	require.Equal(t, armour.Category, fetchedArmour.Category)
}

func TestListAllArmourTx(t *testing.T) {
	store := NewStore(testDB)

	_ = createManyWeapons(t)

	args := ListAllWeaponsParams{
		Limit:  10,
		Offset: 5,
	}

	listArmours, err := store.ListAllWeapons(context.Background(), args)
	require.NoError(t, err)

	for _, armour := range listArmours {
		require.NotEmpty(t, armour)
	}
}

func TestLisArmourByCategoryTx(t *testing.T) {
	store := NewStore(testDB)

	aux := createManyArmours(t)

	arg := ListArmoursByCategoryParams{
		Category: aux[0].Category,
		Limit:    10,
		Offset:   5,
	}

	listArmours, err := store.ListArmoursByCategoryTx(context.Background(), arg)
	require.NoError(t, err)

	for _, armour := range listArmours {
		require.NotEmpty(t, armour)
		require.Equal(t, aux[0].Category, armour.Category)
	}
}

func TestUpdateArmourTx(t *testing.T) {
	store := NewStore(testDB)

	arg := createRandomArmour()
	armour, err := store.CreateArmourTx(context.Background(), arg)
	require.NoError(t, err)

	updateArg := UpdateArmourParams{
		ID:          armour.ID,
		Name:        armour.Name + "_updated",
		Description: armour.Description + "_updated",
		Price:       armour.Price + 50,
		Slot:        armour.Slot + 1,
		Origin:      armour.Origin + "_updated",
		CaBonus:     armour.CaBonus + 1,
		Penality:    armour.Penality + 1,
		Category:    armour.Category + "_update",
	}

	updatedArmour, err := store.UpdateArmourTx(context.Background(), updateArg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedArmour)

	require.Equal(t, updateArg.ID, updatedArmour.ID)
	require.Equal(t, updateArg.Name, updatedArmour.Name)
	require.Equal(t, updateArg.Description, updatedArmour.Description)
	require.Equal(t, updateArg.Price, updatedArmour.Price)
	require.Equal(t, updateArg.Slot, updatedArmour.Slot)
	require.Equal(t, updateArg.Origin, updatedArmour.Origin)
	require.Equal(t, updateArg.CaBonus, updatedArmour.CaBonus)
	require.Equal(t, updateArg.Penality, updatedArmour.Penality)
	require.Equal(t, updateArg.Category, updatedArmour.Category)
}

func TestDeleteArmourTx(t *testing.T) {
	store := NewStore(testDB)

	arg := createRandomArmour()
	armour, err := store.CreateArmourTx(context.Background(), arg)
	require.NoError(t, err)

	err = store.DeleteArmourTx(context.Background(), armour.ID)
	require.NoError(t, err)

	deletedArmour, err := store.GetArmourTx(context.Background(), armour.ID)
	require.Error(t, err)
	require.Empty(t, deletedArmour)
}

func createRandomItem() CreateItemParams {
	return CreateItemParams{
		Name:        "Cajado arcano",
		Description: "Descrição para teste do Cajado arcano",
		Category:    "Esotérico",
		Price:       1000,
		Slot:        2,
		Origin:      "T20 - Base",
	}
}

func createManyItems(t *testing.T) []Items {
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

	store := NewStore(testDB)
	var listItems []Items

	for _, i := range excel {
		arg := CreateItemParams{
			Name:        i.Name,
			Description: i.Description,
			Category:    i.Category,
			Price:       i.Price,
			Slot:        i.Slot,
			Origin:      i.Origin,
		}

		item, err := store.CreateItemTx(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, item)

		require.Equal(t, arg.Name, item.Name)
		require.Equal(t, arg.Description, item.Description)
		require.Equal(t, arg.Category, item.Category)
		require.Equal(t, arg.Price, item.Price)
		require.Equal(t, arg.Slot, item.Slot)

		require.NotZero(t, item.ID)

		listItems = append(listItems, item)
	}

	return listItems
}

func TestCreateItemTx(t *testing.T) {
	store := NewStore(testDB)

	arg := createRandomItem()

	item, err := store.CreateItemTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, arg.Name, item.Name)
	require.Equal(t, arg.Description, item.Description)
	require.Equal(t, arg.Price, item.Price)
	require.Equal(t, arg.Slot, item.Slot)
	require.Equal(t, arg.Origin, item.Origin)
	require.Equal(t, arg.Category, item.Category)

	require.NotZero(t, item.ID)
}

func TestGetItemTx(t *testing.T) {
	store := NewStore(testDB)

	arg := createRandomItem()
	item, err := store.CreateItemTx(context.Background(), arg)
	require.NoError(t, err)

	fetchedItem, err := store.GetItemTx(context.Background(), item.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedItem)

	require.Equal(t, item.ID, fetchedItem.ID)
	require.Equal(t, item.Name, fetchedItem.Name)
	require.Equal(t, item.Description, fetchedItem.Description)
	require.Equal(t, item.Price, fetchedItem.Price)
	require.Equal(t, item.Slot, fetchedItem.Slot)
	require.Equal(t, item.Origin, fetchedItem.Origin)
	require.Equal(t, item.Category, fetchedItem.Category)
}

func TestListAllItemsTx(t *testing.T) {
	store := NewStore(testDB)

	_ = createManyItems(t)
	arg := ListAllItemsParams{
		Limit:  10,
		Offset: 5,
	}

	items, err := store.ListAllItemsTx(context.Background(), arg)
	require.NoError(t, err)

	for _, item := range items {
		require.NotEmpty(t, item)
	}
}

func TestListItemsByCategoryTx(t *testing.T) {
	store := NewStore(testDB)

	aux := createManyItems(t)

	args := ListItemsByCategoryParams{
		Category: aux[0].Category,
		Limit:    10,
		Offset:   5,
	}

	items, err := store.ListItemsByCategoryTx(context.Background(), args)
	require.NoError(t, err)

	for _, item := range items {
		require.NotEmpty(t, item)
		require.Equal(t, aux[0].Category, item.Category)
	}
}

func TestUpdateItemTx(t *testing.T) {
	store := NewStore(testDB)

	arg := createRandomItem()
	item, err := store.CreateItemTx(context.Background(), arg)
	require.NoError(t, err)

	updateArg := UpdateItemParams{
		ID:          item.ID,
		Name:        item.Name + "_updated",
		Description: item.Description + "_updated",
		Price:       item.Price + 50,
		Slot:        item.Slot + 1,
		Origin:      item.Origin + "_updated",
		Category:    item.Category + "_update",
	}

	updatedItem, err := store.UpdateItemTx(context.Background(), updateArg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedItem)

	require.Equal(t, updateArg.ID, updatedItem.ID)
	require.Equal(t, updateArg.Name, updatedItem.Name)
	require.Equal(t, updateArg.Description, updatedItem.Description)
	require.Equal(t, updateArg.Price, updatedItem.Price)
	require.Equal(t, updateArg.Slot, updatedItem.Slot)
	require.Equal(t, updateArg.Origin, updatedItem.Origin)
	require.Equal(t, updateArg.Category, updatedItem.Category)
}

func TestDeleteItemTx(t *testing.T) {
	store := NewStore(testDB)

	arg := createRandomItem()
	item, err := store.CreateItemTx(context.Background(), arg)
	require.NoError(t, err)

	err = store.DeleteItemTx(context.Background(), item.ID)
	require.NoError(t, err)

	deletedItem, err := store.GetItemTx(context.Background(), item.ID)
	require.Error(t, err)
	require.Empty(t, deletedItem)
}
