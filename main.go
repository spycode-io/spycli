package main

import (
	"embed"

	"github.com/spycode-io/spycli.git/cmd"
)

//go:embed assets/*
var assetData embed.FS

func main() {
	cmd.AssetData = assetData
	cmd.Execute()
}

// func NewCommand() *cobra.Command {
// 	var slice []string
// 	var arr []string
// 	c := &cobra.Command{
// 		Use: "cmd",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			for _, v := range slice {
// 				fmt.Println(v)
// 			}

// 			for _, v := range arr {
// 				fmt.Println(v)
// 			}
// 		},
// 	}

// 	c.Flags().StringSliceVarP(&slice, "slice", "s", []string{}, "")
// 	c.Flags().StringArrayVarP(&arr, "arr", "a", []string{}, "")

// 	return c
// }

// func main() {
// 	NewCommand().Execute()
// }
