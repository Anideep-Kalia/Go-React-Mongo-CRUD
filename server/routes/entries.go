package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Anideep-Kalia/GO-React-Mongo-Crud/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Package level variables initialised when packages are imported
var validate = validator.New()
var entryCollection *mongo.Collection = OpenCollection(Client, "calories")

// adding enteries into DB
func AddEntry(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)  	//context is package used to manage API calls like deadline, timeout etc..
	var entry models.Entry   														// from models folder

	if err := c.BindJSON(&entry); err != nil {										// incoming json body conversion to struct (not a ternary operation)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	validationErr := validate.Struct(entry)
	if validationErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}
	entry.ID = primitive.NewObjectID()
	result, insertErr := entryCollection.InsertOne(ctx, entry)
	if insertErr != nil {
		msg := fmt.Sprintf("order item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(insertErr)
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result)
}

