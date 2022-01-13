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
	ModuleUrl  string
	ModulePath string
	IsLocal    bool
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

const LibraryLocalRelativePath = "../../../.."

func NewModule(
	base *model.Scaffold,
	moduleUrl string,
	isLocal bool) (*Module, error) {

	if isLocal {
		moduleUrl = fmt.Sprintf("%s/%s", LibraryLocalRelativePath, slug.Make(base.Name))
	}

	module := &Module{
		Scaffold:   base,
		IsLocal:    isLocal,
		ModuleUrl:  moduleUrl,
		ModulePath: fmt.Sprintf("%s/%s", base.BasePath, slug.Make(base.Name)),
	}

	module.Scaffold.FileSet.WithMap("new", NewFileSet)

	return module, module.InitModule()
}

func (m *Module) InitModule() (err error) {

	log.Printf("Initializing module %s [%s]", m.Scaffold.Name, m.ModuleUrl)

	//Write module files
	err = m.Scaffold.FileSet.WriteObjToPath("new", "module", m.ModulePath, m)
	if nil != err {
		return
	}

	return
}
