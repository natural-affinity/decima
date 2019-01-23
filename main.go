package main

import (
	"log"
	"os"

	docopt "github.com/docopt/docopt-go"
	"github.com/natural-affinity/decima/clerk"
)

// Version Identifier
const Version = "0.0.1"

// Usage message (docopt interface)
const Usage = `
  Decima (la dîme)
    Tithe calculator

  Usage:
    decima [--percent p] [--extra e] [--breakdown] <amount>...
    decima --help
    decima --version

  Options:
    -h, --help          display help information
    -v, --version       display version information
    -b, --breakdown     display detailed breakdown
    -p, --percent p     tithe percentage [default: 10]
    -e, --extra e    	  add extra amount post-tithe
`

// DecimaArgs structure
type DecimaArgs struct {
	Extra     float64  `docopt:"--extra"`
	Amount    []string `docopt:"<amount>"`
	Percent   float64  `docopt:"--percent"`
	Breakdown bool     `docopt:"--breakdown"`
}

func main() {
	log.SetFlags(log.Lshortfile)

	// parse usage string and collect args
	args, err := docopt.ParseArgs(Usage, os.Args[1:], Version)
	if err != nil {
		log.Fatalf("invalid usage string: %s", err.Error())
	}

	var dargs DecimaArgs
	if err := args.Bind(&dargs); err != nil {
		log.Fatalf("invalid bindings: %s", err.Error())
	}

	tithe := &clerk.Tithe{Amount: dargs.Amount, Percentage: dargs.Percent}
	receipt := clerk.Collect(tithe, dargs.Extra)
	receipt.Print(dargs.Breakdown)
}
