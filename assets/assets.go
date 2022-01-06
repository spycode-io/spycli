package assets

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/*
var TemplatesData embed.FS

type FileTmpl struct {
	TmplFile string
	File     string
}

type FileSetInterface interface {
	WriteObjToFile(tmplFile string, file string, obj interface{}) error
}

type FileSet struct {
	RelativeRootPath string
	Set              map[string]map[string][]FileTmpl
}

func (f *FileSet) WriteObjToFile(tmplFile string, file string, obj interface{}) error {
	tmplPath := tmplFile
	if f.RelativeRootPath != "" {
		tmplPath = fmt.Sprintf("%s/%s", f.RelativeRootPath, tmplFile)
	}

	log.Printf("Writing file %s from template %s", file, tmplPath)

	pTmpl, err := template.ParseFS(TemplatesData, tmplPath)
	if err != nil {
		return err
	}

	//Create base folder if necessary
	dir := filepath.Dir(file)
	_, err = os.Stat(dir)
	if os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
		err = nil
	}

	wf, err := os.Create(file)
	if err != nil {
		return err
	}

	defer wf.Close()

	return pTmpl.Execute(wf, obj)
}
