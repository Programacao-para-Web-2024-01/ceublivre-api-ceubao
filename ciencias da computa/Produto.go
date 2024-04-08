package main

import (
	"errors"
	"sync"
)

type Produtos struct {
	Id        int64   `json:"id"`
	Name      string  `json:"name"`
	Preco     float32 `json:"preco"`
	Descricao string  `json:"descricao"`
	Categoria string  `json:"categoria"`
}

type ProdutoRepository struct {
	m  map[int]Produtos
	mu *sync.RWMutex
}

func NewProdutoRepository() *ProdutoRepository {
	return &ProdutoRepository{
		m:  make(map[int]Produtos),
		mu: &sync.RWMutex{},
	}
}

func (sr *ProdutoRepository) List() ([]Produtos, error) {
	sr.mu.RLock()
	defer sr.mu.RUnlock()
	Produtoss := make([]Produtos, len(sr.m))
	for id, Produtos := range sr.m {
		Produtoss[id-1] = Produtos
	}
	return Produtoss, nil
}

func (sr *ProdutoRepository) Get(id int) (*Produtos, error) {
	sr.mu.RLock()
	defer sr.mu.RUnlock()

	Produtos, ok := sr.m[id]
	if !ok {
		return nil, errors.New("Produtos not found")
	}

	return &Produtos, nil
}

func (sr *ProdutoRepository) Create(Produtos Produtos) (int, error) {
	sr.mu.Lock()
	defer sr.mu.Unlock()

	Produtos.Id = len(sr.m) + 1
	sr.m[Produtos.Id] = Produtos

	return Produtos.Id, nil
}

func (sr *ProdutoRepository) Update(id int, Produtos Produtos) error {
	sr.mu.Lock()
	defer sr.mu.Unlock()

	sr.m[id] = Produtos

	return nil
}

func (sr *ProdutoRepository) Delete(id int) error {
	sr.mu.Lock()
	defer sr.mu.Unlock()

	delete(sr.m, id)

	return nil
}
