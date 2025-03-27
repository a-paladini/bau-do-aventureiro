package db

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func ParseInt(val string) int32 {
	s, err := strconv.Atoi(val)
	if err == nil {
		return int32(s)
	}

	return 0
}

func ParseFloat(val string) float64 {
	s, err := strconv.ParseFloat(val, 64)
	if err == nil {
		return s
	}

	return 0
}

func ReadExcelSheets(filePath string) (map[string][][]string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sheetList := f.GetSheetList()
	result := make(map[string][][]string)

	for _, sheetName := range sheetList {
		rows, err := f.GetRows(sheetName)
		if err != nil {
			return nil, fmt.Errorf("error reading sheet %s: %v", sheetName, err)
		}
		result[sheetName] = rows
	}

	return result, nil
}

func ProcessExcelDataWeapons(sheets map[string][][]string) ([]Weapons, error) {
	var weapons []Weapons

	if weaponsRows, ok := sheets["Armas"]; ok {
		for i, row := range weaponsRows {
			if i == 0 {
				continue
			}

			weapons = append(weapons, Weapons{
				Name:        row[0],
				Price:       ParseInt(row[1]),
				Damage:      row[2],
				Critical:    row[3],
				Range:       row[4],
				TypeDamage:  row[5],
				Slot:        ParseFloat(row[6]),
				Property:    row[7],
				Proficiency: row[8],
				Special:     sql.NullString{String: row[9], Valid: row[9] != ""},
				Origin:      row[10],
				Description: row[11],
			})
		}
	}

	return weapons, nil
}

func ProcessExcelDataArmours(sheets map[string][][]string) ([]Armours, error) {
	var armours []Armours

	if armoursRows, ok := sheets["Armaduras"]; ok {
		for i, row := range armoursRows {
			if i == 0 {
				continue
			}

			armours = append(armours, Armours{
				Name:        row[0],
				Price:       ParseInt(row[1]),
				Slot:        ParseFloat(row[2]),
				Category:    row[3],
				CaBonus:     ParseInt(row[4]),
				Penality:    ParseInt(row[5]),
				Origin:      row[6],
				Description: row[7],
			})
		}
	}

	return armours, nil
}

func ProcessExcelDataItems(sheets map[string][][]string) ([]Items, error) {
	var items []Items

	if itemsRows, ok := sheets["Itens"]; ok {
		for i, row := range itemsRows {
			if i == 0 {
				continue
			}

			items = append(items, Items{
				Name:        row[0],
				Category:    row[1],
				Price:       ParseInt(row[2]),
				Slot:        ParseFloat(row[3]),
				Origin:      row[4],
				Description: row[5],
			})
		}
	}

	return items, nil
}
