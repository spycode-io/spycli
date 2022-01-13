package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spycode-io/spycli/model"
)

var rootCmd = &cobra.Command{
	Use:   "spycli",
	Short: "spycli the killer iac project tool",
	Long: `

   ███████╗██████╗ ██╗   ██╗ ██████╗██╗     ██╗
   ██╔════╝██╔══██╗╚██╗ ██╔╝██╔════╝██║     ██║
   ███████╗██████╔╝ ╚████╔╝ ██║     ██║     ██║
   ╚════██║██╔═══╝   ╚██╔╝  ██║     ██║     ██║
   ███████║██║        ██║   ╚██████╗███████╗██║
   ╚══════╝╚═╝        ╚═╝    ╚═════╝╚══════╝╚═╝												
  v0.0.0
  
  SpyCLI is a command library tool for work with iac
  projects, blueprints, modules, etc`,
}

var (
	Verbose, CleanCache bool
	Name, BasePath      string
)

func init() {
	rootCmd.Flags().BoolVarP(&Verbose, "verbose", "V", false, "verbose output")
	log.SetOutput(os.Stdout)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func initCmd(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&Name, "name", "n", "", "Element name (ex: my-project or my-blueprint)")
	cmd.Flags().StringVarP(&BasePath, "directory", "d", ".", "Base directory where the files will be writen")
}

func getScaffold(assetsPath string) *model.Scaffold {
	return model.NewScaffold(Name, BasePath, assetsPath)
}
