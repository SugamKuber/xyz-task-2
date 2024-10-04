package models

type Error struct {
	Category    string `json:"category"`
	Subcategory string `json:"subcategory"`
	Frequency   int    `json:"frequency"`
}
