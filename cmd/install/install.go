package install

import (
	"github.com/spf13/cobra"

	"github.com/mss-boot-io/micro-service-gen-tool/pkg"
)

var (
	StartCmd = &cobra.Command{
		Use:     "install",
		Short:   "Install generate-tool",
		Example: "generate-tool install",
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
