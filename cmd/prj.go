package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spycode-io/spycli/project"
)

var (
	Platform, ProjectName, Stack, Blueprint, RemoteStateBucket, RemoteStateRegion string
	Environments, Regions                                                         []string
	LinkInit, LocalBlueprint, UseRemoteState                                      bool
)

func init() {

	initCmd(newProjectCmd)

	newProjectCmd.Flags().StringVarP(&Platform, "platform", "p", "aws", "Plataform or service (aws|azure)")
	newProjectCmd.Flags().StringVarP(&Blueprint, "blueprint", "b", "", "Blueprint")
	newProjectCmd.Flags().StringVarP(&Stack, "stack", "s", "", "Stack name")

	newProjectCmd.Flags().StringSliceVarP(&Regions, "region", "r", project.DefaultRegions, "Pass a list of environments")
	newProjectCmd.Flags().StringSliceVarP(&Environments, "environment", "e", project.DefaultEnvironments, "Pass a list of environments")
	newProjectCmd.Flags().BoolVarP(&LocalBlueprint, "local", "l", false, "Local blueprint")

	newProjectCmd.Flags().BoolVarP(&UseRemoteState, "remote-state", "t", false, "Use remote state")
	newProjectCmd.Flags().StringVarP(&RemoteStateBucket, "remote-bucket", "u", "", "Stack name")
	newProjectCmd.Flags().StringVarP(&RemoteStateRegion, "remote-bucket-region", "v", "", "Stack name")

	newProjectCmd.MarkFlagRequired("name")
	newProjectCmd.MarkFlagRequired("blueprint")
	newProjectCmd.MarkFlagRequired("stack")

	initProjectCmd.Flags().BoolVarP(&LinkInit, "link", "l", false, "Link files locally instead of copy. This option is the best when editing blueprint files")
	initProjectCmd.Flags().StringVarP(&BasePath, "directory", "d", ".", "Base directory where the files will be writen")

	projectCmd.AddCommand(newProjectCmd)
	projectCmd.AddCommand(initProjectCmd)

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
	Long: `Creates a new project with local or remote reference for blueprint and components

Examples:

Create a project that:

- Is called "My Project"
- Uses blueprint bp-aws-nearform and stack simple-web-app locally
- Have two environments: develop and production
- Have two regions: us-east-1 and us-west-1

spycli project new -n "My Project" -b bp-aws-nearform -l -s simple-web-app -r us-east-1 -e develop -e production

The same project but using remove blueprint:

spycli project new -n "My Project" -b git@github.com:nearform/bp-aws-nearform.git -s simple-web-app -r us-east-1 -e develop -e production

The same project but using remote blueprint and remote state in terraform:

spycli project new -n "My Project" -b git@github.com:nearform/bp-aws-nearform.git -s simple-web-app -r us-east-1 -e develop -e production -t -u my-bucket -v us-east-1

`,
	Run: func(cmd *cobra.Command, args []string) {

		base := getScaffold("templates/prj")

		_, err := project.NewProject(
			base,
			Platform,
			Stack,
			Blueprint,
			UseRemoteState,
			RemoteStateBucket,
			RemoteStateRegion,
			Environments,
			Regions)

		if nil != err {
			log.Fatal(err)
		}

		log.Println("Project created successfully!")
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

		log.Println("Project initialized successfully!")
	},
}

// var cleanProjectCmd = &cobra.Command{
// 	Use:   "clean",
// 	Short: "Clean a project",
// 	Long:  `Use project clean to remove all bp files`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		err := project.CleanStackFolder(BasePath, project.DefaultIgnoredFiles)
// 		if nil != err {
// 			log.Fatal(err)
// 		}
// 	},
// }
