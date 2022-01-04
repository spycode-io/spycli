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
	BasePath       string
	Kind           string
	Name           string
	SlugName       string
	AssetsBasePath string
	ProjectPath    string
	AssetsData     embed.FS
	Environments   []string
}

type ProjectFile struct {
	TmplFile string
	File     string
}

type ProjectInterface interface {
	WriteFile(tmplPath string, filePath string) error
	InitProject() error
}

var (
	DefaultRegions      []string                 = []string{"east-us1", "east-us2", "west-us1"}
	DefaultEnvironments []string                 = []string{"stage", "qa", "prod"}
	ProjectFileSet      map[string][]ProjectFile = map[string][]ProjectFile{
		"aws": {
			ProjectFile{TmplFile: "gitignore.tmpl", File: ".gitignore"},
			ProjectFile{TmplFile: "terragrunt.hcl.tmpl", File: "aws/terragrunt.hcl"},
		},
	}
)

func New(assetsData embed.FS, baseDirectory string, kind string, name string) (*Project, error) {

	project := &Project{
		BasePath:       baseDirectory,
		Kind:           kind,
		Name:           name,
		SlugName:       slug.Make(name),
		ProjectPath:    fmt.Sprintf("%s/%s", baseDirectory, slug.Make(name)),
		AssetsData:     assetsData,
		AssetsBasePath: fmt.Sprintf("assets/%s", kind),
		Environments:   DefaultEnvironments,
	}

	return project, project.InitProject()
}

func (p *Project) InitProject() (err error) {

	log.Printf("Initialing %s [%s] %s project on path: %s", p.Name, p.SlugName, p.Kind, p.BasePath)

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

	//Create folder structure
	for _, env := range p.Environments {
		envPath := fmt.Sprintf("%s/%s/%s", p.ProjectPath, p.Kind, env)
		for _, reg := range DefaultRegions {
			regPath := fmt.Sprintf("%s/%s", envPath, reg)
			err = os.MkdirAll(regPath, os.ModePerm)
			if nil != err {
				return
			}

			err = p.WriteFile("prj.hcl.tmpl", fmt.Sprintf("%s/%s/%s", p.Kind, env, "prj.hcl"))
			if nil != err {
				return
			}
		}

	}

	//Create files from template by kind
	for _, pf := range ProjectFileSet[p.Kind] {
		err = p.WriteFile(pf.TmplFile, pf.File)
		if nil != err {
			return
		}
	}

	return
}

func (p *Project) WriteFile(tmplFile string, file string) (err error) {

	tmplPath := fmt.Sprintf("%s/%s", p.AssetsBasePath, tmplFile)
	filePath := fmt.Sprintf("%s/%s", p.ProjectPath, file)

	log.Printf("Writing file %s from template %s on project on path: %s", filePath, tmplPath, p.BasePath)

	pTmpl, err := template.ParseFS(p.AssetsData, tmplPath)
	if err != nil {
		return err
	}

	wf, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer wf.Close()

	return pTmpl.Execute(wf, p)
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
