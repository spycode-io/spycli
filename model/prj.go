package model

import "embed"

type Project struct {
	Name             string
	SlugName         string
	BasePath         string
	Platform         string
	PlatformBasePath string
	AssetsBasePath   string
	ProjectPath      string
	Stack            string
	Blueprint        string
	BlueprintVersion string
	Environments     []Environment
	Regions          []Region
	AssetsData       embed.FS
}
