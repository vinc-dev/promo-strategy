package mock

import (
	"context"

	"github.com/vinc-dev/promo-strategy/internal/domain/product"
)

type ProductRepositoryMock struct {
	FnFind           func(ctx context.Context, sku string) (*product.Model, error)
	FnFindByName     func(ctx context.Context, name string) (*product.Model, error)
	FnRemoveQuantity func(ctx context.Context, sku string, removeQuantity int64) error
}

func (m *ProductRepositoryMock) Find(ctx context.Context, sku string) (*product.Model, error) {
	return m.FnFind(ctx, sku)
}

func (m *ProductRepositoryMock) FindByName(ctx context.Context, name string) (*product.Model, error) {
	return m.FnFindByName(ctx, name)
}

func (m *ProductRepositoryMock) RemoveQuantity(ctx context.Context, sku string, removeQuantity int64) error {
	return m.FnRemoveQuantity(ctx, sku, removeQuantity)
}
