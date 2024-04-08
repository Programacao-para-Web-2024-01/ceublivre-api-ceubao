package produto

import (
	"errors"
)

type ProdutoRepository interface {
	List() ([]Produto, error)
	Get(id int) (*Produto, error)
	Create(produto Produto) (int64, error)
	Update(id int, produto Produto) error
	Delete(id int) error
}

type ProdutoService struct {
	produtoRepository ProdutoRepository
}

func NewProdutoService(produtoRepository ProdutoRepository) *ProdutoService {
	return &ProdutoService{produtoRepository}
}

func (ps *ProdutoService) List() ([]Produto, error) {
	return ps.produtoRepository.List()
}

func (ps *ProdutoService) Get(id int) (*Produto, error) {
	return ps.produtoRepository.Get(id)
}

func (ps *ProdutoService) Create(produto Produto) (int64, error) {
	if produto.Name == "" {
		return 0, errors.New("nome do produto é obrigatório")
	}
	return ps.produtoRepository.Create(produto)
}

func (ps *ProdutoService) Update(id int, produto Produto) error {
	if produto.Name == "" {
		return errors.New("nome do produto é obrigatório")
	}
	return ps.produtoRepository.Update(id, produto)
}

func (ps *ProdutoService) Delete(id int) error {
	return ps.produtoRepository.Delete(id)
}
