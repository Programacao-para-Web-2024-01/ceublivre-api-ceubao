package products

import "net/http"

type ProductController struct {
	service *ProductService
}

func NewProductController(service *ProductService) *ProductController {
	return &ProductController{
		service: service,
	}
}

func (pc *ProductController) List(w http.ResponseWriter, r *http.Request) {
	pc.service.List(w, r)
}

func (pc *ProductController) Get(w http.ResponseWriter, r *http.Request) {
	pc.service.Get(w, r)
}

func (pc *ProductController) Create(w http.ResponseWriter, r *http.Request) {
	pc.service.Create(w, r)
}

func (pc *ProductController) Update(w http.ResponseWriter, r *http.Request) {
	pc.service.Update(w, r)
}

func (pc *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	pc.service.Delete(w, r)
}
