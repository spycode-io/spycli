package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/spycode-io/spycli/blueprint"
	"github.com/spycode-io/spycli/project"
)

var (
	BluePrintName string
)

func init() {

	initCmd(newBlueprintCmd)

	newBlueprintCmd.Flags().StringVarP(&Blueprint, "blueprint", "b", "", "Blueprint")
	newBlueprintCmd.Flags().StringVarP(&BlueprintVersion, "version", "v", "v0.0.0", "Blueprint version")
	newBlueprintCmd.Flags().StringVarP(&Stack, "stack", "s", "", "Stack name")
	newBlueprintCmd.Flags().StringSliceVarP(&Regions, "region", "r", project.DefaultRegions, "Pass a list of regions")

	newBlueprintCmd.MarkFlagRequired("blueprint")
	newBlueprintCmd.MarkFlagRequired("stack")

	blueprintCmd.AddCommand(newBlueprintCmd)
	rootCmd.AddCommand(blueprintCmd)
}

var blueprintCmd = &cobra.Command{
	Use:   "blueprint",
	Short: "Manipulate iac blueprints",
	Long: `Use blueprint commands
new: creates a new blueprint
`,
}

var newBlueprintCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new project",
	Long:  `Use project new`,
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
