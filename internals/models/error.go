package models

type Error struct {
	Category    string `json:"category"`
	Subcategory string `json:"subcategory"`
	Count       int    `json:"count"`
}
