package external

import (
	
	"be-task2/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"


)

func GetCountries (countries *[]models.Countries) (int, error) {
	

	client := &http.Client{Timeout:10 * time.Second} 

	request, err := http.NewRequest(http.MethodGet, os.Getenv("COUNTRY_API"), nil)

	if err != nil{

	   return http.StatusInternalServerError, fmt.Errorf("Internal Server Error:Failed to create request object")
	}

	
	response, err := client.Do(request)
    if err != nil{
    return http.StatusInternalServerError, fmt.Errorf("Internal data source unavailable:Couldn't fetch data from RestCountries API")
	}

 defer response.Body.Close()

 if response.StatusCode != http.StatusOK{

	   return http.StatusServiceUnavailable, fmt.Errorf("Internal data source unavailable:RestCountriesAPI returned %d", response.StatusCode)
 }


//  Decode data to a struct
	if err := json.NewDecoder(response.Body).Decode(countries); err != nil{
		 return http.StatusServiceUnavailable, fmt.Errorf("Internal data source unavailable:Invalid JSON response from RestCountriesAPI")
	}


  return 200, nil
}


func GetExchangeRate  (rates *models.ExchangeRate)( int, error ){


	client := &http.Client{Timeout:10 * time.Second}
	

	request, err := http.NewRequest(http.MethodGet, os.Getenv("EXCHANGERATE_API"), nil)

	if err != nil{
  return http.StatusInternalServerError, fmt.Errorf("Internal Server Error:Failed to create request object")
	}

	response, err := client.Do(request)

	if  err != nil{
	  return http.StatusInternalServerError, fmt.Errorf("Internal data source unavailable:Couldn't fetch data from Exchange Rates API")
	}

    defer  response.Body.Close()

	 if response.StatusCode != http.StatusOK{

	   return http.StatusServiceUnavailable, fmt.Errorf("Internal data source unavailable:RestCountriesAPI returned %d", response.StatusCode)
 }


	if err := json.NewDecoder(response.Body).Decode(rates); err != nil{
	  return http.StatusServiceUnavailable, fmt.Errorf("Internal data source unavailable:Invalid JSON response from Exchange Rates API")
	}

	return 200, nil
}