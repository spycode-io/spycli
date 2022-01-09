package project

import (
	"github.com/spycode-io/spycli/assets"
	"github.com/spycode-io/spycli/model"
)

type ProjectScaffold struct {
	model.Scaffold
	Platform         string
	PlatformPath     string
	ProjectPath      string
	Stack            string
	Blueprint        string
	BlueprintVersion string
	Environments     []model.Environment
	Regions          []model.Region
}

var (
	DefaultRegions      []string       = []string{"us-east-1"}
	DefaultEnvironments []string       = []string{"stage", "qa", "prod"}
	DefaultFileSet      assets.FileSet = assets.FileSet{
		AssetsPath: "templates/prj",
		Set: map[string]map[string][]assets.FileTmpl{
			"aws": {
				"platform": []assets.FileTmpl{
					{TmplFile: "gitignore.tmpl", File: ".gitignore"},
					{TmplFile: "terragrunt.hcl.tmpl", File: "terragrunt.hcl"},
				},
				"project": []assets.FileTmpl{
					{TmplFile: "prj.hcl.tmpl", File: "prj.hcl"},
				},
				"environment": []assets.FileTmpl{
					{TmplFile: "env.hcl.tmpl", File: "env.hcl"},
				},
				"region": []assets.FileTmpl{
					{TmplFile: "region.hcl.tmpl", File: "region.hcl"},
				},
			},
		},
	}
)
