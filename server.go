package main

import (
	"admin-rt/config"
	"admin-rt/middleware"
	"admin-rt/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// db := config.GetDB()
	//when server.go start, it will be run function InitDB (connecting to database)
	// config.InitDB()
	config.GetDB()

	router := gin.Default()

	v1 := router.Group("api/v1") // /api/v1
	{
		v1.POST("/login", routes.Login) // /api/v1/login
		adminrt := v1.Group("/admin")   // /api/v1/sewaAset
		{
			adminrt.GET("/account", middleware.IsAuth(), routes.GetAccount) // /api/v1/admin/account
			adminrt.POST("/account", routes.PostAccount)                    // /api/v1/admin/account
			adminrt.PUT("/account/:id", middleware.IsAuth(), routes.UpdateAccount)
			adminrt.DELETE("/account/:id", middleware.IsAuth(), routes.DeleteAccount)

			adminrt.POST("/iuranSampah", middleware.IsAuth(), routes.PostIuranSampah) // /api/v1/admin/iuranSampah
		}
	}

	router.Run()
}
