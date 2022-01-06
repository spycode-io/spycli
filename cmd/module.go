package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spycode-io/spycli/module"
)

var (
	Module         string
	Name           string
	Library        string
	LibraryVersion string
)

func init() {

	newModuleCmd.Flags().StringVarP(&Name, "name", "n", "", "Name (ex: My VPC)")
	newModuleCmd.Flags().StringVarP(&BaseDirectory, "directory", "d", "", "Module (ex: aws/vpc)")
	newModuleCmd.Flags().StringVarP(&Module, "module", "m", "", "Module (ex: aws/vpc)")
	newModuleCmd.Flags().StringVarP(&Library, "library", "l", "", "Library (ex: git@github.com:spycode-io/tf-components.git")
	newModuleCmd.Flags().StringVarP(&LibraryVersion, "version", "v", "", "Library version (or tag if it's a git repository)")
	newModuleCmd.Flags().StringVarP(&Stack, "stack", "s", "", "Stack name")

	newModuleCmd.MarkFlagRequired("module")
	newModuleCmd.MarkFlagRequired("name")
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

		mdl, err := module.NewModule(Name, BaseDirectory, Library, LibraryVersion, Stack, Module)

		if nil != err {
			log.Fatal(err)
		}

		log.Println(fmt.Printf("%+v\n", mdl))
	},
}
