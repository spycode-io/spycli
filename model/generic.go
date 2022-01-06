package model

import (
	"log"

	"github.com/spycode-io/spycli/assets"
)

type GenericScaffold struct {
	*Scaffold
	Foo string
	Bar string
}

func NewGenericScafold(base *Scaffold, foo string, bar string) (gen *GenericScaffold, err error) {

	gen = &GenericScaffold{
		Scaffold: base,
		Foo:      foo,
		Bar:      bar,
	}

	gen.FileSet.WithFiles("default", "generic", []assets.FileTmpl{
		{TmplFile: "bar.tmpl", File: "bar"},
		{TmplFile: "foo.tmpl", File: "foo"},
	})

	err = gen.Init()
	return
}

func (g *GenericScaffold) Init() (err error) {
	log.Printf("Initializing scaffold %s on path %s", g.Name, g.BasePath)
	g.FileSet.WriteObjToFiles("default", "generic", g)
	return
}
