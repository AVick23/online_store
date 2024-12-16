package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/AVick23/online_store/database"
	"github.com/AVick23/online_store/models"
	"github.com/gin-gonic/gin"
)

var products = []models.Product{
	{ID: 1, Name: "product_1", Price: 10.99},
	{ID: 2, Name: "product_2", Price: 20.99},
}

func GetProducts(db *sql.DB, c *gin.Context) {
	products, err := database.GetAllProducts(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func GetProductsById(db *sql.DB, c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid product ID"})
		return
	}

	product, err := database.GetIdProduct(db, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func CreateProduct(db *sql.DB, c *gin.Context) {
	var newProduct models.Productss
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	productID, err := database.CreateProductDb(db, newProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failes to sabe product", "error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully", "product_id": productID, "product": newProduct})
}

func UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid product ID"})
		return
	}

	var updateProduct models.Product
	if err := c.ShouldBindJSON(&updateProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	for i, product := range products {
		if product.ID == id {
			products[i] = updateProduct
			c.JSON(http.StatusOK, updateProduct)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
}

func DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid product ID"})
		return
	}

	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
}
