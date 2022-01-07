package module

import (
	"fmt"
	"log"

	"github.com/gosimple/slug"
	"github.com/spycode-io/spycli/assets"
	"github.com/spycode-io/spycli/model"
)

type Module struct {
	Scaffold   *model.Scaffold
	Module     string
	ModulePath string
	Lib        string
	LibVersion string
	Stack      string
	StackPath  string
	Regions    []model.Region
}

type ModuleInterface interface {
	InitModule() error
	WriteFile(tmplFile string, file string) (err error)
}

var DefaultFileSet map[string][]assets.FileTmpl = map[string][]assets.FileTmpl{
	"stack": {},
	"module": {
		{TmplFile: "terragrunt.hcl.tmpl", File: "terragrunt.hcl"},
	},
	"region": {
		{TmplFile: "gitignore.tmpl", File: ".gitignore"},
	},
}

func NewModule(
	base *model.Scaffold,
	stack string,
	lib string,
	libVersion string,
	regions []string) (*Module, error) {

	module := &Module{
		Scaffold:   base,
		Lib:        lib,
		LibVersion: libVersion,
		Stack:      stack,
		StackPath:  fmt.Sprintf("%s/%s", base.BasePath, stack),
		Module:     slug.Make(base.Name),
		//ModulePath: fmt.Sprintf("%s/%s/%s", base.BasePath, stack, slug.Make(base.Name)),
	}

	for _, reg := range regions {
		module.Regions = append(module.Regions,
			model.Region{Region: reg},
		)
	}

	module.Scaffold.FileSet.WithMap("default", DefaultFileSet)

	return module, module.InitModule()
}

func (m *Module) InitModule() (err error) {

	log.Printf("Initializing module %s [%s/%s]", m.Scaffold.Name, m.Lib, m.Module)

	//Create module folder if necessary
	// _, err = os.Stat(m.ModulePath)
	// if os.IsNotExist(err) {
	// 	os.MkdirAll(m.ModulePath, os.ModePerm)
	// 	err = nil
	// }

	//Write stack files
	err = m.Scaffold.FileSet.WriteObjToPath("default", "stack", m.StackPath, m)
	if nil != err {
		return
	}

	//Write stack files

	//Write region files
	for _, reg := range m.Regions {

		regPath := fmt.Sprintf("%s/%s", m.StackPath, reg.Region)
		err = m.Scaffold.FileSet.WriteObjToPath("default", "region", regPath, m)
		if nil != err {
			return
		}

		modulePath := fmt.Sprintf("%s/%s", regPath, m.Scaffold.SlugName)
		err = m.Scaffold.FileSet.WriteObjToPath("default", "module", modulePath, m)
		if nil != err {
			return
		}
	}

	return
}