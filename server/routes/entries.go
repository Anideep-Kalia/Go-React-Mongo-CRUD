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
func AddEntry(c *gin.Context) {				// c.JSON -> just like a response 
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)  	//context is package used to manage API calls like deadline, timeout etc..
	var entry models.Entry   														// from models folder

	if err := c.BindJSON(&entry); err != nil {										// incoming json body conversion to struct (not a ternary operation)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	validationErr := validate.Struct(entry)
	if validationErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validationErr.Error()})	//gin.H is used to mapping so message can be sent in tabular format
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

// Getting all enteries in the DB
func GetEntries(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var entries []bson.M											// slice of bson.M used to store the result of the request
	cursor, err := entryCollection.Find(ctx, bson.M{})				// empty bson.M{} means take all the enteries 

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &entries); err != nil {				// appending all results in the enetries variable
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	fmt.Println(entries)
	c.JSON(http.StatusOK, entries)
}