package routes

import (
	"be-task2/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)


func CountryRoutes  (rg *gin.RouterGroup, db *sqlx.DB){
  rg.POST("/refresh", func(ctx *gin.Context) {
	handlers.AddCountriesToDb(ctx, db)
  })	

//   get all countries from the db
  rg.GET("", func(ctx *gin.Context) {
     handlers.SortAndFilterCountry(ctx, db)
  })

  // get one country by name
  rg.GET("/:name", func(ctx *gin.Context) {

  })

//   delete a country record
  rg.DELETE("/:name", func(ctx *gin.Context) {

  })
  

}