package version

import (
	"github.com/spf13/cobra"

	"github.com/mss-boot-io/micro-service-gen-tool/pkg"
	"github.com/mss-boot-io/micro-service-gen-tool/version"
)

var (
	StartCmd = &cobra.Command{
		Use:     "version",
		Short:   "Get version info",
		Example: "generate-tool version",
		PreRun: func(cmd *cobra.Command, args []string) {
			pkg.Upgrade(true)
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
