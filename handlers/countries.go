package handlers

import (
	"be-task2/external"
	"be-task2/helpers"
	"be-task2/models"
	"fmt"
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

 fmt.Println(exchangeRates)
  
 return

}