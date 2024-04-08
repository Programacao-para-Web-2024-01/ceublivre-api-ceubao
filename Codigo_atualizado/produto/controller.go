package produto

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProdutoController struct {
	produtoService ProdutoService
}

func NewProdutoController(produtoService ProdutoService) *ProdutoController {
	return &ProdutoController{produtoService}
}

func (pc *ProdutoController) List(w http.ResponseWriter, r *http.Request) {
	produtos, err := pc.produtoService.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(produtos)
}

func (pc *ProdutoController) Create(w http.ResponseWriter, r *http.Request) {
	var produto Produto
	err := json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		http.Error(w, "Erro ao decodificar o JSON do produto", http.StatusBadRequest)
		return
	}

	id, err := pc.produtoService.Create(produto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(id)))
}

func (pc *ProdutoController) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID do produto inválido", http.StatusBadRequest)
		return
	}

	produto, err := pc.produtoService.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(produto)
}

func (pc *ProdutoController) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID do produto inválido", http.StatusBadRequest)
		return
	}

	var produto Produto
	err = json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		http.Error(w, "Erro ao decodificar o JSON do produto", http.StatusBadRequest)
		return
	}

	err = pc.produtoService.Update(id, produto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (pc *ProdutoController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID do produto inválido", http.StatusBadRequest)
		return
	}

	err = pc.produtoService.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
