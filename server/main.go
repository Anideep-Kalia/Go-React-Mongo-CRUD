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
	router.GET("entries",routes.GetEnteries)
	router.GET("entry/:id/",routes.GetEntryById)
	router.GET("/ingredient/:ingredient", routes.GetEntriesByIngredient)

	router.PUT("/entry/update/:id",routes.UpdateEntry)
	router.PUT("/ingrdient/update/:id", routes.UpdateIngredient)
	router.DELETE("/entry/delete/:id", routes.DeleteEntry)
	router.Run(":" + port)
}