package clerk_test

import (
	"bytes"
	"errors"
	"flag"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/natural-affinity/decima/clerk"
	"github.com/natural-affinity/gotanda"
)

var update = flag.Bool("update", false, "update .golden files")

func TestTally(t *testing.T) {
	cases := []struct {
		Name    string
		Amounts []string
		Total   float64
		Err     error
	}{
		{"empty", []string{}, 0, nil},
		{"bad", []string{"a"}, 0, errors.New("a")},
		{"one.bad", []string{"100.25", "str"}, 0, errors.New("str")},
		{"multi.int", []string{"100", "25", "4", "-4", "0"}, 125, nil},
		{"multi.float", []string{"3500.31", "3498.69", "1.0001"}, 7000.0001, nil},
	}

	for _, tc := range cases {
		actualTotal, actualError := clerk.Tally(tc.Amounts)

		total := !(actualTotal == tc.Total)
		err := !gotanda.CompareError(actualError, tc.Err)

		if total || err {
			t.Errorf("\nTest: %s\n %s\nExpected:\n %f %s\nActual:\n %f %s",
				tc.Name, tc.Amounts,
				tc.Total, tc.Err,
				actualTotal, actualError)
		}
	}
}

func TestSubmitAndPrint(t *testing.T) {
	cases := []struct {
		Name     string
		Earnings float64
		Tithe    *clerk.Tithe
	}{
		{"print.zero", 0, &clerk.Tithe{Extra: 0, Percentage: 10}},
		{"print.extra", 1000, &clerk.Tithe{Extra: 100, Percentage: 10}},
		{"print.default", 1500, &clerk.Tithe{Extra: 0, Percentage: 10}},
		{"print.percent", 1500, &clerk.Tithe{Extra: 0, Percentage: 12.5}},
		{"print.percent.extra", 1250.28, &clerk.Tithe{Extra: 50.25, Percentage: 12.5}},
	}

	for _, tc := range cases {
		golden := filepath.Join("../testdata", tc.Name+".golden")

		// Submit and capture results
		tc.Tithe.Submit(func() float64 {
			return tc.Earnings
		})

		aout, _ := gotanda.Capture(func() {
			tc.Tithe.Print(true)
		})

		if *update {
			ioutil.WriteFile(golden, aout, 0644)
		}

		expected, _ := ioutil.ReadFile(golden)
		if !bytes.Equal(aout, expected) {
			t.Errorf("Test: %s\n Expected: %s\n Actual: %s\n", tc.Name, aout, expected)
		}
	}
}
