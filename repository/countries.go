package repository

import (
	"be-task2/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func AddCountriesToDB(countries []models.DBData, db *sqlx.DB) error{
	tx, err := db.Beginx()
	if err != nil {
     return fmt.Errorf("Error starting transaction", err)
	}

	defer func (){
		if p := recover(); p != nil{
			tx.Rollback()
			panic(p)
		}else if err != nil{
			tx.Rollback()
		}else {
			err = tx.Commit()
		}

	}()


	for _, c := range countries{
		_, err := tx.Exec(`INSERT INTO countries (name, capital, region, population, currency_code, exchange_rate, estimated_gdp, flag_url, last_refreshed_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, NOW()) ON
		DUPLICATE KEY UPDATE capital=VALUES(capital), region=VALUES(region), population=VALUES(population), currency_code=VALUES(currency_code),
		 exchange_rate=VALUES(exchange_rate), estimated_gdp=VALUES(estimated_gdp), flag_url=VALUES(flag_url), last_refreshed_at=VALUES(last_refreshed_at)
		
		`, c.Name, c.Capital, c.Region, c.Population, c.Currency_code, c.Exchange_rate, c.Estimated_gdp, c.Flag_url)
	 if err != nil{
		return fmt.Errorf("Failed Inserting Data %s: %w", c.Name, err)
	 }
	
	}

return nil

}