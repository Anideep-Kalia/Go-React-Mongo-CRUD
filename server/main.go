package main

import (
	"os"

	"github.com/Anideep-Kalia/GO-React-Mongo-Crud/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){
	println("hello")
	port := os.Getenv("PORT")
	
	if port == "" {
		port = "8000"
	}

	router :=gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/entry/create", routes.AddEntry)
}