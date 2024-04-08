package main

import (
	"Produto-Database/Produto"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	if err := createServer(); err != nil {
		log.Panic(err)
	}
}

func connectDB() *sql.DB {
	config := mysql.NewConfig()
	config.User = "root"
	config.Passwd = "12345678"
	config.DBName = "Catalogo_de_produtos"
	conn, err := mysql.NewConnector(config)
	if err != nil {
		panic(err)
	}
	return sql.OpenDB(conn)
}

func createServer() error {
	db := connectDB()

	ProdutoRepository := Produto.NewProdutoRepository(db)
	ProdutoService := Produto.NewProdutoService(ProdutoRepository)
	ProdutoController := Produto.NewProdutoController(ProdutoService)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /Produtos/", ProdutoController.List)
	mux.HandleFunc("POST /Produtos/", ProdutoController.Create)
	mux.HandleFunc("GET /Produtos/{id}", ProdutoController.Get)
	mux.HandleFunc("PUT /Produtos/{id}", ProdutoController.Update)
	mux.HandleFunc("DELETE /Produtos/{id}", ProdutoController.Delete)

	return http.ListenAndServe("localhost:8080", mux)
}
