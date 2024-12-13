package main

import (
	"github.com/AVick23/online_store/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run("0.0.0.0:8080")
}
