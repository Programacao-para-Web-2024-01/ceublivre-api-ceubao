package Produto

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type ProdutoController struct {
	Controller *ProdutoController
}

func NewProdutoController(Controller *ProdutoController) *ProdutoController {
	return &ProdutoController{Controller: Controller}
}

func (s *ProdutoController) List(w http.ResponseWriter, req *http.Request) {
	Produtos, err := s.Controller.List()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.NewEncoder(w).Encode(Produtos)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (s *ProdutoController) Get(w http.ResponseWriter, req *http.Request) {
	// Input
	idRaw := req.PathValue("id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// processamento
	Produto, err := s.Controller.Get(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// output
	err = json.NewEncoder(w).Encode(Produto)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (s *ProdutoController) Create(w http.ResponseWriter, req *http.Request) {
	// Leitura do corpo (INPUT)
	var Produto db.Produto
	err := json.NewDecoder(req.Body).Decode(&Produto)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// Lógica da função/Algoritmo
	newProduto, err := s.Controller.Create(Produto)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Output / Resposta
	err = json.NewEncoder(w).Encode(newProduto)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (s *ProdutoController) Update(w http.ResponseWriter, req *http.Request) {
	// Input
	idRaw := req.PathValue("id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var Produto db.Produto
	err = json.NewDecoder(req.Body).Decode(&Produto)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	Produto.Id = int64(id)

	err = s.Controller.Update(Produto)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.NewEncoder(w).Encode(Produto)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

func (s *ProdutoController) Delete(w http.ResponseWriter, req *http.Request) {
	// Input
	idRaw := req.PathValue("id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = s.Controller.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(204)
}
