package main

import (
	"travel-api/api/v1/router"
	"travel-api/internal/model/datasource"
	"travel-api/locale"
)

func main() {
	datasource.Setup()
	locale.Setup()
	router.NewGin().CreateRouter()
}
