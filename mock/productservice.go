package mock

import (
	"context"

	"github.com/vinc-dev/promo-strategy/internal/domain/product"
)

type ProductServiceMock struct {
	FnFind           func(ctx context.Context, sku string) (*product.Model, error)
	FnFindByName     func(ctx context.Context, name string) (*product.Model, error)
	FnRemoveQuantity func(ctx context.Context, sku string, removeQuantity int64) error
}

func (m *ProductServiceMock) Find(ctx context.Context, sku string) (*product.Model, error) {
	return m.FnFind(ctx, sku)
}

func (m *ProductServiceMock) FindByName(ctx context.Context, name string) (*product.Model, error) {
	return m.FnFindByName(ctx, name)
}

func (m *ProductServiceMock) RemoveQuantity(ctx context.Context, sku string, removeQuantity int64) error {
	return m.FnRemoveQuantity(ctx, sku, removeQuantity)
}
