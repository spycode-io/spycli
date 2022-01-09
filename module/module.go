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
}

type ModuleInterface interface {
	InitModule() error
	WriteFile(tmplFile string, file string) (err error)
}

var NewFileSet map[string][]assets.FileTmpl = map[string][]assets.FileTmpl{
	"module": {
		{TmplFile: "terragrunt.hcl.tmpl", File: "terragrunt.hcl"},
	},
}

var InstallFileSet map[string][]assets.FileTmpl = map[string][]assets.FileTmpl{
	"module": {
		{TmplFile: "terragrunt.hcl.tmpl", File: "terragrunt.hcl"},
	},
}

func NewModule(
	base *model.Scaffold,
	moduleName string) (*Module, error) {

	module := &Module{
		Scaffold:   base,
		Module:     moduleName,
		ModulePath: fmt.Sprintf("%s/%s", base.BasePath, slug.Make(base.Name)),
	}

	module.Scaffold.FileSet.WithMap("new", NewFileSet)

	return module, module.InitModule()
}

func (m *Module) InitModule() (err error) {

	log.Printf("Initializing module %s [%s/%s]", m.Scaffold.Name, m.Lib, m.Module)

	//Write module files
	err = m.Scaffold.FileSet.WriteObjToPath("new", "module", m.ModulePath, m)
	if nil != err {
		return
	}

	return
}
