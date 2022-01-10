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

type ProjectConfig struct {
	Name          string   `yaml:"name"`
	SlugName      string   `yaml:"slugName"`
	BluePrint     string   `yaml:"blueprint"`
	BluePrintPath string   `yaml:"blueprintPath"`
	Stack         string   `yaml:"stack"`
	BasePath      string   `yaml:"basePath"`
	ProjectPath   string   `yaml:"projectPath"`
	Ignore        []string `yaml:"ignore"`
}

type RegionConfig struct {
	Region   string `yaml:"region"`
	BasePath string `yaml:"basePath"`
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
					{TmplFile: "gitignore_region.tmpl", File: ".gitignore"},
				},
			},
		},
	}
)
