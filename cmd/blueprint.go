package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spycode-io/spycli.git/blueprint"
)

func init() {
	newBlueprintCmd.Flags().StringVarP(&BaseDirectory, "directory", "d", ".", "Base projects directory to execute command")
	newBlueprintCmd.Flags().StringVarP(&Blueprint, "blueprint", "b", "", "Blueprint")
	newBlueprintCmd.Flags().StringVarP(&BlueprintVersion, "blueprint-version", "v", "", "Blueprint version")

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

		bp, err := blueprint.NewBlueprint(AssetData, BaseDirectory, Blueprint, BlueprintVersion)

		if nil != err {
			log.Fatal(err)
		}

		log.Println(fmt.Printf("%+v\n", bp))
	},
}
