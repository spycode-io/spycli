package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spycode-io/spycli/project"
)

var (
	Platform, ProjectName, Stack, Blueprint, BlueprintVersion, LibraryRelativePath string
	Environments, Regions                                                          []string
	LinkInit                                                                       bool
)

func init() {

	initCmd(newProjectCmd)

	newProjectCmd.Flags().StringVarP(&Platform, "platform", "p", "aws", "Plataform or service (aws|azure)")
	newProjectCmd.Flags().StringVarP(&Blueprint, "blueprint", "b", "", "Blueprint")
	newProjectCmd.Flags().StringVarP(&BlueprintVersion, "blueprint-version", "v", "", "Blueprint version")
	newProjectCmd.Flags().StringVarP(&Stack, "stack", "s", "", "Stack name")
	newProjectCmd.Flags().StringVarP(&Library, "library", "l", "", "Library (ex: git@github.com:spycode-io/tf-components.git")
	newProjectCmd.Flags().StringVarP(&LibraryVersion, "version", "k", "", "Library version (or tag if it's a git repository)")
	newProjectCmd.Flags().StringVarP(&LibraryRelativePath, "library-relative-path", "q", "", "Library version (or tag if it's a git repository)")
	newProjectCmd.Flags().StringSliceVarP(&Regions, "region", "r", project.DefaultRegions, "Pass a list of environments")
	newProjectCmd.Flags().StringSliceVarP(&Environments, "environment", "e", project.DefaultEnvironments, "Pass a list of environments")

	newProjectCmd.MarkFlagRequired("name")
	newBlueprintCmd.MarkFlagRequired("blueprint")

	initProjectCmd.PersistentFlags().StringVarP(&BasePath, "directory", "d", ".", "Base directory where the files will be writen")
	initProjectCmd.PersistentFlags().BoolVarP(&LinkInit, "link", "l", false, "Base directory where the files will be writen")
	initProjectCmd.MarkFlagRequired("directory")

	cleanProjectCmd.PersistentFlags().StringVarP(&BasePath, "directory", "d", ".", "Base directory where the files will be writen")
	cleanProjectCmd.MarkFlagRequired("directory")

	projectCmd.AddCommand(newProjectCmd)
	projectCmd.AddCommand(initProjectCmd)
	//projectCmd.AddCommand(cleanProjectCmd)

	rootCmd.AddCommand(projectCmd)
}

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manipulate iac projects",
	Long:  `Use project commands`,
}

var newProjectCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new project",
	Long: `Use project new
new: creates a new project with local reference for blueprint and components
Ex:

Create a local solution with all folders (project, blueprint and modules) in same folder

spycli project new -n "Prj Simple Web App" -b bp-aws-nearform -s simple-web-app -l tf-modules-aws -r us-east-1 -e dev -e prd -q "../../../../"

`,
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
			LibraryRelativePath,
			Environments,
			Regions)

		if nil != err {
			log.Fatal(err)
		}

		log.Println(fmt.Printf("%+v\n", prj))
	},
}

var initProjectCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a project",
	Long: `Use project init on a project folder

Ex:
spycli project init`,
	Run: func(cmd *cobra.Command, args []string) {

		err := project.InitProject(BasePath, LinkInit)

		if nil != err {
			log.Fatal(err)
		}
	},
}

var cleanProjectCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean a project",
	Long:  `Use project clean to remove all bp files`,
	Run: func(cmd *cobra.Command, args []string) {
		err := project.CleanStackFolder(BasePath, project.DefaultIgnoredFiles)
		if nil != err {
			log.Fatal(err)
		}
	},
}
