package main

import (
	"golang-restaurant-management/database"
	"golang-restaurant-management/routes"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	// routes.UserRoutes(router)
	// router.Use(middlewares.Authentication())

	routes.FoodRoutes(router)
	// routes.MenuRoutes(router)
	// routes.TableRoutes(router)
	// routes.OrderRoutes(router)
	// routes.OrderItemRoutes(router)
	// routes.InvoiceRoutes(router)

	router.Run(":" + port)
}
