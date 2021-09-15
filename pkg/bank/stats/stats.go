package stats

import "github.com/fm2901/bank/pkg/bank/types"

// Calculates avg sum of payments
func Avg(payments []types.Payment) types.Money {
	allSum := types.Money(0)
	
	if len(payments) < 1 {
		return allSum
	}

	for _, payment := range payments {
		allSum += payment.Amount
	}

	avgSum := allSum / types.Money(len(payments))

	return avgSum
}