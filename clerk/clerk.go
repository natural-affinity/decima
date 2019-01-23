package clerk

import (
	"fmt"
	"strconv"
)

// Tithe breakdown
type Tithe struct {
	Percentage float64
	Amount     []string
}

// Receipt with tithe and calcuations
type Receipt struct {
	Tithe   float64
	Extra   float64
	Total   float64
	Savings float64
	Tribute float64
}

// Collect tithe and issue receipt
func Collect(t *Tithe, extra float64) *Receipt {
	sum := 0.0
	for _, amount := range t.Amount {
		f, _ := strconv.ParseFloat(amount, 64)
		sum += f
	}

	receipt := &Receipt{
		Total: sum,
		Tithe: (sum * t.Percentage / 100.0),
		Extra: extra,
	}
	receipt.Savings = (sum - receipt.Tithe)
	receipt.Tribute = (receipt.Tithe + receipt.Extra)

	return receipt
}

// Print receipt with correct verbosity
func (r *Receipt) Print(verbose bool) {
	if verbose {
		fmt.Printf("Savings: $%.2f -Tithe: $%.2f -Extra: $%.2f\n", r.Total, r.Tithe, r.Extra)
		fmt.Printf("= $%.2f", r.Savings)
	} else {
		if r.Extra == 0 {
			fmt.Printf("You owe $%.2f", r.Tribute)
		} else {
			fmt.Printf("You owe $%.2f (tithe: $%.2f, extra: $%.2f)", r.Tribute, r.Tithe, r.Extra)
		}
	}
}
