package main

import (
	"fmt"
	"log"

	"github.com/Splash07/go_user/config"
	_ "github.com/Splash07/go_user/docs"
	mw "github.com/Splash07/go_user/middleware"
	"github.com/Splash07/go_user/route"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Go User API
// @version 1.0
// @description This documentation contains APIs of Go User
// @termsOfService http://swagger.io/terms/
// @host localhost:9090
// @BasePath /api/user/v1/
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	fmt.Printf("config app: %+v", config.Config)
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(mw.SimpleLogger())
	route.All(e)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	log.Println(e.Start(":9090"))
}
