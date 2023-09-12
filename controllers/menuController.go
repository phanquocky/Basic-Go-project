package controllers

import (
	"context"
	"golang-restaurant-management/database"
	"golang-restaurant-management/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "gopkg.in/mgo.v2/bson"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var menus []models.Menu
		ctx, cancer := context.WithTimeout(context.Background(), 100*time.Second)
		cursor, err := menuCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		for cursor.Next(context.TODO()) {
			var menu models.Menu
			err := cursor.Decode(&menu)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err.Error())
			}

			menus = append(menus, menu)
		}

		c.JSON(http.StatusOK, menus)
		defer cancer()
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		menuId := c.Params.ByName("menu_id")
		var menu models.Menu
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": menuId}).Decode(&menu)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.JSON(http.StatusOK, menu)
		defer cancel()
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		var menu models.Menu
		err := c.BindJSON(&menu)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		validateErr := validate.Struct(menu)
		if validateErr != nil {
			c.JSON(http.StatusInternalServerError, validateErr.Error())
		}
		menu.ID = primitive.NewObjectID()
		menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		menu.Menu_id = menu.ID.Hex()
		result, err := menuCollection.InsertOne(ctx, menu)
		if err != nil {
			c.JSON(http.StatusInternalServerError, validateErr.Error())
		}
		defer cancel()

		c.JSON(http.StatusOK, result)
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu

		err := c.BindJSON(&menu)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		menuId := c.Param("menu_id")

		var updateObj primitive.D

		/*
			TODO: check Start_date and End_date is inTimeSpan
		*/
		// if menu.Start_date != nil && menu.End_date != nil {
		// }

		if menu.Name != "" {
			updateObj = append(updateObj, bson.E{Key: "name", Value: menu.Name})
		}

		if menu.Category != "" {
			updateObj = append(updateObj, bson.E{Key: "category", Value: menu.Category})
		}

		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: menu.Updated_at})

		result, err := menuCollection.UpdateOne(ctx, bson.M{"menu_id": menuId}, bson.D{{"$set", updateObj}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}
