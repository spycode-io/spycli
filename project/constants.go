package project

import "github.com/spycode-io/spycli/assets"

var (
	DefaultRegions      []string       = []string{"east-us1", "east-us2", "west-us1"}
	DefaultEnvironments []string       = []string{"stage", "qa", "prod"}
	DefaultFileSet      assets.FileSet = assets.FileSet{
		AssetsPath: "templates/prj",
		Set: map[string]map[string][]assets.FileTmpl{
			"aws": {
				"platform": []assets.FileTmpl{
					{TmplFile: "gitignore.tmpl", File: ".gitignore"},
				},
				"project": []assets.FileTmpl{
					{TmplFile: "prj.hcl.tmpl", File: "prj.hcl"},
				},
				"environment": []assets.FileTmpl{
					{TmplFile: "env.hcl.tmpl", File: "env.hcl"},
				},
				"region": []assets.FileTmpl{
					{TmplFile: "region.hcl.tmpl", File: "region.hcl"},
					{TmplFile: "gitignore_region.tmpl", File: ".gitignore"},
				},
			},
		},
	}
)
