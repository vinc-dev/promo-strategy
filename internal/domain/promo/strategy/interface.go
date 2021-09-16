package strategy

import "context"

type StrategyInterface interface {
	CalculateDiscount(ctx context.Context, payload map[string]int64) (float64, error)
}
