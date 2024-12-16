package routes

import (
	"database/sql"

	"github.com/AVick23/online_store/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *sql.DB, r *gin.Engine) {
	r.GET("/products/all", func(c *gin.Context) {
		controllers.GetProducts(db, c)
	})
	r.GET("/products/:id", func(c *gin.Context) {
		controllers.GetProductsById(db, c)
	})
	r.POST("/products", func(c *gin.Context) {
		controllers.CreateProduct(db, c)
	})
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/porducts/:id", controllers.DeleteProduct)

}
