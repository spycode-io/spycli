package model

import (
	"embed"

	"github.com/gosimple/slug"
	"github.com/spycode-io/spycli/assets"
)

/* Entities */
type Environment struct {
	Name string
	Path string
}

type Region struct {
	Region string
}

type Scaffold struct {
	Name           string
	SlugName       string
	BasePath       string
	AssetsBasePath string
	AssetsData     embed.FS
	FileSet        assets.FileSet
}

type ProjectScaffold struct {
	Scaffold
	Platform         string
	ProjectPath      string
	Stack            string
	Blueprint        string
	BlueprintVersion string
	Environments     []Environment
	Regions          []Region
}

/* Construtores */
func NewScaffold(name string, basePath string, assetsPath string) *Scaffold {
	return &Scaffold{
		Name:           name,
		SlugName:       slug.Make(name),
		BasePath:       basePath,
		AssetsBasePath: assetsPath,
		FileSet:        *assets.NewFileSet(assetsPath),
	}
}

/* Interfaces */
type ScaffoldInterface interface {
	Init() error
	WriteFile(tmplFile string, file string) (err error)
}
