package main

import (
	"aula-database/products"
	"database/sql"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

func connectDB() *sql.DB {
	config := mysql.NewConfig()
	config.User = "root"
	config.Passwd = "12345678"
	config.DBName = "market"
	config.Net = "tcp"
	config.Addr = "127.0.0.1:3306"
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

	// Rota para deletar um produto por ID
	mux.HandleFunc("/products/delete/", productController.Delete) // Delete a product by id

	// Adicione middleware para habilitar o CORS
	handler := allowCORS(mux)

	return http.ListenAndServe("localhost:8080", handler)
}

// Middleware para habilitar o CORS
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func main() {
	if err := createServer(); err != nil {
		log.Panic(err)
	}
}
