package blueprint

import (
	"fmt"
	"log"
	"os"

	"github.com/spycode-io/spycli/assets"
	"github.com/spycode-io/spycli/model"
)

type BlueprintScaffold struct {
	*model.Scaffold
	BluePrintPath string
	Stack         string
	StackPath     string
	Regions       []model.Region
}

type Module struct {
	Name string
}

var (
	DefaultFileSet map[string][]assets.FileTmpl = map[string][]assets.FileTmpl{
		"blueprint": {
			{TmplFile: "gitignore.tmpl", File: ".gitignore"},
		},
		"stack": {},
		"region": {
			{TmplFile: "region.yml.tmpl", File: "region.yml"},
		},
	}
)

func NewBlueprint(
	base *model.Scaffold,
	stack string,
	regions []string) (*BlueprintScaffold, error) {

	blueprint := &BlueprintScaffold{
		Scaffold:      base,
		BluePrintPath: fmt.Sprintf("%s/%s", base.BasePath, base.SlugName),
		Stack:         stack,
		StackPath:     fmt.Sprintf("%s/%s/%s", base.BasePath, base.SlugName, stack),
		Regions:       []model.Region{{Region: "_any"}},
	}

	for _, reg := range regions {
		blueprint.Regions = append(blueprint.Regions,
			model.Region{Region: reg},
		)
	}

	blueprint.FileSet.WithMap("default", DefaultFileSet)

	return blueprint, blueprint.InitBlueprint()
}

func (b *BlueprintScaffold) InitBlueprint() (err error) {

	log.Printf("Initializing a new blueprint %s", b.SlugName)

	//Create base folder if necessary
	_, err = os.Stat(b.BluePrintPath)
	if os.IsNotExist(err) {
		os.MkdirAll(b.BluePrintPath, os.ModePerm)
		err = nil
	}

	//Write blueprint files
	err = b.FileSet.WriteObjToPath("default", "blueprint", b.BluePrintPath, b)
	if nil != err {
		return
	}

	//Write stack files
	err = b.FileSet.WriteObjToPath("default", "stack", b.StackPath, b)
	if nil != err {
		return
	}

	// //Create region files from template
	for _, reg := range b.Regions {
		regPath := fmt.Sprintf("%s/%s", b.StackPath, reg.Region)
		err = b.FileSet.WriteObjToPath("default", "region", regPath, reg)
		if nil != err {
			return
		}
	}

	return
}
