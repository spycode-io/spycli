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

	initCmd(includeModuleCmd)

	includeModuleCmd.Flags().StringVarP(&Module, "module", "m", "", "Module (ex: aws/vpc)")
	includeModuleCmd.MarkFlagRequired("module")

	moduleCmd.AddCommand(includeModuleCmd)
	rootCmd.AddCommand(moduleCmd)
}

var moduleCmd = &cobra.Command{
	Use:   "module",
	Short: "Manipulate modules",
	Long:  `Use module commands`,
}

var includeModuleCmd = &cobra.Command{
	Use:   "include",
	Short: "Include module",
	Long: `include: includes a new module in a blueprint region

Ex: To create a vpc module called web-app-vpc
spycli module new -m vpc -n "VPC Web App"

`,
	Run: func(cmd *cobra.Command, args []string) {

		base := getScaffold("templates/mdl")

		mdl, err := module.NewModule(base, Module)

		if nil != err {
			log.Fatal(err)
		}

		log.Println(fmt.Printf("%+v\n", mdl))
	},
}
