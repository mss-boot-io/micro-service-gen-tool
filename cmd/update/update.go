package update

import (
	"github.com/spf13/cobra"

	"github.com/mss-boot-io/mss-boot-generator/pkg"
)

var (
	StartCmd = &cobra.Command{
		Use:     "update",
		Short:   "Install mss-boot-generator",
		Example: "mss-boot-generator update",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	pkg.Upgrade(false)
	return nil
}
