
package db

import (
	"errors"
	"sync"
)

type Produto struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Descricao string  `json:"descricao"`
	Preco     float64 `json:"preco"`
	Categoria string  `json:"categoria"`
}

type ProdutoRepository struct {
	m  map[int]Produto
	mu *sync.RWMutex
}

func NewProdutoRepository() *ProdutoRepository {
	return &ProdutoRepository{
		m:  make(map[int]Produto),
		mu: &sync.RWMutex{},
	}
}

func (sr *ProdutoRepository) List() ([]Produto, error) {
	sr.mu.RLock()
	defer sr.mu.RUnlock()
	produto := make([]Produto, len(sr.m))
	for id, p := range sr.m {
		produto[id-1] = p
	}

	return produto, nil
}

func (sr *ProdutoRepository) Get(id int) (*Produto, error) {
	sr.mu.RLock()
	defer sr.mu.RUnlock()

	produto, ok := sr.m[id]
	if !ok {
		return nil, errors.New("produto not found")
	}

	return &produto, nil
}

func (sr *ProdutoRepository) Create(produto Produto) (int, error) {
	sr.mu.Lock()
	defer sr.mu.Unlock()

	produto.Id = len(sr.m) + 1
	sr.m[produto.Id] = produto

	return produto.Id, nil
}

func (sr *ProdutoRepository) Update(id int, produto Produto) error {
	sr.mu.Lock()
	defer sr.mu.Unlock()

	sr.m[id] = produto

	return nil
}

func (sr *ProdutoRepository) Delete(id int) error {
	sr.mu.Lock()
	defer sr.mu.Unlock()

	delete(sr.m, id)

	return nil
}
