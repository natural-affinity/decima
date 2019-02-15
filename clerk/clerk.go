package clerk

import (
	"errors"
	"fmt"
	"strconv"
)

// Tithe breakdown
type Tithe struct {
	Extra      float64
	Percentage float64

	amount   float64
	tribute  float64
	earnings float64
	takehome float64
}

// Submission of tithe
type Submission interface {
	Submit(earnings func() float64)
	Print(verbose bool)
}

// Tally amounts
func Tally(amounts []string) (float64, error) {
	earnings := 0.0
	for _, amt := range amounts {
		value, err := strconv.ParseFloat(amt, 64)
		if err != nil {
			return 0, errors.New(amt)
		}

		earnings += value
	}

	return earnings, nil
}

// Submit tithe and compute totals
func (t *Tithe) Submit(earnings func() float64) {
	t.earnings = earnings()
	t.amount = (t.earnings * t.Percentage / 100.0)
	t.tribute = (t.amount + t.Extra)
	t.takehome = (t.earnings - t.tribute)
}

// Print tithe (optional: detailed breakdown)
func (t *Tithe) Print(verbose bool) {
	fmt.Printf("You owe: $%.2f\n\n", t.tribute)

	if verbose {
		fmt.Printf("Takehome: $%.2f\n", t.takehome)
		fmt.Printf("Earnings: $%.2f\n", t.earnings)
		fmt.Printf(" Percent: %.2f%%\n", t.Percentage)

		if t.Extra != 0 {
			fmt.Printf("   Tithe: $%.2f\n", t.amount)
			fmt.Printf("   Extra: $%.2f\n", t.Extra)
		}
	}
}
