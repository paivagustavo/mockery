package cmd

import (
	"fmt"

	"github.com/paivagustavo/mockery/v3/internal/logging"
	"github.com/spf13/cobra"
)

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version of mockery",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(logging.GetSemverInfo())
		},
	}
}
