package main

import (
	"github.com/gin-gonic/gin"
	"lab1/controllers"
	"lab1/db"
)

func main() {
	db.Init()
	r := gin.Default()

	//Store
	r.POST("/store/add", controllers.AddItem)
	r.GET("/store/get", controllers.GetItems)
	r.GET("/store/get/:id", controllers.GetItem)
	r.GET("/store/search", controllers.SearchItems)
	r.DELETE("/store/remove/:id", controllers.RemoveItem)
	r.PUT("/store/update", controllers.UpdateItem)

	// Users
	r.POST("/user/add", controllers.AddUser)
	r.POST("/user/rate", controllers.RateItem)
	r.POST("/user/:id/saveItem", controllers.SaveItem)
	r.DELETE("/user/:id/removeItem", controllers.RemoveSavedItem)
	r.GET("/user/:id/getSavedItems", controllers.GetSavedItems)
	r.GET("/user/get", controllers.GetUsers)
	r.GET("/user/get/:name", controllers.GetUserByName)
	r.GET("/user/login", controllers.LoginUser)
	r.DELETE("/user/remove/:id", controllers.RemoveUser)
	r.PUT("/user/update/:id", controllers.UpdateUser)

	r.Run()
}
