package routes

import (
	"golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/menus", controllers.GetMenus())
	incommingRoutes.GET("/menus/:menu_id", controllers.GetMenu())
	incommingRoutes.POST("/menus", controllers.CreateMenu())
	incommingRoutes.PATCH("/menus/:menu_id", controllers.UpdateMenu())
}
