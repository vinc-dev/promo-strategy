package strategy

import (
	"context"
	"fmt"

	"github.com/vinc-dev/promo-strategy/internal/domain/product"
)

type discount struct {
	productSKU      string
	minimalQuantity int64
	discountType    string
	discountValue   float64
	productService  product.ServiceInterface
}

func Discount(
	productSKU string,
	minimalQuantity int64,
	discountType string,
	discountValue float64,
	productService product.ServiceInterface,
) StrategyInterface {
	return &discount{
		productSKU:      productSKU,
		minimalQuantity: minimalQuantity,
		discountType:    discountType,
		discountValue:   discountValue,
		productService:  productService,
	}
}

// CalculateDiscount calculate discount base on discount strategy that has fulfilled the right conditions
func (p *discount) CalculateDiscount(ctx context.Context, payload map[string]int64) (float64, error) {
	if payload[p.productSKU] < p.minimalQuantity {
		return 0, nil
	}
	productModel, err := p.productService.Find(ctx, p.productSKU)
	if nil != err {
		return 0, err
	}
	var discountAmount float64
	switch p.discountType {
	case DISCOUNT_TYPE_AMOUNT:
		discountAmount = p.discountValue * float64(payload[p.productSKU])
	case DISCOUNT_TYPE_PERCENTAGE:
		discountAmount = p.discountValue / 100 * productModel.Price * float64(payload[p.productSKU])
	default:
		return 0, fmt.Errorf("invalid discount type")
	}
	return discountAmount, nil
}
