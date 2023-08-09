package main

import (
	"go-collab/cmd/handlers"
	"go-collab/cmd/middleware"
	"go-collab/cmd/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	storage.InitDB()
  
	e.GET("/logout", handlers.LogoutAkun)
	e.GET("/username", handlers.GetAkunByID)
  
	e.POST("/register", handlers.CreateAkun)
	e.POST("/login", handlers.LoginAkun)
	e.Use(middleware.LogRequest)
	e.Logger.Fatal(e.Start(":9000"))
  }