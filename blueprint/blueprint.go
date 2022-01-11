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
	Url           string
	Version       string
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
	url string,
	version string,
	stack string,
	regions []string) (*BlueprintScaffold, error) {

	blueprint := &BlueprintScaffold{
		Scaffold:      base,
		Url:           url,
		Version:       version,
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

	log.Printf("Initializing blueprint %s %s", b.Url, b.Version)

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

	//Create base folder if necessary
	// _, err = os.Stat(b.StackPath)
	// if os.IsNotExist(err) {
	// 	os.MkdirAll(b.StackPath, os.ModePerm)
	// 	err = nil
	// }

	// //Create blueprint files from template
	// for _, pf := range BpFileSet.Set["default"]["blueprint"] {
	// 	bpFile := fmt.Sprintf("%s/%s", b.BluePrintPath, pf.File)
	// 	err = b.WriteFile(pf.TmplFile, bpFile)
	// 	if nil != err {
	// 		return
	// 	}
	// }

	// //Create stack files from template
	// for _, pf := range BpFileSet.Set["default"]["stack"] {
	// 	bpFile := fmt.Sprintf("%s/%s", b.StackPath, pf.File)
	// 	err = b.WriteFile(pf.TmplFile, bpFile)
	// 	if nil != err {
	// 		return
	// 	}
	// }

	// //Create region files from template
	// for _, r := range b.Regions {
	// 	for _, f := range BpFileSet.Set["default"]["region"] {
	// 		bpFile := fmt.Sprintf("%s/%s/%s", b.StackPath, r.Region, f.File)
	// 		err = BpFileSet.WriteObjToFile(f.TmplFile, bpFile, r)
	// 		if nil != err {
	// 			return
	// 		}
	// 	}
	// }

	return
}

// func (p *BlueprintScaffold) WriteFile(tmplFile string, file string) (err error) {
// 	return BpFileSet.WriteObjToFile(tmplFile, file, p)
// }
