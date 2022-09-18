package main

import (
	"assignment2go/route"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	serverAddress := os.Getenv("SERVICE_PORT")
	r := gin.Default()
	route.Route(r)

	r.Run(serverAddress)
}
