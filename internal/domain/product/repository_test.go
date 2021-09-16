package product_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vinc-dev/promo-strategy/internal/domain/product"
)

func TestRepository(t *testing.T) {
	repository := product.Repository()

	t.Run("Find", func(t *testing.T) {
		// SETUP
		var testcases = []struct {
			SKU          string
			Success      bool
			ErrorMessage string
		}{
			{"120P90", true, ""},
			{"non-existing-sku", false, "product sku (non-existing-sku) not found"},
			{"", false, "product sku () not found"},
		}

		// TESTING
		for _, testcase := range testcases {
			result, err := repository.Find(context.Background(), testcase.SKU)
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
		}{
			{"MacBook Pro", true, ""},
			{"invalid name", false, "product name (invalid name) not found"},
			{"", false, "product name () not found"},
		}

		// TESTING
		for _, testcase := range testcases {
			result, err := repository.FindByName(context.Background(), testcase.ProductName)
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
		}{
			{"120P90", 1, true, ""},
			{"234234", 3, false, "item (SKU: 234234) is out of stocks"},
			{"invalid", 1, false, "product sku (invalid) not found"},
			{"", 1, false, "product sku () not found"},
		}

		// TESTING
		for _, testcase := range testcases {
			err := repository.RemoveQuantity(context.Background(), testcase.SKU, testcase.RemoveQuantity)
			if testcase.Success {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.Equal(t, testcase.ErrorMessage, err.Error())
			}
		}
	})
}
