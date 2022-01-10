package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spycode-io/spycli/model"
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
												
SpyCLI is a CLI library for work with iac
projects, blueprints, tf-components, etc`,
}

var (
	Verbose, CleanCache bool
	Name, BasePath      string
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "V", false, "verbose output")
	log.SetOutput(os.Stdout)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func initCmd(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&Name, "name", "n", "", "Element name (ex: my-project or my-blueprint)")
	cmd.PersistentFlags().StringVarP(&BasePath, "directory", "d", "", "Base directory where the files will be writen")
}

func getScaffold(assetsPath string) *model.Scaffold {
	return model.NewScaffold(Name, BasePath, assetsPath)
}
