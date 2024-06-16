package products

type ProductService struct {
	repo *ProductRepository
}

func NewProductService(repo *ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) List() ([]Product, error) {
	return s.repo.List()
}

func (s *ProductService) Get(id int) (*Product, error) {
	return s.repo.Get(id)
}

func (s *ProductService) Create(product Product) (*Product, error) {
	id, err := s.repo.Create(product)
	if err != nil {
		return nil, err
	}

	product.Id = id

	return &product, nil
}

func (s *ProductService) Update(product Product) error {
	id := int(product.Id)

	_, err := s.Get(id)
	if err != nil {
		return err
	}

	return s.repo.Update(id, product)
}

func (s *ProductService) Delete(id int) error {
	_, err := s.Get(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
