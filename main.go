package main

import (
	"myapp/router"
)

func main() {
	router := router.SetupRouter()
	router.Run() // listen and serve on 0.0.0.0:8080
}
