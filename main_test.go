package main_test

import (
	"flag"
	"testing"

	"github.com/natural-affinity/gotanda"
)

var update = flag.Bool("update", false, "update .golden files")

func TestUsage(t *testing.T) {
	cases := []struct {
		Name string
	}{
		{"help.long"},
		{"help.short"},
		{"version.long"},
		{"version.short"},
		{"invalid.earnings"},
		{"invalid.extra"},
		{"invalid.percentage"},
		{"earnings"},
		{"earnings.multi"},
		{"earnings.multi.extra"},
		{"earnings.extra.long"},
		{"earnings.breakdown.long"},
		{"earnings.breakdown.short"},
		{"earnings.breakdown.extra"},
		{"earnings.percentage.long"},
		{"earnings.percentage.breakdown"},
		{"earnings.percentage.extra"},
		{"earnings.percentage.extra.breakdown"},
	}

	for _, tc := range cases {
		r := gotanda.CompareCommand(t, tc, update)
		r.Assert(t, tc)
	}
}
