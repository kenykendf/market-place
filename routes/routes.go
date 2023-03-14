package routes

import (
	"superindo/controllers/categories"
	"superindo/controllers/items"
	"superindo/controllers/orders"
	user "superindo/controllers/users"
	"superindo/middlewares"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	r.POST("/user/login", user.LoginUser)
	r.POST("/user", user.CreatorUser)

	r.Use(middlewares.Authentication())
	r.GET("/users", user.ListsUsers)
	r.GET("/user/:userID", middlewares.UserAuthorization(), user.DetailUser)
	r.PUT("/user/:userID", middlewares.UserAuthorization(), user.UpdaterUser)
	r.DELETE("/user/:userID", middlewares.UserAuthorization(), user.DeleteUser)

	r.POST("/order", orders.CreateOrder)
	r.GET("/orders", orders.ListsOrders)
	r.GET("/order/:orderID", orders.DetailOrder)
	r.DELETE("/order/:orderID", orders.DeleteOrder)

	r.POST("/item", items.CreateItem)
	r.GET("/items", items.ListsItems)
	r.GET("/item/:itemID", items.DetailItem)

	r.POST("/category", categories.CreateCategories)
	r.GET("/categories", categories.ListCategories)
	return r
}
