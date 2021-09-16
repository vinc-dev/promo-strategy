package container

import (
	"github.com/vinc-dev/promo-strategy/internal/domain/product"
	"github.com/vinc-dev/promo-strategy/internal/domain/promo"
	"github.com/vinc-dev/promo-strategy/internal/domain/promo/strategy"
)

// ServiceContainer handle all service used in project
type ServiceContainer struct {
	Product product.ServiceInterface
	Promo   promo.ServiceInterface
}

// CreateServiceContainer construct all services available in the app
func CreateServiceContainer(repositories *RepositoryContainer) *ServiceContainer {
	productService := product.Service(repositories.Product)
	promoService := promo.Service(
		productService,
		[]strategy.StrategyInterface{
			strategy.PurchaseWithBonus("43N23P", 1, "234234", 1, productService),
			strategy.PurchaseWithBonus("120P90", 2, "120P90", 1, productService),
			strategy.Discount("A304SD", 3, strategy.DISCOUNT_TYPE_PERCENTAGE, 10, productService),
		},
	)
	return &ServiceContainer{
		Product: productService,
		Promo:   promoService,
	}
}
