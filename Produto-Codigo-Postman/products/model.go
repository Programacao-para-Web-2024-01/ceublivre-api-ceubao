package products

type Product struct {
	Id                    int64   `json:"id"`
	Name                  string  `json:"name"`
	Price                 float64 `json:"price"`
	Description           string  `json:"description"`
	Categoria_idcategoria int64   `json:"categoria_idcategoria"`
}
