package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spycli",
	Short: "Spy iac CLI",
	Long: `
 ███████╗██████╗ ██╗   ██╗ ██████╗██╗     ██╗
 ██╔════╝██╔══██╗╚██╗ ██╔╝██╔════╝██║     ██║
 ███████╗██████╔╝ ╚████╔╝ ██║     ██║     ██║
 ╚════██║██╔═══╝   ╚██╔╝  ██║     ██║     ██║
 ███████║██║        ██║   ╚██████╗███████╗██║
 ╚══════╝╚═╝        ╚═╝    ╚═════╝╚══════╝╚═╝
												
Spy CLI is a CLI library for work with iac
projects, blueprints, tf-components, etc`,
}

var (
	Verbose, CleanCache bool
	BaseDirectory       string
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
