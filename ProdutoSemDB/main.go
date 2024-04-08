package main

import (
	"aula-database/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	if err := createServer(); err != nil {
		log.Panic(err)
	}
}

func createServer() error {
	produtoRepository := db.NewProdutoRepository()

	mux := http.NewServeMux()

	mux.HandleFunc(
		"/produto",
		func(w http.ResponseWriter, req *http.Request) {
			switch req.Method {
			case "GET":
				produto, err := produtoRepository.List()
				if err != nil {
					http.Error(w, err.Error(), 500)
					return
				}

				err = json.NewEncoder(w).Encode(produto)
				if err != nil {
					http.Error(w, err.Error(), 500)
					return
				}

				w.WriteHeader(200)
			case "POST":
				var produto db.Produto
				err := json.NewDecoder(req.Body).Decode(&produto)
				if err != nil {
					http.Error(w, err.Error(), 400)
					return
				}

				id, err := produtoRepository.Create(produto)
				if err != nil {
					http.Error(w, err.Error(), 500)
					return
				}

				produto.Id = id

				err = json.NewEncoder(w).Encode(produto)
				if err != nil {
					http.Error(w, err.Error(), 500)
					return
				}
			default:
				http.Error(w, "method not supported", 405)
				return
			}

		})

	mux.HandleFunc(
		"/produto/{id}",
		func(w http.ResponseWriter, req *http.Request) {
			idRaw := req.PathValue("id")

			id, err := strconv.Atoi(idRaw)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}

			switch req.Method {
			case "GET":
				prodtu, err := produtoRepository.Get(id)
				if err != nil {
					http.Error(w, err.Error(), 500)
					return
				}

				err = json.NewEncoder(w).Encode(prodtu)
				if err != nil {
					http.Error(w, err.Error(), 500)
					return
				}
			case "PUT":
				var produto db.Produto
				err := json.NewDecoder(req.Body).Decode(&produto)
				if err != nil {
					http.Error(w, err.Error(), 400)
					return
				}

				err = produtoRepository.Update(id, produto)
				if err != nil {
					http.Error(w, err.Error(), 500)
					return
				}

				err = json.NewEncoder(w).Encode(produto)
				if err != nil {
					http.Error(w, err.Error(), 500)
					return
				}
			case "DELETE":
				err := produtoRepository.Delete(id)
				if err != nil {
					http.Error(w, err.Error(), 500)
					return
				}

				fmt.Fprint(w, "excluido com sucesso")
				w.WriteHeader(204)

			default:
				http.Error(w, "method not supported", 405)
				return
			}

		})

	return http.ListenAndServe("localhost:8080", mux)
}
