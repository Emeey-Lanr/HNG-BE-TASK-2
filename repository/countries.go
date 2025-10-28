package repository

import (
	"be-task2/models"
	"database/sql"
	"errors"
	"fmt"
	"strings"

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



func GetImageSummaryFromDB (db *sqlx.DB)( models.Status, []models.TopGDP, error){
	var totalLastRefreshed models.Status
	var top5 []models.TopGDP

     err := db.Get(&totalLastRefreshed, `SELECT COUNT(*) AS total_countries, MAX(last_refreshed_at) AS last_refreshed_at  FROM countries`)
	if  err != nil {
		return  totalLastRefreshed, top5, fmt.Errorf("Error Getting Total Country %w", err)
	}


	topCountryErr := db.Select(&top5, `SELECT name, estimated_gdp FROM countries ORDER BY estimated_gdp DESC LIMIT 5`)

	if topCountryErr != nil {
       	return  totalLastRefreshed, top5, fmt.Errorf("Error Getting Total Country %w", err)
	}
	
	
    return   totalLastRefreshed,  top5, nil
}



func SortAndFilterDBQuery (db *sqlx.DB, region, currency, sort string) ([]models.DBData, error){
	
	query := `SELECT name, capital, region, population, currency_code, exchange_rate,
	 estimated_gdp, flag_url, last_refreshed_at FROM countries WHERE 1=1`

	 
	 parameter := []interface{}{}



	 if region != "" {
		
	 query += " AND region = ?"

	 //  changed region from to Upperlower..
	 Region := strings.ToUpper(region[:1]) + strings.ToLower(region[1:])

	 parameter = append(parameter,  Region)
	 }

	 
	
	 if currency != "" {
		 query += " AND currency_code = ?"
		 parameter = append(parameter, strings.ToUpper(currency))
	 }



	 // sorting
	 switch sort {
	 case "gdp_asc":
		query += " ORDER BY estimate_gdp ASC"
	 case "gdp_desc":
		query += " ORDER BY estimate_gdp DESC"
	 case "population_desc":
		query += " ORDER BY population DESC"
		 case "population_asc":
		query += " ORDER BY population ASC"
		 default:
			query += " ORDER BY name ASC"

	 }


	 var selectedCountries [] models.DBData

	 if err := db.Select(&selectedCountries, query, parameter...); err != nil{
		return nil, fmt.Errorf("Failed to fetch country data %w", err)
	 }

	 return selectedCountries, nil

}



func SelectASingleCountry (name string, db *sqlx.DB)( models.DBData, error){

	query := `SELECT* FROM countries WHERE name = ? LIMIT 1`


	  var country models.DBData

	  if err := db.Get(&country, query, name); err != nil{
		if errors.Is(err, sql.ErrNoRows){
			return models.DBData{}, fmt.Errorf("country not found")
		}

		return models.DBData{}, fmt.Errorf("Failed to fetch country %w", err)

	  }
	  

 return country,  nil
}


func DeleteACountryFromDB (name string, db *sqlx.DB) error{
	query := `DELETE FROM countries WHERE LOWER(name) = LOWER(?) `
	
	deleted, err := db.Exec(query, name)

	if  err != nil{
		return fmt.Errorf ("Failed to delete country %w", err)
	}

	rowsAffected, _ := deleted.RowsAffected()
	if rowsAffected == 0{
		return fmt.Errorf ("country not found")
	}

	return nil
}