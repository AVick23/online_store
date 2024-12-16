package main

import (
	"log"

	"github.com/AVick23/online_store/database"
	"github.com/AVick23/online_store/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Не получилось подключиться к БД: %v", err)
	}
	defer db.Close()

	routes.SetupRoutes(db, r)

	r.Run("0.0.0.0:8080")
}
