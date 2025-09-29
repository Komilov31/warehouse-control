package main

import (
	"log"
	"wharehouse-control/cmd/app"
	_ "wharehouse-control/docs"
)

// @title Warehouse Control API
// @version 1.0
// @description Api to create and manage items in warehouse
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Тип аутентификации - Bearer token. В поле авторизации введите: Bearer {token}
func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
