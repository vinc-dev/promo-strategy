package promo

import (
	"context"
	"math"

	"github.com/vinc-dev/promo-strategy/internal/domain/product"
	"github.com/vinc-dev/promo-strategy/internal/domain/promo/strategy"
)

// ServiceInterface interface of product service
type ServiceInterface interface {
	Checkout(ctx context.Context, payload []string) (float64, error)
}

type service struct {
	productService product.ServiceInterface
	strategies     []strategy.StrategyInterface
}

// Service return concrete service of ServiceInterface
func Service(
	productService product.ServiceInterface,
	strategies []strategy.StrategyInterface,
) ServiceInterface {
	return &service{
		productService: productService,
		strategies:     strategies,
	}
}

// Checkout is a function that receive your checkout product and calculate the discount base on available promo(s)
func (s *service) Checkout(ctx context.Context, payload []string) (float64, error) {
	var amountToPay float64
	mapProductQuantity := make(map[string]int64)
	for _, val := range payload {
		p, err := s.productService.FindByName(ctx, val)
		if nil != err {
			return 0, err
		}
		mapProductQuantity[p.SKU] += 1
		amountToPay += p.Price
		if err := s.productService.RemoveQuantity(ctx, p.SKU, 1); nil != err {
			return 0, err
		}
	}
	for _, strats := range s.strategies {
		discountAmount, err := strats.CalculateDiscount(ctx, mapProductQuantity)
		if nil != err {
			return 0, err
		}
		amountToPay -= discountAmount
	}
	return math.Round(amountToPay*100) / 100, nil
}
