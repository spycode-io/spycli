package blueprint

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/gosimple/slug"
	"github.com/spycode-io/spycli.git/assets"
)

type Blueprint struct {
	Name          string
	SlugName      string
	BasePath      string
	Url           string
	Version       string
	BluePrintPath string
	Stack         string
	StackPath     string
	Regions       []string
}

type BlueprintInterface interface {
	InitBlueprint() error
}

var (
	BpFileTmplSet map[string]map[string][]assets.FileTmpl = map[string]map[string][]assets.FileTmpl{
		"default": {
			"blueprint": []assets.FileTmpl{
				{TmplFile: "gitignore.tmpl", File: ".gitignore"},
			},
			"stack": []assets.FileTmpl{
				{TmplFile: "prj.hcl.tmpl", File: "prj.hcl"},
			},
			"region": []assets.FileTmpl{
				{TmplFile: "region.hcl.tmpl", File: "region.hcl"},
				{TmplFile: "gitignore_region.tmpl", File: ".gitignore"},
			},
		},
	}
)

func NewBlueprint(
	templatesData embed.FS,
	basePath string,
	name string,
	url string,
	version string,
	stack string,
	regions []string) (*Blueprint, error) {

	blueprint := &Blueprint{
		Name:          name,
		SlugName:      slug.Make(name),
		BluePrintPath: fmt.Sprintf("%s/%s", basePath, slug.Make(name)),
		StackPath:     fmt.Sprintf("%s/%s/%s", basePath, slug.Make(name), slug.Make(stack)),
		Url:           url,
		Version:       version,
		BasePath:      basePath,
		Stack:         stack,
		Regions:       regions,
	}

	return blueprint, blueprint.InitBlueprint()
}

func (b *Blueprint) InitBlueprint() (err error) {

	log.Printf("Initializing blueprint %s %s", b.Url, b.Version)

	//Create base folder if necessary
	_, err = os.Stat(b.StackPath)
	if os.IsNotExist(err) {
		os.MkdirAll(b.StackPath, os.ModePerm)
		err = nil
	}

	//Create platform files from template by platform
	// for _, pf := range BpFileTmplSet["default"]["blueprint"] {
	// 	platformFile := fmt.Sprintf("%s/%s", b.BluePrintPath, pf.File)
	// 	err = b.WriteFile(pf.TmplFile, platformFile)
	// 	if nil != err {
	// 		return
	// 	}
	// }

	return
}
