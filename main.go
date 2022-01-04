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
