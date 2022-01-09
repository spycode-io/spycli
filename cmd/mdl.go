package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spycode-io/spycli/module"
)

var (
	Module         string
	Library        string
	LibraryVersion string
)

func init() {

	initCmd(newModuleCmd)

	newModuleCmd.Flags().StringVarP(&Module, "module", "m", "", "Module (ex: aws/vpc)")
	newModuleCmd.MarkFlagRequired("module")

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

		base := getScaffold("templates/mdl")

		mdl, err := module.NewModule(base, Module)

		if nil != err {
			log.Fatal(err)
		}

		log.Println(fmt.Printf("%+v\n", mdl))
	},
}
