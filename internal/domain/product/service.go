package product

import "context"

// ServiceInterface interface of product service
type ServiceInterface interface {
	Find(ctx context.Context, sku string) (*Model, error)
	FindByName(ctx context.Context, name string) (*Model, error)
	RemoveQuantity(ctx context.Context, sku string, removeQuantity int64) error
}

type service struct {
	repository RepositoryInterface
}

// Service return concrete service of ServiceInterface
func Service(productRepository RepositoryInterface) ServiceInterface {
	return &service{
		repository: productRepository,
	}
}

// Find call product repository and find the product by given sku
func (s *service) Find(ctx context.Context, sku string) (*Model, error) {
	return s.repository.Find(ctx, sku)
}

// FindByName call product repository and find the product by given name
func (s *service) FindByName(ctx context.Context, name string) (*Model, error) {
	return s.repository.FindByName(ctx, name)
}

// RemoveQuantity call product repository and remove the quantity of given sku
func (s *service) RemoveQuantity(ctx context.Context, sku string, removeQuantity int64) error {
	return s.repository.RemoveQuantity(ctx, sku, removeQuantity)
}
