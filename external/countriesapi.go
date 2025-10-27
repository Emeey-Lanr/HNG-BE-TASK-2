package external

import (
	"be-task2/models"
	"encoding/json"
	"net/http"
	"time"
	"os"
)

func GetCountries () ( []models.Countries, error ){
	var countries []models.Countries

	client := &http.Client{Timeout:10 * time.Second}

	req, err := http.NewRequest(http.MethodGet, os.Getenv("COUNTRY_API"), nil)

	if err != nil{
      return nil,  err
	}

	
	res, err := client.Do(req)
    if err != nil{
       return nil, err
	}

 defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&countries); err != nil{
		return nil, err
	}

   return  countries, nil

}