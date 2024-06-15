package products

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type ProductController struct {
	service *ProductService
}

func NewProductController(service *ProductService) *ProductController {
	return &ProductController{service: service}
}

func (p *ProductController) List(w http.ResponseWriter, req *http.Request) {
	products, err := p.service.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *ProductController) Get(w http.ResponseWriter, req *http.Request) {
	idRaw := strings.TrimPrefix(req.URL.Path, "/products/")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := p.service.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *ProductController) Create(w http.ResponseWriter, req *http.Request) {
	var product Product
	err := json.NewDecoder(req.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if product.Name == "" || product.Description == "" || product.Price == 0 || product.Categoria_idcategoria == 0 {
		http.Error(w, "Todos os campos são obrigatórios", http.StatusBadRequest)
		return
	}

	newProduct, err := p.service.Create(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *ProductController) Update(w http.ResponseWriter, req *http.Request) {
	idRaw := strings.TrimPrefix(req.URL.Path, "/products/update/")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var product Product
	err = json.NewDecoder(req.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product.Id = int64(id)

	err = p.service.Update(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *ProductController) Delete(w http.ResponseWriter, req *http.Request) {
	idRaw := strings.TrimPrefix(req.URL.Path, "/products/delete/")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = p.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
