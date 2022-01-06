package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spycode-io/spycli/assets"
	"github.com/spycode-io/spycli/project"
)

var (
	BaseDirectory, Platform, ProjectName, Stack, Blueprint, BlueprintVersion string
	Environments, Regions                                                    []string
)

func init() {

	newProjectCmd.Flags().StringVarP(&BaseDirectory, "directory", "d", ".", "Base projects directory to execute command")
	newProjectCmd.Flags().StringVarP(&Platform, "kind", "k", "aws", "Kind of project (aws|azure)")
	newProjectCmd.Flags().StringVarP(&ProjectName, "name", "n", "New Project", "Name of project")
	newProjectCmd.Flags().StringVarP(&Stack, "stack", "s", "", "Stack name")
	newProjectCmd.Flags().StringVarP(&Blueprint, "blueprint", "b", "", "Blueprint")
	newProjectCmd.Flags().StringVarP(&BlueprintVersion, "blueprint-version", "v", "", "Blueprint version")
	newProjectCmd.Flags().StringSliceVarP(&Environments, "environment", "e", project.DefaultEnvironments, "Pass a list of environments")
	newProjectCmd.Flags().StringSliceVarP(&Regions, "region", "r", project.DefaultRegions, "Pass a list of regions")

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
		prj, err := project.New(
			assets.TemplatesData,
			BaseDirectory,
			Platform,
			ProjectName,
			Stack,
			Blueprint,
			BlueprintVersion,
			Environments,
			Regions)

		if nil != err {
			log.Fatal(err)
		}

		log.Println(fmt.Printf("%+v\n", prj))
	},
}
