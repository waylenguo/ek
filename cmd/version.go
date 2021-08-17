package cmd

import (
	"github.com/ek/pkg/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "查看版本信息",
	Run: func(cmd *cobra.Command, args []string) {
		version.PrintVersion()
	},
}

var printVersion bool

func init() {
	versionCmd.Flags().BoolVar(&printVersion, "version", false, "显示版本号")
	rootCmd.AddCommand(versionCmd)
}