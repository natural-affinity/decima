package clerk

import (
	"errors"
	"fmt"
	"strconv"
)

// Tithe breakdown
type Tithe struct {
	Percentage float64
	Earnings   float64
	Amount     float64
	Extra      float64
}

// Receipt with breakdown
type Receipt struct {
	*Tithe
	Takehome float64
	Tribute  float64
}

// Submit and compute tribute amount
func (t *Tithe) Submit() *Receipt {
	t.Amount = (t.Earnings * t.Percentage / 100.0)

	tribute := (t.Amount + t.Extra)
	return &Receipt{
		Tithe:    t,
		Tribute:  tribute,
		Takehome: (t.Earnings - tribute),
	}
}

// Print receipt
func (r *Receipt) Print(verbose bool) {
	if verbose {
		fmt.Printf("Savings: $%.2f -Tithe: $%.2f -Extra: $%.2f\n", r.Earnings, r.Amount, r.Extra)
		fmt.Printf("= $%.2f\n\n", r.Takehome)

		fmt.Printf("Tribute: Tithe: $%.2f + Extra: $%.2f\n", r.Amount, r.Extra)
		fmt.Printf("= $%.2f", r.Tribute)
	} else {
		if r.Extra == 0 {
			fmt.Printf("You owe $%.2f", r.Tribute)
		} else {
			fmt.Printf("You owe $%.2f (tithe: $%.2f, extra: $%.2f)", r.Tribute, r.Amount, r.Extra)
		}
	}
}

// Collect(Percent, Earnings)
// (t *tithe) Submit
// (r *receipt) Print

// VerifyEarnings and sum
func VerifyEarnings(amounts []string) (float64, error) {
	earnings := 0.0
	for _, a := range amounts {
		value, err := strconv.ParseFloat(a, 64)
		if err != nil {
			msg := fmt.Sprintf("invalid amount: %s", a)
			return 0.0, errors.New(msg)
		}

		earnings += value
	}

	return earnings, nil
}
