package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var genCommandsHelpDoc = &cobra.Command{
	Use:   "man",
	Short: "Generate markdown commands manual",
	Long:  `Generate markdown commands manual`,
	Run: func(cmd *cobra.Command, args []string) {
		err := doc.GenMarkdownTree(rootCmd, BasePath)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	genCommandsHelpDoc.PersistentFlags().StringVarP(&BasePath, "directory", "d", ".", "Base directory where the files will be writen")
	rootCmd.AddCommand(genCommandsHelpDoc)
}
