package models

type Error struct {
	Category    string `json:"category"`
	Subcategory string `json:"subcategory"`
	Frequency   int    `json:"frequency"`
}
type CategoryErrors struct {
	Category string  `json:"error_category"`
	Errors   []Error `json:"errors"`
}
