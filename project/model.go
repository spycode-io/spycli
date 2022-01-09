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
	DefaultRegions      []string = []string{"us-east-1"}
	DefaultEnvironments []string = []string{"dev"}

	DefaultFileSet assets.FileSet = assets.FileSet{
		AssetsPath: "templates/prj",
		Set: map[string]map[string][]assets.FileTmpl{
			"aws": {
				"platform": []assets.FileTmpl{},
				"project": []assets.FileTmpl{
					{TmplFile: "prj.yml.tmpl", File: "prj.yml"},
					{TmplFile: "gitignore.tmpl", File: ".gitignore"},
					{TmplFile: "terragrunt.hcl.tmpl", File: "terragrunt.hcl"},
				},
				"environment": []assets.FileTmpl{
					{TmplFile: "env.yml.tmpl", File: "env.yml"},
				},
				"region": []assets.FileTmpl{
					{TmplFile: "region.yml.tmpl", File: "region.yml"},
				},
			},
		},
	}
)
