package blueprint

import "github.com/spycode-io/spycli/assets"

var (
	DefaultFileSet map[string][]assets.FileTmpl = map[string][]assets.FileTmpl{
		"blueprint": []assets.FileTmpl{
			{TmplFile: "gitignore.tmpl", File: ".gitignore"},
		},
		"stack": []assets.FileTmpl{},
		"region": []assets.FileTmpl{
			{TmplFile: "region.hcl.tmpl", File: "region.hcl"},
		},
	}
)
