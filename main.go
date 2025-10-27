package main

import (
	"be-task2/config"
	"be-task2/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)





func main (){

	r := gin.Default()

	// Godotenv
	if err := godotenv.Load(); err != nil {
		log.Println("error loading err:", err.Error())
		return
	 }

	//  Database
    db := config.ConnectoDb()
   

	// Routes
	CountryRoutes := r.Group("/countries")  

	routes.CountryRoutes(CountryRoutes, db)

     defer db.Close()
	

	r.Run(`:8080`)
}