package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/Splash07/go_user/config"
	"github.com/Splash07/go_user/route"
	mw "github.com/Splash07/go_user/middleware"
)

func main() {
	fmt.Printf("config app: %+v", config.Config)
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(mw.SimpleLogger())
	route.All(e)

	log.Println(e.Start(":9090"))
}
