package blueprint

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/gosimple/slug"
	"github.com/spycode-io/spycli/assets"
	"github.com/spycode-io/spycli/model"
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
	AssetsData    embed.FS
	Regions       []model.Region
}

type Module struct {
	Name string
}

type BlueprintInterface interface {
	InitBlueprint() error
	WriteFile(tmplFile string, file string) (err error)
}

var (
	BpFileSet assets.FileSet = assets.FileSet{
		AssetsPath: "templates/bp",
		Set: map[string]map[string][]assets.FileTmpl{
			"default": {
				"blueprint": []assets.FileTmpl{
					{TmplFile: "gitignore.tmpl", File: ".gitignore"},
				},
				"stack": []assets.FileTmpl{},
				"region": []assets.FileTmpl{
					{TmplFile: "region.hcl.tmpl", File: "region.hcl"},
				},
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
		BasePath:      basePath,
		Url:           url,
		Version:       version,
		BluePrintPath: fmt.Sprintf("%s/%s", basePath, slug.Make(name)),
		Stack:         stack,
		StackPath:     fmt.Sprintf("%s/%s/%s", basePath, slug.Make(name), slug.Make(stack)),
		AssetsData:    embed.FS{},
		Regions:       []model.Region{{Region: "any"}},
	}

	for _, reg := range regions {
		blueprint.Regions = append(blueprint.Regions,
			model.Region{Region: reg},
		)
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

	//Create blueprint files from template
	for _, pf := range BpFileSet.Set["default"]["blueprint"] {
		bpFile := fmt.Sprintf("%s/%s", b.BluePrintPath, pf.File)
		err = b.WriteFile(pf.TmplFile, bpFile)
		if nil != err {
			return
		}
	}

	//Create stack files from template
	for _, pf := range BpFileSet.Set["default"]["stack"] {
		bpFile := fmt.Sprintf("%s/%s", b.StackPath, pf.File)
		err = b.WriteFile(pf.TmplFile, bpFile)
		if nil != err {
			return
		}
	}

	//Create region files from template
	for _, r := range b.Regions {
		for _, f := range BpFileSet.Set["default"]["region"] {
			bpFile := fmt.Sprintf("%s/%s/%s", b.StackPath, r.Region, f.File)
			err = BpFileSet.WriteObjToFile(f.TmplFile, bpFile, r)
			if nil != err {
				return
			}
		}
	}

	return
}

func (p *Blueprint) WriteFile(tmplFile string, file string) (err error) {
	return BpFileSet.WriteObjToFile(tmplFile, file, p)
}
