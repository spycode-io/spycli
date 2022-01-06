package module

import (
	"fmt"
	"log"
	"os"

	"github.com/gosimple/slug"
	"github.com/spycode-io/spycli/assets"
)

type Module struct {
	Name       string
	BasePath   string
	SlugName   string
	Module     string
	ModulePath string
	Lib        string
	LibVersion string
	Stack      string
	StackPath  string
}

type ModuleInterface interface {
	InitModule() error
	WriteFile(tmplFile string, file string) (err error)
}

var ModuleFileSet assets.FileSet = assets.FileSet{
	RelativeRootPath: "templates/mdl",
	Set: map[string]map[string][]assets.FileTmpl{
		"default": {
			"stack": []assets.FileTmpl{
				{TmplFile: "gitignore.tmpl", File: ".gitignore"},
			},
			"module": []assets.FileTmpl{
				{TmplFile: "terragrunt.hcl.tmpl", File: "terragrunt.hcl"},
			},
		},
	},
}

func NewModule(
	name string,
	basePath string,
	lib string,
	libVersion string,
	stack string,
	module string) (*Module, error) {

	mdl := &Module{
		Name:       name,
		BasePath:   basePath,
		Lib:        lib,
		LibVersion: libVersion,
		Stack:      stack,
		Module:     module,
		SlugName:   slug.Make(name),
		StackPath:  fmt.Sprintf("%s/%s", basePath, stack),
		ModulePath: fmt.Sprintf("%s/%s/%s", basePath, stack, slug.Make(name)),
	}

	return mdl, mdl.InitModule()
}

func (m *Module) InitModule() (err error) {

	log.Printf("Initializing module %s [%s/%s]", m.Name, m.Lib, m.Module)

	//Create module folder if necessary
	_, err = os.Stat(m.ModulePath)
	if os.IsNotExist(err) {
		os.MkdirAll(m.ModulePath, os.ModePerm)
		err = nil
	}

	//Create module files from template
	for _, pf := range ModuleFileSet.Set["default"]["stack"] {
		bpFile := fmt.Sprintf("%s/%s", m.StackPath, pf.File)
		err = m.WriteFile(pf.TmplFile, bpFile)
		if nil != err {
			return
		}
	}

	//Create module files from template
	for _, pf := range ModuleFileSet.Set["default"]["module"] {
		bpFile := fmt.Sprintf("%s/%s", m.ModulePath, pf.File)
		err = m.WriteFile(pf.TmplFile, bpFile)
		if nil != err {
			return
		}
	}

	return
}

func (p *Module) WriteFile(tmplFile string, file string) (err error) {
	return ModuleFileSet.WriteObjToFile(tmplFile, file, p)
}
