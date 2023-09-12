package controllers

import (
	"context"
	"fmt"
	"golang-restaurant-management/database"
	"golang-restaurant-management/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sqrthree/toFixed"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/bluesuncorp/validator.v5"
	"gopkg.in/mgo.v2/bson"
)

var validate = validator.New("validate", validator.BakedInValidators)
var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var foods []models.Food

		mongoCursor, err := foodCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		for mongoCursor.Next(context.TODO()) {
			var food models.Food
			if err := mongoCursor.Decode(&food); err != nil {
				c.JSON(http.StatusOK, foods)
			}
			foods = append(foods, food)
		}

		defer cancel()
		c.JSON(http.StatusOK, foods)
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food_id")

		var food models.Food
		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured when fetch one food "})
		}
		c.JSON(http.StatusOK, food)

	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		// var menu models.Menu
		var food models.Food

		if err := c.BindJSON(&food); err != nil {
			fmt.Println("Cannot bind json")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		validationErr := validate.Struct(food)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// TODO : do this when the menu controller done
		// err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// }

		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		var num = toFixed.ToFixed(food.Price, 2)

		food.Price = num

		result, insertErr := foodCollection.InsertOne(ctx, food)
		if insertErr != nil {
			msg := fmt.Sprintf("food item was not crearted!")
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}
