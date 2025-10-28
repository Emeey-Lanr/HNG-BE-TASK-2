package repository

import (
	"be-task2/models"

	"github.com/jmoiron/sqlx"
	"fmt"
)


func GetTotalCountryAndTimeStamp (db *sqlx.DB)(models.Status, error){
	var status models.Status

	query := "SELECT COUNT (*) AS total_countries, MAX(last_refreshed_at) AS last_refreshed_at FROM countries"

	 err := db.Get(&status, query)
	 
	 if err != nil {
		return models.Status{}, fmt.Errorf("Failed to get status:%w", err)
	 }

	 return status, nil
}