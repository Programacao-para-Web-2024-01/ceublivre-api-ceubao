package main

import (
	"aula-database/produto"
	"database/sql"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

func main() {
	if err := createServer(); err != nil {
		log.Panic(err)
	}
}

func connectDB() (*sql.DB, error) {
	config := mysql.NewConfig()
	config.User = "root"
	config.Passwd = "12345678"
	config.DBName = "CatalogoDeProdutos"
	conn, err := mysql.NewConnector(config)
	if err != nil {
		return nil, err
	}
	return sql.OpenDB(conn), nil
}

func createServer() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	produtoRepo := produto.NewProdutoRepository(db)
	produtoService := produto.NewProdutoService(produtoRepo)
	produtoController := produto.NewProdutoController(produtoService)

	mux := http.NewServeMux()

	mux.HandleFunc("/produtos/", produtoController.List)
	mux.HandleFunc("/produtos/", produtoController.Create)
	mux.HandleFunc("/produtos/", produtoController.Get)
	mux.HandleFunc("/produtos/", produtoController.Update)
	mux.HandleFunc("/produtos/", produtoController.Delete)

	log.Println("Servidor iniciado em localhost:8080")
	return http.ListenAndServe(":8080", mux)
}
