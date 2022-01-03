package cmd

import (
	"embed"

	"github.com/spf13/cobra"
	"github.com/spycode-io/spycli.git/project"
)

var (
	BaseDirectory, Kind, Name string
	AssetData                 embed.FS
)

func init() {
	newProjectCmd.Flags().StringVarP(&BaseDirectory, "directory", "d", ".", "Base projects directory to execute command")
	newProjectCmd.Flags().StringVarP(&Kind, "kind", "k", "aws", "Kind of project (aws|azure)")
	newProjectCmd.Flags().StringVarP(&Name, "name", "n", "New Project", "Name of project")

	projectCmd.AddCommand(newProjectCmd)
	rootCmd.AddCommand(projectCmd)
}

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manipulate iac projects",
	Long: `Use project commands
new: creates a new project
`,
}

var newProjectCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new project",
	Long:  `Use project new`,
	Run: func(cmd *cobra.Command, args []string) {
		project.New(AssetData, BaseDirectory, Kind, Name)
	},
}
