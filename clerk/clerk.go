package clerk

import (
	"fmt"
	"strconv"
)

// Transaction with clerk
type Transaction struct {
	Extra   float64
	Tithe   *Tithe
	Receipt *Receipt
	Verbose bool
}

// Tithe breakdown
type Tithe struct {
	Percentage float64
	Amount     []string
	Value      float64
}

// Receipt with tithe and calcuations
type Receipt struct {
	Total   float64
	Savings float64
	Tribute float64
}

// Collect tithe and issue receipt
func Collect(t *Tithe, extra float64) *Transaction {
	r := &Receipt{}
	for _, amount := range t.Amount {
		f, _ := strconv.ParseFloat(amount, 64)
		r.Total += f
	}

	t.Value = (r.Total * t.Percentage / 100.0)
	r.Tribute = (t.Value + extra)
	r.Savings = (r.Total - r.Tribute)

	return &Transaction{Tithe: t, Receipt: r, Extra: extra}
}

// Print receipt with correct verbosity
func (tx *Transaction) Print(verbose bool) {
	r := tx.Receipt
	t := tx.Tithe
	if verbose {
		fmt.Printf("Savings: $%.2f -Tithe: $%.2f -Extra: $%.2f\n", r.Total, t.Value, tx.Extra)
		fmt.Printf("= $%.2f\n\n", r.Savings)

		fmt.Printf("Tribute: Tithe: $%.2f + Extra: $%.2f\n", t.Value, tx.Extra)
		fmt.Printf("= $%.2f", r.Tribute)
	} else {
		if tx.Extra == 0 {
			fmt.Printf("You owe $%.2f", r.Tribute)
		} else {
			fmt.Printf("You owe $%.2f (tithe: $%.2f, extra: $%.2f)", r.Tribute, t.Value, tx.Extra)
		}
	}
}
