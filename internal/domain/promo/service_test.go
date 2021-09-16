package promo_test

import (
	"context"
	"testing"

	"github.com/vinc-dev/promo-strategy/mock"

	"github.com/vinc-dev/promo-strategy/internal/domain/promo"

	"github.com/vinc-dev/promo-strategy/internal/domain/promo/strategy"

	"github.com/vinc-dev/promo-strategy/internal/domain/product"

	"github.com/stretchr/testify/require"
)

func TestRepository(t *testing.T) {
	t.Run("Checkout", func(t *testing.T) {
		// SETUP
		fnFindResult := []struct {
			Model *product.Model
			Err   error
		}{
			{
				Model: &product.Model{
					SKU:      "234234",
					Name:     "Raspberry Pi B",
					Price:    30.00,
					Quantity: 2,
				},
				Err: nil,
			},
			{
				Model: &product.Model{
					SKU:      "120P90",
					Name:     "Google Home",
					Price:    49.99,
					Quantity: 10,
				},
				Err: nil,
			},
			{
				Model: &product.Model{
					SKU:      "A304SD",
					Name:     "Alexa Speaker",
					Price:    109.50,
					Quantity: 10,
				},
				Err: nil,
			},
			{
				Model: &product.Model{
					SKU:      "NewSKU",
					Name:     "New SKU",
					Price:    10,
					Quantity: 10,
				},
				Err: nil,
			},
			{
				Model: &product.Model{
					SKU:      "234234",
					Name:     "Raspberry Pi B",
					Price:    30.00,
					Quantity: 2,
				},
				Err: nil,
			},
		}
		fnFindByNameResult := []struct {
			Model *product.Model
			Err   error
		}{
			{
				Model: &product.Model{
					SKU:      "43N23P",
					Name:     "MacBook Pro",
					Price:    5399.99,
					Quantity: 5,
				},
				Err: nil,
			},
			{
				Model: &product.Model{
					SKU:      "234234",
					Name:     "Raspberry Pi B",
					Price:    30.00,
					Quantity: 2,
				},
				Err: nil,
			},
			{
				Model: &product.Model{
					SKU:      "120P90",
					Name:     "Google Home",
					Price:    49.99,
					Quantity: 10,
				},
				Err: nil,
			},
			{
				Model: &product.Model{
					SKU:      "120P90",
					Name:     "Google Home",
					Price:    49.99,
					Quantity: 10,
				},
				Err: nil,
			},
			{
				Model: &product.Model{
					SKU:      "120P90",
					Name:     "Google Home",
					Price:    49.99,
					Quantity: 10,
				},
				Err: nil,
			},
			{
				Model: &product.Model{
					SKU:      "A304SD",
					Name:     "Alexa Speaker",
					Price:    109.50,
					Quantity: 10,
				},
				Err: nil,
			},
			{
				Model: &product.Model{
					SKU:      "A304SD",
					Name:     "Alexa Speaker",
					Price:    109.50,
					Quantity: 10,
				},
				Err: nil,
			},
			{
				Model: &product.Model{
					SKU:      "A304SD",
					Name:     "Alexa Speaker",
					Price:    109.50,
					Quantity: 10,
				},
				Err: nil,
			},
			{
				Model: &product.Model{
					SKU:      "NewSKU",
					Name:     "New SKU",
					Price:    10,
					Quantity: 10,
				},
				Err: nil,
			},
			{
				Model: &product.Model{
					SKU:      "NewSKU",
					Name:     "New SKU",
					Price:    10,
					Quantity: 10,
				},
				Err: nil,
			},
			{
				Model: &product.Model{
					SKU:      "NewSKU",
					Name:     "New SKU",
					Price:    10,
					Quantity: 10,
				},
				Err: nil,
			},
			{
				Model: &product.Model{
					SKU:      "43N23P",
					Name:     "MacBook Pro",
					Price:    5399.99,
					Quantity: 5,
				},
				Err: nil,
			},
		}
		strategies := []strategy.StrategyInterface{
			strategy.PurchaseWithBonus("43N23P", 1, "234234", 1, &mock.ProductRepositoryMock{
				FnFind: func(ctx context.Context, sku string) (*product.Model, error) {
					if len(fnFindResult) == 0 {
						return nil, nil
					}
					currResult := fnFindResult[0]
					fnFindResult = fnFindResult[1:]
					return currResult.Model, currResult.Err
				},
			}),
			strategy.PurchaseWithBonus("120P90", 2, "120P90", 1, &mock.ProductRepositoryMock{
				FnFind: func(ctx context.Context, sku string) (*product.Model, error) {
					if len(fnFindResult) == 0 {
						return nil, nil
					}
					currResult := fnFindResult[0]
					fnFindResult = fnFindResult[1:]
					return currResult.Model, currResult.Err
				},
			}),
			strategy.Discount("A304SD", 3, strategy.DISCOUNT_TYPE_PERCENTAGE, 10, &mock.ProductRepositoryMock{
				FnFind: func(ctx context.Context, sku string) (*product.Model, error) {
					if len(fnFindResult) == 0 {
						return nil, nil
					}
					currResult := fnFindResult[0]
					fnFindResult = fnFindResult[1:]
					return currResult.Model, currResult.Err
				},
			}),
			strategy.Discount("NewSKU", 3, "invalid type", 10, &mock.ProductRepositoryMock{
				FnFind: func(ctx context.Context, sku string) (*product.Model, error) {
					if len(fnFindResult) == 0 {
						return nil, nil
					}
					currResult := fnFindResult[0]
					fnFindResult = fnFindResult[1:]
					return currResult.Model, currResult.Err
				},
			}),
		}

		var testcases = []struct {
			Result         float64
			Success        bool
			ErrorMessage   string
			Payload        []string
			ProductService product.ServiceInterface
		}{
			{
				Result:       5399.99,
				Success:      true,
				ErrorMessage: "",
				Payload:      []string{"MacBook Pro", "Raspberry Pi B"},
				ProductService: &mock.ProductServiceMock{
					FnFindByName: func(ctx context.Context, name string) (*product.Model, error) {
						if len(fnFindByNameResult) == 0 {
							return nil, nil
						}
						currResult := fnFindByNameResult[0]
						fnFindByNameResult = fnFindByNameResult[1:]
						return currResult.Model, currResult.Err
					},
					FnRemoveQuantity: func(ctx context.Context, sku string, removeQuantity int64) error {
						return nil
					},
				},
			},
			{
				Result:       99.98,
				Success:      true,
				ErrorMessage: "",
				Payload:      []string{"Google Home", "Google Home", "Google Home"},
				ProductService: &mock.ProductServiceMock{
					FnFindByName: func(ctx context.Context, name string) (*product.Model, error) {
						if len(fnFindByNameResult) == 0 {
							return nil, nil
						}
						currResult := fnFindByNameResult[0]
						fnFindByNameResult = fnFindByNameResult[1:]
						return currResult.Model, currResult.Err
					},
					FnRemoveQuantity: func(ctx context.Context, sku string, removeQuantity int64) error {
						return nil
					},
				},
			},
			{
				Result:       295.65,
				Success:      true,
				ErrorMessage: "",
				Payload:      []string{"Alexa Speaker", "Alexa Speaker", "Alexa Speaker"},
				ProductService: &mock.ProductServiceMock{
					FnFindByName: func(ctx context.Context, name string) (*product.Model, error) {
						if len(fnFindByNameResult) == 0 {
							return nil, nil
						}
						currResult := fnFindByNameResult[0]
						fnFindByNameResult = fnFindByNameResult[1:]
						return currResult.Model, currResult.Err
					},
					FnRemoveQuantity: func(ctx context.Context, sku string, removeQuantity int64) error {
						return nil
					},
				},
			},
			{
				Result:       0,
				Success:      false,
				ErrorMessage: "invalid discount type",
				Payload:      []string{"New SKU", "New SKU", "New SKU"},
				ProductService: &mock.ProductServiceMock{
					FnFindByName: func(ctx context.Context, name string) (*product.Model, error) {
						if len(fnFindByNameResult) == 0 {
							return nil, nil
						}
						currResult := fnFindByNameResult[0]
						fnFindByNameResult = fnFindByNameResult[1:]
						return currResult.Model, currResult.Err
					},
					FnRemoveQuantity: func(ctx context.Context, sku string, removeQuantity int64) error {
						return nil
					},
				},
			},
			{
				Result:       0,
				Success:      false,
				ErrorMessage: "you have unclaimed promotional product(s) (SKU: 234234 - Raspberry Pi B, Quantity: 1), include your bonus before checking out",
				Payload:      []string{"MacBook Pro"},
				ProductService: &mock.ProductServiceMock{
					FnFindByName: func(ctx context.Context, name string) (*product.Model, error) {
						if len(fnFindByNameResult) == 0 {
							return nil, nil
						}
						currResult := fnFindByNameResult[0]
						fnFindByNameResult = fnFindByNameResult[1:]
						return currResult.Model, currResult.Err
					},
					FnRemoveQuantity: func(ctx context.Context, sku string, removeQuantity int64) error {
						return nil
					},
				},
			},
		}

		// TESTING
		for _, testcase := range testcases {
			service := promo.Service(testcase.ProductService, strategies)
			result, err := service.Checkout(context.Background(), testcase.Payload)
			if testcase.Success {
				require.Equal(t, testcase.Result, result)
				require.NoError(t, err)
			} else {
				require.Equal(t, testcase.Result, result)
				require.Error(t, err)
				require.Equal(t, testcase.ErrorMessage, err.Error())
			}
		}
	})
}
