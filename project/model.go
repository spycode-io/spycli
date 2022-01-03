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
}

type ProjectInterface interface {
	WriteFile(tmplPath string, filePath string) error
	InitProject() error
}

func New(assetsData embed.FS, baseDirectory string, kind string, name string) (*Project, error) {

	project := &Project{
		BasePath:       baseDirectory,
		Kind:           kind,
		Name:           name,
		SlugName:       slug.Make(name),
		ProjectPath:    fmt.Sprintf("%s/%s", baseDirectory, slug.Make(name)),
		AssetsData:     assetsData,
		AssetsBasePath: fmt.Sprintf("assets/%s", kind),
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

	err = os.MkdirAll(p.ProjectPath, os.ModePerm)
	if nil != err {
		return
	}

	err = p.WriteFile("gitignore.tmpl", ".gitignore")
	if nil != err {
		return
	}

	err = p.WriteFile("terragrunt.hcl.tmpl", "terragrunt.hcl")
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
