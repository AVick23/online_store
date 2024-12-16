package database

import (
	"database/sql"
	"fmt"

	"github.com/AVick23/online_store/models"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func ConnectDB() (*sql.DB, error) {
	connStr := "postgres://avick23:212226@127.0.0.1:5432/database?sslmode=disable"
	return sql.Open("postgres", connStr)
}

func GetAllProducts(db *sql.DB) ([]models.Products, error) {
	rows, err := db.Query("SELECT id, name, description, price, brand FROM products")
	if err != nil {
		return nil, fmt.Errorf("произошла ошибка %v", err)
	}
	defer rows.Close()

	var products []models.Products
	for rows.Next() {
		var p models.Products
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Brand)
		if err != nil {
			return nil, fmt.Errorf("неудалось сохранить данные %v", err)
		}
		products = append(products, p)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func GetIdProduct(db *sql.DB, id int) (*models.Products, error) {
	rows, err := db.Query(`SELECT name, description, price, brand, category_id FROM products WHERE id = $1`, id)
	if err != nil {
		return nil, fmt.Errorf("произошла ошибка %v", err)
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, fmt.Errorf("продукт с этим id не найден %v", err)
	}

	var product models.Products
	err = rows.Scan(&product.Name, &product.Description, &product.Price, &product.Brand, &product.CategoryID)
	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе продукта %v", err)
	}

	return &product, nil
}

func CreateProductDb(db *sql.DB, newProduct models.Productss) (int, error) {
	var productid int
	err := db.QueryRow(`INSERT INTO products (name, description, price, brand) VALUES ($1, $2, $3, $4) RETURNING id`, newProduct.Name, newProduct.Description, newProduct.Price, newProduct.Brand).Scan(&productid)
	if err != nil {
		return 0, fmt.Errorf("произошла ошибка при добовление в бд %v", err)
	}
	return productid, nil
}
