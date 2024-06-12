package main

import (
	"aula-database/products"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func connectDB() *sql.DB {
	config := mysql.NewConfig()
	config.User = "root"
	config.Passwd = "12345678"
	config.DBName = "market"
	conn, err := mysql.NewConnector(config)
	if err != nil {
		panic(err)
	}
	db := sql.OpenDB(conn)
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func createServer() error {
	db := connectDB()

	productRepository := products.NewProductRepository(db)
	productService := products.NewProductService(productRepository)
	productController := products.NewProductController(productService)

	mux := http.NewServeMux()

	mux.HandleFunc("/products", productController.List)           // List all products
	mux.HandleFunc("/products/", productController.Get)           // Get a product by id
	mux.HandleFunc("/products/create", productController.Create)  // Create a new product
	mux.HandleFunc("/products/update/", productController.Update) // Update a product by id
	mux.HandleFunc("/products/delete/", productController.Delete) // Delete a product by id

	return http.ListenAndServe("localhost:8080", mux)
}

func main() {
	if err := createServer(); err != nil {
		log.Panic(err)
	}
}
