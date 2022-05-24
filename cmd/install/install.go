package install

import (
	"github.com/spf13/cobra"

	"github.com/mss-boot-io/mss-boot-generator/pkg"
)

var (
	StartCmd = &cobra.Command{
		Use:     "install",
		Short:   "Install mss-boot-generator",
		Example: "mss-boot-generator install",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	pkg.Install()
	return nil
}
