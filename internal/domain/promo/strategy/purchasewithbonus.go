package strategy

import (
	"context"
	"fmt"

	"github.com/vinc-dev/promo-strategy/internal/domain/product"
)

type purchaseWithBonus struct {
	productSKU      string
	minimalQuantity int64
	bonusProductSKU string
	bonusQuantity   int64
	productService  product.ServiceInterface
}

func PurchaseWithBonus(
	productSKU string,
	minimalQuantity int64,
	bonusProductSKU string,
	bonusQuantity int64,
	productService product.ServiceInterface,
) StrategyInterface {
	return &purchaseWithBonus{
		productSKU:      productSKU,
		minimalQuantity: minimalQuantity,
		bonusProductSKU: bonusProductSKU,
		bonusQuantity:   bonusQuantity,
		productService:  productService,
	}
}

// CalculateDiscount calculate discount base on purchase with bonus strategy that has fulfilled the right conditions
func (p *purchaseWithBonus) CalculateDiscount(ctx context.Context, payload map[string]int64) (float64, error) {
	if payload[p.productSKU] < p.minimalQuantity {
		return 0, nil
	}
	productModel, err := p.productService.Find(ctx, p.bonusProductSKU)
	if nil != err {
		return 0, err
	}
	if productModel.Quantity < p.bonusQuantity && payload[p.bonusProductSKU] == 0 {
		return 0, nil
	}
	if payload[p.bonusProductSKU] < p.bonusQuantity {
		return 0, fmt.Errorf("you have unclaimed promotional product(s) (SKU: %s - %s, Quantity: %d), include your bonus before checking out", productModel.SKU, productModel.Name, p.bonusQuantity)
	}
	count := payload[p.productSKU] / p.minimalQuantity
	return float64(count) * productModel.Price, nil
}
