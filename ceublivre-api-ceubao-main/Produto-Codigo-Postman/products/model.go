package products

type Product struct {
	ID                    int     `json:"id"`
	Name                  string  `json:"name"`
	Description           string  `json:"description"`
	Price                 float64 `json:"price"`
	Categoria_idcategoria int     `json:"categoria_idcategoria"`
}
