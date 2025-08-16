package main

import (
	"os"

	"github.com/paivagustavo/mockery/tools/cmd"
)

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
