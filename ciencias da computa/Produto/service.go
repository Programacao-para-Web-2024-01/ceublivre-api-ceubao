package Produto

type ProdutoService struct {
	repo *ProdutoRepository
}

func NewProdutoService(repo *ProdutoRepository) *ProdutoService {
	return &ProdutoService{repo: repo}
}

func (s *ProdutoService) List() ([]Produto, error) {
	return s.repo.List()
}

func (s *ProdutoService) Get(id int) (*Produto, error) {
	return s.repo.Get(id)
}

func (s *ProdutoService) Create(produto Produto) (*Produto, error) {
	id, err := s.repo.Create(produto)
	if err != nil {
		return nil, err
	}

	Produto.Id = id

	return &Produto, nil
}

func (s *ProdutoService) Update(produto Produto) error {
	_, err := s.Get(produto.Id)
	if err != nil {
		return err
	}

	return s.repo.Update(produto.Id, produto)
}

func (s *ProdutoService) Delete(id int) error {
	_, err := s.Get(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
