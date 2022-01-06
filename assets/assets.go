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

type FileSet struct {
	AssetsPath string
	Set        map[string]map[string][]FileTmpl
	Data       embed.FS
}

type FileSetInterface interface {
	WithSet(set map[string]map[string][]FileTmpl) *FileSet
	WithMap(platform string, fileMap map[string][]FileTmpl) *FileSet
	WithFiles(platform string, level string, files []FileTmpl) *FileSet
	WriteObjToFile(tmplFile string, file string, obj interface{}) error
}

func NewFileSet(assetsPath string) *FileSet {
	return &FileSet{
		AssetsPath: assetsPath,
		Data:       TemplatesData,
		Set:        make(map[string]map[string][]FileTmpl),
	}
}

func (f *FileSet) WithMap(platform string, fileMap map[string][]FileTmpl) *FileSet {
	f.Set[platform] = fileMap
	return f
}

func (f *FileSet) WithSet(set map[string]map[string][]FileTmpl) *FileSet {
	f.Set = set
	return f
}

func (f *FileSet) WithFiles(platform string, level string, files []FileTmpl) *FileSet {
	if _, ok := f.Set[platform][level]; !ok {
		f.Set[platform] = make(map[string][]FileTmpl)
	}
	f.Set[platform][level] = files
	return f
}

func (f *FileSet) WriteObjToFile(tmplFile string, file string, obj interface{}) error {
	tmplPath := tmplFile
	if f.AssetsPath != "" {
		tmplPath = fmt.Sprintf("%s/%s", f.AssetsPath, tmplFile)
	}

	log.Printf("Writing file %s from template %s", file, tmplPath)

	pTmpl, err := template.ParseFS(f.Data, tmplPath)
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

func (f *FileSet) WriteObjToFiles(platform string, level string, obj interface{}) (err error) {
	for _, tf := range f.Set[platform][level] {
		err = f.WriteObjToFile(tf.TmplFile, tf.File, obj)
		if nil != err {
			return
		}
	}
	return
}
