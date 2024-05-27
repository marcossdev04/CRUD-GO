package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/marcossdev04/crud-go/src/configuration/logger"
	"github.com/marcossdev04/crud-go/src/controller/routes"
)

func main() {
	logger.Info("About to start user aplication")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":3333"); err != nil {
		log.Fatal(err)
	}
}
