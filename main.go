package main

import (
	"log"
	"net/http"
	"rescues/infrastructure"
	"rescues/router"
)

// @title Swagger UI for E-Rescues
// @version 1.0
// @description API lists for E-Rescues

// @host localhost:19000
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	log.Println("Database name: ", infrastructure.GetDBName())
	log.Fatal(http.ListenAndServe(":"+infrastructure.GetAppPort(), router.Router()))

}