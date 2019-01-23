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
	Tribute    float64
	Takehome   float64
}

// UpdateEarnings inline using amounts
func (t *Tithe) UpdateEarnings(amounts []string) error {
	t.Earnings = 0.0
	for _, amt := range amounts {
		value, err := strconv.ParseFloat(amt, 64)
		if err != nil {
			return errors.New(amt)
		}

		t.Earnings += value
	}

	return nil
}

// Submit tithe and return takehome pay
func (t *Tithe) Submit(amounts []string) (float64, error) {
	if err := t.UpdateEarnings(amounts); err != nil {
		return 0, err
	}

	// calculate totals
	t.Amount = (t.Earnings * t.Percentage / 100.0)
	t.Tribute = (t.Amount + t.Extra)
	t.Takehome = (t.Earnings - t.Tribute)

	return t.Takehome, nil
}

// Print tithe (optional takehome)
func (t *Tithe) Print(verbose bool) {
	fmt.Printf("You owe: $%.2f\n\n", t.Tribute)

	if verbose {
		fmt.Printf("Takehome: $%.2f\n", t.Takehome)
		fmt.Printf("Earnings: $%.2f\n", t.Earnings)
		fmt.Printf(" Percent: %.2f%%\n", t.Percentage)

		if t.Extra != 0 {
			fmt.Printf("   Tithe: $%.2f\n", t.Amount)
			fmt.Printf("   Extra: $%.2f\n", t.Extra)
		}
	}
}
