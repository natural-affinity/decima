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
    decima [--percent p] [--extra e] [--breakdown] <amounts>...
    decima --help
    decima --version

  Options:
    -h, --help          display help information
    -v, --version       display version information
    -b, --breakdown     display detailed breakdown
    -p, --percent p     tithe percentage [default: 10]
    -e, --extra e       add extra amount post-tithe [default: 0]
`

func main() {
	log.SetFlags(log.Lshortfile)

	// parse usage string and collect args
	args, err := docopt.ParseArgs(Usage, os.Args[1:], Version)
	if err != nil {
		log.Fatalf("invalid usage string: %s", err.Error())
	}

	// extract options and args
	extra, err := args.Float64("--extra")
	if err != nil {
		log.Fatalf("invalid extra amount: %s", err.Error())
	}

	percent, err := args.Float64("--percent")
	if err != nil {
		log.Fatalf("invalid percentage: %s", err.Error())
	}

	amount := args["<amounts>"].([]string)
	verbose := args["--breakdown"].(bool)

	tithe := &clerk.Tithe{Amount: amount, Percentage: percent}
	transaction := clerk.Collect(tithe, extra)
	transaction.Print(verbose)
}
