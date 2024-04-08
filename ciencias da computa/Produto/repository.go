package Produto

type ProdutoRepository struct {
	repo *Produto_Repository
}

func NewProdutoRepository(repo *Produto_Repository) *ProdutoRepository {
	return &ProdutoRepository{repo: repo}
}

func (s *ProdutoRepository) List() ([]Produto, error) {
	return s.repo.List()
}

func (s *ProdutoRepository) Get(id int) (*Produto, error) {
	return s.repo.Get(id)
}

func (s *ProdutoRepository) Create(Produto Produto) (*Produto, error) {
	id, err := s.repo.Create(Produto)
	if err != nil {
		return nil, err
	}

	Produto.Id = id

	return &Produto, nil
}

func (s *ProdutoRepository) Update(Produto Produto) error {
	_, err := s.Get(Produto.Id)
	if err != nil {
		return err
	}

	return s.repo.Update(Produto.Id, Produto)
}

func (s *ProdutoRepository) Delete(id int) error {
	_, err := s.Get(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
