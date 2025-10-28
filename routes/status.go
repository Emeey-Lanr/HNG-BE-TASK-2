package routes

import (
	"be-task2/handlers"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func StatusRoute (r *gin.Engine, db *sqlx.DB){
	r.GET("/status",func(ctx *gin.Context) {
    handlers.GetStatus(ctx, db)
  }) 
}