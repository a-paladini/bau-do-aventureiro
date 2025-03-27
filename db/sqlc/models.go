// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

type Armours struct {
	ID          int32   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Slot        float64 `json:"slot"`
	Origin      string  `json:"origin"`
	CaBonus     int32   `json:"ca_bonus"`
	Penality    int32   `json:"penality"`
}

type Items struct {
	ID          int32   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Slot        float64 `json:"slot"`
	Origin      string  `json:"origin"`
}

type Weapons struct {
	ID          int32   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Slot        float64 `json:"slot"`
	Origin      string  `json:"origin"`
	Damage      string  `json:"damage"`
	Critical    string  `json:"critical"`
	Range       string  `json:"range"`
	TypeDamage  string  `json:"type_damage"`
	Property    string  `json:"property"`
	Proficiency string  `json:"proficiency"`
	Special     string  `json:"special"`
}
