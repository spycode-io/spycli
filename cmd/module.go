package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spycode-io/spycli/module"
	"github.com/spycode-io/spycli/project"
)

var (
	Module         string
	Library        string
	LibraryVersion string
)

func init() {

	initCmd(newModuleCmd)

	newModuleCmd.Flags().StringVarP(&Module, "module", "m", "", "Module (ex: aws/vpc)")
	newModuleCmd.Flags().StringVarP(&Library, "library", "l", "", "Library (ex: git@github.com:spycode-io/tf-components.git")
	newModuleCmd.Flags().StringVarP(&LibraryVersion, "version", "v", "", "Library version (or tag if it's a git repository)")
	newModuleCmd.Flags().StringVarP(&Stack, "stack", "s", "", "Stack name")
	newProjectCmd.Flags().StringSliceVarP(&Regions, "region", "r", project.DefaultRegions, "Pass a list of regions")

	newModuleCmd.MarkFlagRequired("module")
	newModuleCmd.MarkFlagRequired("library")
	newModuleCmd.MarkFlagRequired("version")
	newModuleCmd.MarkFlagRequired("stack")

	moduleCmd.AddCommand(newModuleCmd)
	rootCmd.AddCommand(moduleCmd)
}

var moduleCmd = &cobra.Command{
	Use:   "module",
	Short: "Manipulate modules",
	Long: `Use module commands
new: creates a new module
install: install a module
`,
}

var newModuleCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new module",
	Long:  `Use module new`,
	Run: func(cmd *cobra.Command, args []string) {

		base := getScaffold("templates/tmpl")

		mdl, err := module.NewModule(
			base, Stack, Library, LibraryVersion, Regions)

		if nil != err {
			log.Fatal(err)
		}

		log.Println(fmt.Printf("%+v\n", mdl))
	},
}
