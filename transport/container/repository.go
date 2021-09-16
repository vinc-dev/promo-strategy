package container

import "github.com/vinc-dev/promo-strategy/internal/domain/product"

// RepositoryContainer handle all repository used in project
type RepositoryContainer struct {
	Product product.RepositoryInterface
}

// CreateRepositoryContainer construct all repositories available in the app
func CreateRepositoryContainer() *RepositoryContainer {
	return &RepositoryContainer{
		Product: product.Repository(),
	}
}
