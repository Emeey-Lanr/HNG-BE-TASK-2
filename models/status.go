package models


import (
	"time"
)

type Status struct {
	TotalCountries int `db:"total_countries" json:"total_countries"`
	LastRefreshedAt time.Time `db:"last_refreshed_at" json:"last_refreshed_at"`
}