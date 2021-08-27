package main

import (
	"api-go/configs"
	_ "api-go/docs"
	services "api-go/service"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

// @title Coursera RESTful API
// @version 1.0
// @description This is a Coursera API server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email long.chu@savvycomsoftware.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	app := echo.New()

	_, err := configs.Connect()
	if err != nil {
		panic(err.Error)
	}

	app.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello API!")
	})
	apiRoutes := app.Group("/v1")
	courseRoutes := apiRoutes.Group("/courses")
	courseRoutes.GET("/", services.GetAll)
	courseRoutes.GET("/:id", services.GetOne)
	courseRoutes.POST("/", services.Create)
	courseRoutes.PUT("/:id", services.Update)
	courseRoutes.DELETE("/:id", services.Delete)

	app.GET("/swagger/*", echoSwagger.WrapHandler)

	app.Logger.Fatal(app.Start(":8080"))
}
