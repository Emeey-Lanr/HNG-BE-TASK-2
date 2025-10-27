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

}