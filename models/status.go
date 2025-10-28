package models


type Status struct {
	TotalCountries int `db:"total_countries" json:"total_countries"`
	LastRefreshedAt string `db:"last_refreshed_at" json:"last_refreshed_at"`
}