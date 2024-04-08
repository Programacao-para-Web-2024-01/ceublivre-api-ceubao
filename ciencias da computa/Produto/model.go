package Produto

type Produtos struct {
	Id        int64   `json:"id"`
	Name      string  `json:"name"`
	Preco     float32 `json:"preco"`
	Descricao string  `json:"descricao"`
	Categoria string  `json:"categoria"`
}

type Variacao_produtos struct {
	Id         int64  `json:"id"`
	Produto_id int64  `json:"produto_id"`
	Tamanho    string `json:"tamanho"`
	Cor        string `json:"cor"`
	Modelo     string `json:"modelo"`
}
