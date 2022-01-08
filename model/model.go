package model

import (
	"fmt"

	"github.com/gosimple/slug"
	"github.com/spycode-io/spycli/assets"
)

/* Entities */
type Environment struct {
	Name             string
	Path             string
	BlueprintVersion string
	Library          string
	LibraryVersion   string
}

type Region struct {
	Region string
}

type Scaffold struct {
	Name             string
	SlugName         string
	BasePath         string
	ScaffoldBasePath string
	FileSet          assets.FileSet
}

/* Construtores */
func NewScaffold(name string, basePath string, assetsPath string) *Scaffold {
	return &Scaffold{
		Name:             name,
		SlugName:         slug.Make(name),
		BasePath:         basePath,
		ScaffoldBasePath: fmt.Sprintf("%s/%s", basePath, slug.Make(name)),
		FileSet:          *assets.NewFileSet(assetsPath),
	}
}

/* Interfaces */
type ScaffoldInterface interface {
	Init() error
	WriteFile(tmplFile string, file string) (err error)
}
