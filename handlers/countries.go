package handlers

import (
	"be-task2/external"
	"be-task2/helpers"
	"be-task2/models"
	"be-task2/repository"
	"be-task2/services"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func AddCountriesToDb(c *gin.Context, db *sqlx.DB) {
   var countries []models.Countries
 var  exchangeRates models.ExchangeRate

  countryReturnedMethod, countryReturnedErr :=  external.GetCountries(&countries)

  if countryReturnedErr != nil{
	helpers.ErrorResponse(countryReturnedMethod, models.ErrorResp{Error: strings.Split(countryReturnedErr.Error(), ":")[0], Details:strings.Split(countryReturnedErr.Error(), ":")[1]}, c)
	return
  }

  fmt.Println(countries[:5])
  

  exchangeReturnedMethod, exchangeReturnedErr  := external.GetExchangeRate(&exchangeRates)

   if exchangeReturnedErr != nil{
	helpers.ErrorResponse(exchangeReturnedMethod, models.ErrorResp{Error: strings.Split(exchangeReturnedErr.Error(), ":")[0], Details:strings.Split(exchangeReturnedErr.Error(), ":")[1]}, c)
	return
  }

  fmt.Println(exchangeRates.Rates)
//   Handles appending the correct data set into a slice

  
  countriesDBData, err := services.SetCountryDBData(countries, exchangeRates)
  
 

  if err != nil{
	helpers.ErrorResponse(400, models.ErrorResp{Error:"Validation failed", 
	Details:"is required" }, c)
  return
  }

  fmt.Println(countriesDBData[:9])
 
  if err := repository.AddCountriesToDB(countriesDBData, db); err != nil{
    helpers.ErrorResponse(http.StatusInternalServerError, models.ErrorResp{Error: "Internal server error", Details: err.Error()}, c)
   return
  }


  // Get Image Summary from DB

   totalAndLastRefreshed, topGDP, err := repository.GetImageSummaryFromDB(db)

  if  err != nil {
     log.Println(err.Error())
  }

  errFromImage := services.CreateImage(totalAndLastRefreshed.TotalCountries, topGDP, totalAndLastRefreshed.LastRefreshedAt)

  if errFromImage != nil{
     log.Println(errFromImage.Error())
  }
  
  helpers.SuccessResponse(http.StatusOK, gin.H{"message":"refresh succesfull"}, c)

}



func SortAndFilterCountry (c *gin.Context, db *sqlx.DB){

   region :=  c.Query("region")
   currency := c.Query("currency")
   sort := c.Query("sort")

   data, err := repository.SortAndFilterDBQuery(db, region, currency, sort)

   if err != nil {
    helpers.ErrorResponse(http.StatusNotFound, models.ErrorResp{Error:"Data not found", Details: err.Error()}, c)
    return 
  }

 helpers.SuccessResponse(http.StatusOK, data, c)


}


func GetACountry (c *gin.Context, db *sqlx.DB){
    name := c.Param("name")
 
    data, err := repository.SelectASingleCountry(name, db)

    if err != nil {
            helpers.ErrorResponse(http.StatusNotFound, models.ErrorResp{Error:"Data not found", Details: err.Error()},c)
    return
          }

helpers.SuccessResponse(http.StatusOK, data, c)    


}

func DeleteACountry (c *gin.Context, db *sqlx.DB){
  name := c.Param("name")

  err := repository.DeleteACountryFromDB(name, db)
  if err != nil{
    helpers.ErrorResponse(http.StatusNotFound, models.ErrorResp{Error:"An occured deleting from db", Details: err.Error()},c)
  }

  helpers.SuccessResponse(http.StatusOK, gin.H{"message":"delted succesfully"}, c) 
}



func ServeSummaryImage (c *gin.Context){
 imagePath := "cache/summary.png"

 if _, err := os.Stat(imagePath); os.IsNotExist(err){
   helpers.ErrorResponse(http.StatusNotFound, models.ErrorResp{Error: "Summary image not found"}, c)
   return
 }

 //serve the file
 c.File(imagePath)

}