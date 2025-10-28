package services

import (
	"be-task2/models"
	"fmt"
	"image/color"
	"log"

	"math/rand"

	"github.com/fogleman/gg"
)

func SetCountryDBData(countries []models.Countries, rate models.ExchangeRate )([]models.DBData, error ){
     fmt.Println(countries[:4])
	var data []models.DBData



   
	 
     
	   type Country struct {
         CurrencyCode  *string
		 ExchangeRate  *float64 
		 EstimatedGDP  *float64
	   }

	for _, value := range countries {
 	random := rand.Intn(10001) + 1000

	   var 	ifNot Country
        
		if value.Name == ""{
		   continue
		}

		if value.Population <= 0 {
              continue
			  
		}

     
	
      // if it doesn't have currencies,
	  // it can't have an currency code to null, exchange rate null, estimated gdp to 0, 
		if  len(value.Currencies) == 0 {
			 
             ifNot.CurrencyCode = nil
		     estGdp :=  float64(0)
			 ifNot.EstimatedGDP = &estGdp
			 ifNot.ExchangeRate = nil

		}else{
			//let say it has a value currencies code but if the country code 
			// doesn't exist in  exchange rate api, t should set exchange  rate, estimated gbp to to nul  
			
				 _, exist := rate.Rates[value.Currencies[0].Code]
		 if !exist{
           ifNot.ExchangeRate = nil
		   ifNot.EstimatedGDP = nil
		   ifNot.CurrencyCode = &value.Currencies[0].Code
		 }else{

			ifNot.CurrencyCode = &value.Currencies[0].Code
			
			exchangeRate := rate.Rates[value.Currencies[0].Code]
			ifNot.ExchangeRate = &exchangeRate

             gdp := float64(value.Population) * float64(random) / rate.Rates[value.Currencies[0].Code]
		     ifNot.EstimatedGDP = &gdp
		 }
        
       
		
		}
		

		// if currency code is not found, set exchange rate to null,
		// set estimatedgbp to null
	
		
       countriesValue := models.DBData{
		Name: value.Name,
		 Capital: value.Capital,
		  Region: value.Region,
		  Population: value.Population, 
		  Currency_code:  ifNot.CurrencyCode,
		   Exchange_rate: ifNot.ExchangeRate,
		    Estimated_gdp: ifNot.EstimatedGDP,
		  Flag_url: value.Flag,
		}
      
		
		data = append(data, countriesValue)
	}

	return data, nil
}



func CreateImage (total int, top5 []models.TopGDP, lastRefreshed string)(error){

	const width = 800
	const height = 800

	dc := gg.NewContext(width, height)

	dc.Clear() // set background to whitw
	dc.SetColor(color.Black)
	y := 50.0

 dc.DrawStringAnchored(fmt.Sprintf("Total Countries: %d", total), width /2, y, 0.5, 0.5)
 y += 50
 
 dc.DrawStringAnchored("Top Five Countries by Estimated GDP:", width /2, y, 0.5, 0.5)
 y += 40

 for i, c := range top5 {
   if i >= 5 {
	break
   }

   dc.DrawStringAnchored(fmt.Sprintf("%d, %s - %.2f", i+1, c.Name, c.EstimatedGDP), width /2, y, 0.5, 0.5)
   y+= 30
 }

 y += 20
dc.DrawStringAnchored(fmt.Sprintf("Last Refreshed:%s", lastRefreshed), width /2, y, 0.5, 0.5)

// save image 
err := dc.SavePNG("cache/summary.png")
if err != nil {
	return fmt.Errorf("Unable to save image")
}

 log.Println("Image saved successfully")

 return nil

}