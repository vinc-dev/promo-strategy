package product

import (
	"context"
	"fmt"
)

// RepositoryInterface interface for product repository
type RepositoryInterface interface {
	Find(ctx context.Context, sku string) (*Model, error)
	FindByName(ctx context.Context, name string) (*Model, error)
	RemoveQuantity(ctx context.Context, sku string, removeQuantity int64) error
}

type repository struct {
	data       map[string]*Model
	dataByName map[string]*Model
}

// Repository return concrete instance of RepositoryInterface
func Repository() RepositoryInterface {
	return &repository{
		data:       data,
		dataByName: dataByName,
	}
}

// Find get product by sku
func (r *repository) Find(_ context.Context, sku string) (*Model, error) {
	m := r.data[sku]
	if m == nil {
		return nil, fmt.Errorf("product sku (%s) not found", sku)
	}
	return m, nil
}

// FindByName get product by name
func (r *repository) FindByName(_ context.Context, name string) (*Model, error) {
	m := r.dataByName[name]
	if m == nil {
		return nil, fmt.Errorf("product name (%s) not found", name)
	}
	return m, nil
}

// RemoveQuantity get given product by sku and remove the quantity of product
func (r *repository) RemoveQuantity(ctx context.Context, sku string, removeQuantity int64) error {
	result, err := r.Find(ctx, sku)
	if nil != err {
		return err
	}
	if result.Quantity < removeQuantity {
		return fmt.Errorf("item (SKU: %s) is out of stocks", result.SKU)
	}
	data[result.SKU].Quantity -= removeQuantity
	dataByName[result.Name].Quantity -= removeQuantity
	return nil
}
