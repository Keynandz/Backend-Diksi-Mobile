package main

import (
	"golang/handlers"
	"golang/middleware"
	"golang/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	storage.InitDB()
	storage.CreateDB()

	e.GET("/images/latest/:order", handlers.GetImageByTimestampOrder)
	e.GET("/logout", handlers.LogoutAkun)
	e.GET("/username", handlers.GetAkunByID)

	e.POST("/images", handlers.UploadImage)
	e.POST("/register", handlers.CreateAkun)
	e.POST("/login", handlers.LoginAkun)
	e.Use(middleware.LogRequest)
	e.Logger.Fatal(e.Start(":9000"))
}
