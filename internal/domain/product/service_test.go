package product_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vinc-dev/promo-strategy/internal/domain/product"
	"github.com/vinc-dev/promo-strategy/mock"
)

func TestService(t *testing.T) {
	t.Run("Find", func(t *testing.T) {
		// SETUP
		var testcases = []struct {
			SKU          string
			Success      bool
			ErrorMessage string
			Repository   product.RepositoryInterface
		}{
			{"120P90", true, "", &mock.ProductRepositoryMock{
				FnFind: func(ctx context.Context, sku string) (*product.Model, error) {
					return &product.Model{
						SKU:      "120P90",
						Name:     "Name",
						Price:    1000,
						Quantity: 10,
					}, nil
				},
			}},
			{"non-existing-sku", false, fmt.Sprintf("invalid product sku (non-existing-sku)"), &mock.ProductRepositoryMock{
				FnFind: func(ctx context.Context, sku string) (*product.Model, error) {
					return nil, fmt.Errorf("invalid product sku (non-existing-sku)")
				},
			}},
			{"", false, fmt.Sprintf("invalid product sku ()"), &mock.ProductRepositoryMock{
				FnFind: func(ctx context.Context, sku string) (*product.Model, error) {
					return nil, fmt.Errorf("invalid product sku ()")
				},
			}},
		}

		// TESTING
		for _, testcase := range testcases {
			service := product.Service(testcase.Repository)
			result, err := service.Find(context.Background(), testcase.SKU)
			if testcase.Success {
				require.NotNil(t, result)
				require.NoError(t, err)
			} else {
				require.Nil(t, result)
				require.Error(t, err)
				require.Equal(t, testcase.ErrorMessage, err.Error())
			}
		}
	})

	t.Run("FindByName", func(t *testing.T) {
		// SETUP
		var testcases = []struct {
			ProductName  string
			Success      bool
			ErrorMessage string
			Repository   product.RepositoryInterface
		}{
			{"MacBook Pro", true, "", &mock.ProductRepositoryMock{
				FnFindByName: func(ctx context.Context, name string) (*product.Model, error) {
					return &product.Model{
						SKU:      "SKU",
						Name:     "MacBook pro",
						Price:    1000,
						Quantity: 1,
					}, nil
				},
			}},
			{"invalid name", false, fmt.Sprintf("invalid product name (invalid name)"), &mock.ProductRepositoryMock{
				FnFindByName: func(ctx context.Context, name string) (*product.Model, error) {
					return nil, fmt.Errorf("invalid product name (invalid name)")
				},
			}},
			{"", false, fmt.Sprintf("invalid product name ()"), &mock.ProductRepositoryMock{
				FnFindByName: func(ctx context.Context, name string) (*product.Model, error) {
					return nil, fmt.Errorf("invalid product name ()")
				},
			}},
		}

		// TESTING
		for _, testcase := range testcases {
			service := product.Service(testcase.Repository)
			result, err := service.FindByName(context.Background(), testcase.ProductName)
			if testcase.Success {
				require.NotNil(t, result)
				require.NoError(t, err)
			} else {
				require.Nil(t, result)
				require.Error(t, err)
				require.Equal(t, testcase.ErrorMessage, err.Error())
			}
		}
	})

	t.Run("RemoveQuantity", func(t *testing.T) {
		// SETUP
		var testcases = []struct {
			SKU            string
			RemoveQuantity int64
			Success        bool
			ErrorMessage   string
			Repository     product.RepositoryInterface
		}{
			{"120P90", 1, true, "", &mock.ProductRepositoryMock{
				FnRemoveQuantity: func(ctx context.Context, sku string, removeQuantity int64) error {
					return nil
				},
			}},
			{"234234", 3, false, fmt.Sprintf("not enough items"), &mock.ProductRepositoryMock{
				FnRemoveQuantity: func(ctx context.Context, sku string, removeQuantity int64) error {
					return fmt.Errorf("not enough items")
				},
			}},
			{"invalid", 1, false, fmt.Sprintf("invalid product sku (invalid)"), &mock.ProductRepositoryMock{
				FnRemoveQuantity: func(ctx context.Context, sku string, removeQuantity int64) error {
					return fmt.Errorf("invalid product sku (invalid)")
				},
			}},
			{"", 1, false, fmt.Sprintf("invalid product sku ()"), &mock.ProductRepositoryMock{
				FnRemoveQuantity: func(ctx context.Context, sku string, removeQuantity int64) error {
					return fmt.Errorf("invalid product sku ()")
				},
			}},
		}

		// TESTING
		for _, testcase := range testcases {
			service := product.Service(testcase.Repository)
			err := service.RemoveQuantity(context.Background(), testcase.SKU, testcase.RemoveQuantity)
			if testcase.Success {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.Equal(t, testcase.ErrorMessage, err.Error())
			}
		}
	})
}
