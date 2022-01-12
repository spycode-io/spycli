package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spycode-io/spycli/module"
)

var (
	ModuleUrl    string
	LocalLibrary bool
)

func init() {

	initCmd(includeModuleCmd)

	includeModuleCmd.Flags().StringVarP(&ModuleUrl, "url", "u", "", "Module URL. Ex: git@github.com:terraform-aws-modules/terraform-aws-vpc.git or a local folder tf-componets/my-module (using with -l parameter)")
	includeModuleCmd.Flags().BoolVarP(&LocalLibrary, "local", "l", false, "Use a local modules library")

	includeModuleCmd.MarkFlagRequired("url")

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

Ex: To create a vpc module called my-vpc from official terraform aws modules
 
  spycli module include -n "My VPC" -u git@github.com:terraform-aws-modules/terraform-aws-vpc.git

or you can include a local module from a library in same path of the project:

  spycli module include -n "My VPC" -u tf-aws-components/vpc -l

`,
	Run: func(cmd *cobra.Command, args []string) {

		base := getScaffold("templates/mdl")
		_, err := module.NewModule(base, ModuleUrl, LocalLibrary)

		if nil != err {
			log.Fatal(err)
		}

		log.Println("Module created successfully!")
	},
}
