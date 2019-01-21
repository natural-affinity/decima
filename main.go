package main

import (
	"log"
	"os"

	docopt "github.com/docopt/docopt-go"
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
    -e, --extra e    	add extra amount post-tithe
    -p, --percent p     tithe percentage [default: 10]
`

func main() {
	log.SetFlags(log.Lshortfile)

	// parse usage string and collect args
	_, err := docopt.ParseArgs(Usage, os.Args[1:], Version)
	if err != nil {
		log.Fatalf("invalid usage string: %s", err.Error())
	}

	//extract options and args
}
