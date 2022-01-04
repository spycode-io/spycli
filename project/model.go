package project

import (
	"embed"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/gosimple/slug"
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
	Environments     []Environment
	Regions          []Region
	AssetsData       embed.FS
}

type Environment struct {
	EnvName     string
	BaseEnvPath string
}

type Region struct {
	Region string
}

type ProjectFile struct {
	TmplFile string
	File     string
}

type ProjectInterface interface {
	WriteFile(tmplFile string, file string) (err error)
	WriteObjToFile(tmplFile string, file string, obj interface{}) error
	InitProject() error
}

var (
	DefaultRegions      []string                            = []string{"east-us1", "east-us2", "west-us1"}
	DefaultEnvironments []string                            = []string{"stage", "qa", "prod"}
	PlatformFileSet     map[string]map[string][]ProjectFile = map[string]map[string][]ProjectFile{
		"aws": {
			"platform": []ProjectFile{
				{TmplFile: "gitignore.tmpl", File: ".gitignore"},
			},
			"project": []ProjectFile{
				{TmplFile: "prj.hcl.tmpl", File: "prj.hcl"},
			},
			"environment": []ProjectFile{
				{TmplFile: "env.hcl.tmpl", File: "env.hcl"},
			},
			"region": []ProjectFile{
				{TmplFile: "region.hcl.tmpl", File: "region.hcl"},
				{TmplFile: "gitignore_region.tmpl", File: ".gitignore"},
			},
		},
	}
)

func New(
	assetsData embed.FS,
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
		AssetsData:       assetsData,
		AssetsBasePath:   fmt.Sprintf("assets/%s", kind),
		Stack:            stack,
		Blueprint:        blueprint,
		BlueprintVersion: blueprintVersion,
	}

	for _, env := range environments {
		project.Environments = append(project.Environments,
			Environment{EnvName: env, BaseEnvPath: fmt.Sprintf("%s/%s", project.ProjectPath, env)},
		)
	}

	for _, reg := range regions {
		project.Regions = append(project.Regions,
			Region{Region: reg},
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

	//Write profect file
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
	for _, pf := range PlatformFileSet[p.Platform]["platform"] {
		platformFile := fmt.Sprintf("%s/%s", p.PlatformBasePath, pf.File)
		err = p.WriteFile(pf.TmplFile, platformFile)
		if nil != err {
			return
		}
	}

	//Create project files from template by platform
	for _, pf := range PlatformFileSet[p.Platform]["project"] {
		platformFile := fmt.Sprintf("%s/%s", p.ProjectPath, pf.File)
		err = p.WriteFile(pf.TmplFile, platformFile)
		if nil != err {
			return
		}
	}

	//Create environment files from template
	for _, env := range p.Environments {
		for _, pf := range PlatformFileSet[p.Platform]["environment"] {
			platformFile := fmt.Sprintf("%s/%s", env.BaseEnvPath, pf.File)
			err = p.WriteObjToFile(pf.TmplFile, platformFile, env)
			if nil != err {
				return
			}
		}
	}

	//Create region files from template
	for _, env := range p.Environments {
		for _, reg := range p.Regions {
			for _, pf := range PlatformFileSet[p.Platform]["region"] {
				regionFile := fmt.Sprintf("%s/%s/%s", env.BaseEnvPath, reg.Region, pf.File)
				err = p.WriteObjToFile(pf.TmplFile, regionFile, reg)
				if nil != err {
					return
				}
			}
		}
	}

	return
}

func (p *Project) WriteFile(tmplFile string, file string) (err error) {
	return p.WriteObjToFile(tmplFile, file, p)
}

func (p *Project) WriteObjToFile(tmplFile string, file string, obj interface{}) (err error) {
	tmplPath := fmt.Sprintf("assets/%s/%s", p.Platform, tmplFile)
	log.Printf("Writing file %s from template %s on project on path: %s", file, tmplPath, p.ProjectPath)

	pTmpl, err := template.ParseFS(p.AssetsData, tmplPath)
	if err != nil {
		return err
	}

	wf, err := os.Create(file)
	if err != nil {
		return err
	}

	defer wf.Close()

	return pTmpl.Execute(wf, obj)
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
