package main

import (
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "github.com/kuromike0629/shogi-engine-docker/api/handler"
)

func main() {
  e := echo.New()

  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.Get("/hello", handler.MainPage())

  e.Start(":4459")
}
