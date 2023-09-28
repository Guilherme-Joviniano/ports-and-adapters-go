package service

import "github.com/Guilherme-Joviniano/go-hexagonal/application/domain"

type ProductService struct {
	Persistence domain.ProductPersistenceInterface
}

func NewProductService(persistence domain.ProductPersistenceInterface) *ProductService {
	return &ProductService{
		Persistence: persistence,
	}
}

func (s *ProductService) Get(id string) (domain.ProductInterface, error) {
	product, err := s.Persistence.Get(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) Create(name string, price float32) (domain.ProductInterface, error) {
	product := domain.NewProduct(name, price)

	_, err := product.IsValid()

	if err != nil {
		return nil, err
	}

	result, err := s.Persistence.Save(product)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *ProductService) Enable(product domain.ProductInterface) (domain.ProductInterface, error) {
	err := product.Enable()
	if err != nil {
		return nil, err
	}
	result, err := s.Persistence.Save(product)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *ProductService) Disable(product domain.ProductInterface) (domain.ProductInterface, error) {
	err := product.Disable()

	if err != nil {
		return nil, err
	}

	result, err := s.Persistence.Save(product)

	if err != nil {
		return nil, err
	}

	return result, nil
}
