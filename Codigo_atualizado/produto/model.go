package produto

type Produto struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Descricao string  `json:"descricao"`
	Preco     float64 `json:"preco"`
	Categoria string  `json:"categoria"`
}
