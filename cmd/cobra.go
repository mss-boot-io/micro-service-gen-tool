package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lwnmengjing/micro-service-gen-tool/cmd/install"
	"github.com/lwnmengjing/micro-service-gen-tool/cmd/run"
	"github.com/lwnmengjing/micro-service-gen-tool/cmd/update"
	"github.com/lwnmengjing/micro-service-gen-tool/cmd/version"
	"github.com/lwnmengjing/micro-service-gen-tool/pkg"
	v "github.com/lwnmengjing/micro-service-gen-tool/version"
)

var rootCmd = &cobra.Command{
	Use:          "generate-tool",
	Short:        "generate-tool",
	SilenceUsage: true,
	Long:         `generate-tool`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New(pkg.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `欢迎使用 ` + pkg.Green(`generate-tool `+v.Version) + ` 可以使用 ` + pkg.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s\n", usageStr)
}

func init() {
	rootCmd.AddCommand(version.StartCmd)
	rootCmd.AddCommand(run.StartCmd)
	rootCmd.AddCommand(install.StartCmd)
	rootCmd.AddCommand(update.StartCmd)
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
