package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/spycode-io/spycli/blueprint"
)

var (
	BluePrintName string
)

func init() {

	initCmd(newBlueprintCmd)

	newBlueprintCmd.Flags().StringVarP(&Stack, "stack", "s", "", "Stack name")
	newBlueprintCmd.Flags().StringSliceVarP(&Regions, "region", "r", []string{}, "Pass a list of regions")

	newBlueprintCmd.MarkFlagRequired("name")
	newBlueprintCmd.MarkFlagRequired("stack")

	blueprintCmd.AddCommand(newBlueprintCmd)
	rootCmd.AddCommand(blueprintCmd)
}

var blueprintCmd = &cobra.Command{
	Use:   "blueprint",
	Short: "Manipulate iac blueprints",
	Long:  `Use blueprint new`,
}

var newBlueprintCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new blueprint",
	Long: `

Creates a new empty blueprint

Ex: To create a new blueprint called bp-aws-nearform with a empty stack called simple-web-app with a region us-east-1

  spycli blueprint new -n "BP AWS Nearform" -s simple-web-app -r us-east-1
  
`,
	Run: func(cmd *cobra.Command, args []string) {

		base := getScaffold("templates/bp")

		bp, err := blueprint.NewBlueprint(
			base,
			Blueprint,
			BlueprintVersion,
			Stack,
			Regions,
		)

		if nil != err {
			log.Fatal(err)
		}

		log.Println(fmt.Printf("%+v\n", bp))
	},
}
