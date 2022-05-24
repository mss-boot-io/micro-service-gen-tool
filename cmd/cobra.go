package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/mss-boot-io/mss-boot-generator/cmd/install"
	"github.com/mss-boot-io/mss-boot-generator/cmd/run"
	"github.com/mss-boot-io/mss-boot-generator/cmd/update"
	"github.com/mss-boot-io/mss-boot-generator/cmd/version"
	"github.com/mss-boot-io/mss-boot-generator/pkg"
	v "github.com/mss-boot-io/mss-boot-generator/version"
)

var rootCmd = &cobra.Command{
	Use:          "mss-boot-generator",
	Short:        "mss-boot-generator",
	SilenceUsage: true,
	Long:         `mss-boot-generator`,
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
	usageStr := `欢迎使用 ` + pkg.Green(`mss-boot-generator `+v.Version) + ` 可以使用 ` + pkg.Red(`-h`) + ` 查看命令`
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
