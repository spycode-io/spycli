package project

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/gosimple/slug"
	"github.com/spycode-io/spycli/assets"
	"github.com/spycode-io/spycli/model"
)

type Project struct {
	Name             string
	SlugName         string
	BasePath         string
	Platform         string
	PlatformBasePath string
	AssetsBasePath   string
	ProjectPath      string
	Stack            string
	Blueprint        string
	BlueprintVersion string
	Environments     []model.Environment
	Regions          []model.Region
	AssetsData       embed.FS
}

type ProjectInterface interface {
	WriteFile(tmplFile string, file string) (err error)
	// WriteObjToFile(tmplFile string, file string, obj interface{}) error
	InitProject() error
}

var (
	DefaultRegions      []string       = []string{"east-us1", "east-us2", "west-us1"}
	DefaultEnvironments []string       = []string{"stage", "qa", "prod"}
	ProjectFileSet      assets.FileSet = assets.FileSet{
		RelativeRootPath: "templates/prj",
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

func New(
	templatesData embed.FS,
	baseDirectory string,
	kind string,
	name string,
	stack string,
	blueprint string,
	blueprintVersion string,
	environments []string,
	regions []string) (*Project, error) {

	project := &Project{
		Name:             name,
		BasePath:         baseDirectory,
		Platform:         kind,
		PlatformBasePath: fmt.Sprintf("%s/%s", baseDirectory, kind),
		SlugName:         slug.Make(name),
		ProjectPath:      fmt.Sprintf("%s/%s/%s", baseDirectory, kind, slug.Make(name)),
		AssetsData:       templatesData,
		AssetsBasePath:   fmt.Sprintf("assets/%s", kind),
		Stack:            stack,
		Blueprint:        blueprint,
		BlueprintVersion: blueprintVersion,
	}

	for _, env := range environments {
		project.Environments = append(project.Environments,
			model.Environment{EnvName: env, BaseEnvPath: fmt.Sprintf("%s/%s", project.ProjectPath, env)},
		)
	}

	for _, reg := range regions {
		project.Regions = append(project.Regions,
			model.Region{Region: reg},
		)
	}

	return project, project.InitProject()
}

func (p *Project) InitProject() (err error) {

	log.Printf("Initializing %s [%s] %s project on path: %s", p.Name, p.SlugName, p.Platform, p.BasePath)

	//Create base folder if necessary
	_, err = os.Stat(p.BasePath)
	if os.IsNotExist(err) {
		os.MkdirAll(p.BasePath, os.ModePerm)
	}

	//Create the project base foder
	err = os.MkdirAll(p.ProjectPath, os.ModePerm)
	if nil != err {
		return
	}

	//Write project file
	projectFile := fmt.Sprintf("%s/prj.hcl", p.ProjectPath)
	err = p.WriteFile("prj.hcl.tmpl", projectFile)
	if nil != err {
		return
	}

	//Create folder structure
	for _, env := range p.Environments {
		for _, reg := range DefaultRegions {
			regPath := fmt.Sprintf("%s/%s", env.BaseEnvPath, reg)
			err = os.MkdirAll(regPath, os.ModePerm)
			if nil != err {
				return
			}
		}
	}

	//Create platform files from template by platform
	for _, pf := range ProjectFileSet.Set[p.Platform]["platform"] {
		platformFile := fmt.Sprintf("%s/%s", p.PlatformBasePath, pf.File)
		err = p.WriteFile(pf.TmplFile, platformFile)
		if nil != err {
			return
		}
	}

	//Create project files from template by platform
	for _, pf := range ProjectFileSet.Set[p.Platform]["project"] {
		platformFile := fmt.Sprintf("%s/%s", p.ProjectPath, pf.File)
		err = p.WriteFile(pf.TmplFile, platformFile)
		if nil != err {
			return
		}
	}

	//Create environment files from template
	for _, env := range p.Environments {
		for _, pf := range ProjectFileSet.Set[p.Platform]["environment"] {
			platformFile := fmt.Sprintf("%s/%s", env.BaseEnvPath, pf.File)
			err = ProjectFileSet.WriteObjToFile(pf.TmplFile, platformFile, env)
			if nil != err {
				return
			}
		}
	}

	//Create region files from template
	for _, env := range p.Environments {
		for _, reg := range p.Regions {
			for _, pf := range ProjectFileSet.Set[p.Platform]["region"] {
				regionFile := fmt.Sprintf("%s/%s/%s", env.BaseEnvPath, reg.Region, pf.File)
				err = ProjectFileSet.WriteObjToFile(pf.TmplFile, regionFile, reg)
				if nil != err {
					return
				}
			}
		}
	}

	return
}

func (p *Project) WriteFile(tmplFile string, file string) (err error) {
	return ProjectFileSet.WriteObjToFile(tmplFile, file, p)
}

// func writeFile(project *Project, assetsData embed.FS, tmplPath string, filePath string) error {
// 	tmpl, err := template.ParseFS(assetsData, fmt.Sprintf(tmplPath, project.Kind))
// 	if err != nil {
// 		return err
// 	}

// 	f, err := os.Create(filePath)
// 	if err != nil {
// 		return err
// 	}

// 	defer f.Close()
// 	return tmpl.Execute(f, project)
// }
