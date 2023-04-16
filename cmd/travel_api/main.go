package main

import (
	"travel-api/api/v1/router"
	"travel-api/internal/model"
)

func main() {
	model.ConnectDatabase()
	router.CreateRouter()
}
