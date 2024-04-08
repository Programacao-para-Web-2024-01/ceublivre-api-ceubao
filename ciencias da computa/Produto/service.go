package Produto

type ProdutoService struct {
	repo *Produto_Repository
}

func NewProdutoService(repo *Produto_Repository) *ProdutoService {
	return &ProdutoService{repo: repo}
}

func (s *ProdutoService) List() ([]Student, error) {
	return s.repo.List()
}

func (s *ProdutoService) Get(id int) (*Student, error) {
	return s.repo.Get(id)
}

func (s *ProdutoService) Create(student Student) (*Student, error) {
	id, err := s.repo.Create(student)
	if err != nil {
		return nil, err
	}

	student.Id = id

	return &student, nil
}

func (s *ProdutoService) Update(student Student) error {
	_, err := s.Get(student.Id)
	if err != nil {
		return err
	}

	return s.repo.Update(student.Id, student)
}

func (s *ProdutoService) Delete(id int) error {
	_, err := s.Get(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
