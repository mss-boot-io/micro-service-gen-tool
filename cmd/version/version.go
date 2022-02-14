package version

import (
	"github.com/spf13/cobra"

	"github.com/lwnmengjing/micro-service-gen-tool/pkg"
	"github.com/lwnmengjing/micro-service-gen-tool/version"
)

var (
	StartCmd = &cobra.Command{
		Use:     "version",
		Short:   "Get version info",
		Example: "generate-tool version",
		PreRun: func(cmd *cobra.Command, args []string) {
			pkg.Upgrade()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	version.PrintVersion()
	return nil
}
