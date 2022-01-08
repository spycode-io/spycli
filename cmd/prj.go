package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spycode-io/spycli/project"
)

var (
	Platform, ProjectName, Stack, Blueprint, BlueprintVersion string
	Environments, Regions                                     []string
)

func init() {

	initCmd(newProjectCmd)

	newProjectCmd.Flags().StringVarP(&Platform, "platform", "p", "aws", "Plataform or service (aws|azure)")
	newProjectCmd.Flags().StringVarP(&Blueprint, "blueprint", "b", "", "Blueprint")
	newProjectCmd.Flags().StringVarP(&BlueprintVersion, "blueprint-version", "v", "", "Blueprint version")
	newProjectCmd.Flags().StringVarP(&Stack, "stack", "s", "", "Stack name")
	newProjectCmd.Flags().StringVarP(&Library, "library", "l", "", "Library (ex: git@github.com:spycode-io/tf-components.git")
	newProjectCmd.Flags().StringVarP(&LibraryVersion, "version", "k", "", "Library version (or tag if it's a git repository)")
	newProjectCmd.Flags().StringSliceVarP(&Regions, "region", "r", project.DefaultRegions, "Pass a list of environments")
	newProjectCmd.Flags().StringSliceVarP(&Environments, "environment", "e", project.DefaultEnvironments, "Pass a list of environments")

	newBlueprintCmd.MarkFlagRequired("blueprint")

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

		base := getScaffold("templates/prj")

		prj, err := project.NewProject(
			base,
			Platform,
			Stack,
			Blueprint,
			BlueprintVersion,
			Library,
			LibraryVersion,
			Environments,
			Regions)

		if nil != err {
			log.Fatal(err)
		}

		log.Println(fmt.Printf("%+v\n", prj))
	},
}
