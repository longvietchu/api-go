package helpers

import (
	"api-go/configs"
	"github.com/labstack/echo/v4"
)

func NewServer() *configs.Server {
	s := &configs.Server{
		Echo: echo.New(),
		DB:   Init(),
	}
	return s
}
