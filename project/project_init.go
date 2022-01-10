package project

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spycode-io/spycli/lib"

	cp "github.com/otiai10/copy"
	yaml "gopkg.in/yaml.v2"
)

func InitProject(basePath string, linkInit bool) (err error) {

	log.Printf("Initializing project on %s", basePath)
	prjConfigFiles, err := GetConfigFiles(basePath, "prj.yml")

	if err != nil {
		return
	}

	for _, prj := range prjConfigFiles {
		var prjConfig *ProjectConfig
		prjConfig, err = GetProjectConfig(prj)
		if err != nil {
			return
		}

		DoInit(prjConfig, linkInit)
	}
	return
}

func DoInit(prjConfig *ProjectConfig, linkInit bool) (err error) {

	regionConfigFiles, err := GetConfigFiles(prjConfig.BasePath, "region.yml")

	if err != nil {
		return
	}

	for _, reg := range regionConfigFiles {
		var regConfig *RegionConfig
		regConfig, err = GetRegionConfig(reg)
		if err != nil {
			return
		}

		//add modules from _any folder
		source := fmt.Sprintf("%s/%s/%s", prjConfig.BluePrintPath, prjConfig.Stack, "_any")
		SyncBlueprintFolders(source, regConfig.BasePath, prjConfig.Ignore, linkInit)

		regFolder := filepath.Base(regConfig.BasePath)
		source = fmt.Sprintf("%s/%s/%s", prjConfig.BluePrintPath, prjConfig.Stack, regFolder)
		SyncBlueprintFolders(source, regConfig.BasePath, prjConfig.Ignore, linkInit)
	}

	return
}

func SyncBlueprintFolders(workingFolder string, destinyFolder string, ignoreFolders []string, linkInit bool) error {
	if linkInit {
		return LinkBlueprintFolders(workingFolder, destinyFolder, ignoreFolders, true)
	} else {
		return CopyBlueprintFolders(workingFolder, destinyFolder, ignoreFolders, true)
	}
}

func GetProjectConfig(filePath string) (prjConfig *ProjectConfig, err error) {
	var data []byte

	data, err = os.ReadFile(filePath)
	if err != nil {
		return
	}
	err = yaml.Unmarshal([]byte(data), &prjConfig)
	if err != nil {
		return
	}

	prjConfig.BasePath, err = filepath.Abs(filepath.Dir(filePath))
	if err != nil {
		return
	}
	prjConfig.BluePrintPath, err = filepath.Abs(prjConfig.BluePrint)

	return
}

func GetRegionConfig(filePath string) (regConfig *RegionConfig, err error) {
	var data []byte

	data, err = os.ReadFile(filePath)
	if err != nil {
		return
	}
	err = yaml.Unmarshal([]byte(data), &regConfig)
	if err != nil {
		return
	}

	regConfig.BasePath, err = filepath.Abs(filepath.Dir(filePath))
	return
}

func GetConfigFiles(location string, fileName string) (configFiles []string, err error) {

	configFiles = make([]string, 0)

	err = filepath.Walk(location,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && info.Name() == fileName {
				configFiles = append(configFiles, path)
			}
			return nil
		})
	return
}

func LinkBlueprintFolders(workingFolder string, destinyFolder string, ignoreFolders []string, verbose bool) (err error) {

	log.Printf("Linking blueprint folder %s -> %s", workingFolder, destinyFolder)

	folders, err := ioutil.ReadDir(workingFolder)

	for _, f := range folders {

		if f.IsDir() && f.Name() != ".git" && !lib.StringInSlice(f.Name(), ignoreFolders) {

			source := fmt.Sprintf("%s/%s", workingFolder, f.Name())
			dest := fmt.Sprintf("%s/%s", destinyFolder, f.Name())

			CleanStackFolder(dest, verbose)

			log.Println(source)
			err = lib.LinkChild(source, dest, ignoreFolders, verbose)
		}
	}

	return
}

func CleanStackFolder(stackFolder string, verbose bool) (err error) {
	if verbose {
		log.Print("Cleaning folder ", stackFolder)
	}

	files, err := os.ReadDir(stackFolder)
	if err != nil {
		return
	}

	for _, f := range files {
		if !f.Type().IsRegular() {
			fpath := filepath.Join(stackFolder, f.Name())
			if verbose {
				log.Print("Deleting folder ", fpath)
			}
			os.RemoveAll(fpath)
		}
	}
	return
}

func CopyBlueprintFolders(workingFolder string, destinyFolder string, ignoreFolders []string, verbose bool) (err error) {

	opt := cp.Options{
		Skip: func(src string) (bool, error) {
			f, err := os.Stat(src)
			doCopy := f.IsDir() || f.Name() != ".git" && !lib.StringInSlice(f.Name(), ignoreFolders)
			return !doCopy, err
		},
	}

	folders, err := ioutil.ReadDir(workingFolder)

	for _, f := range folders {
		source := fmt.Sprintf("%s/%s", workingFolder, f.Name())
		dest := fmt.Sprintf("%s/%s", destinyFolder, f.Name())

		log.Printf("Copping blueprint folder %s -> %s", source, dest)

		err = cp.Copy(source, dest, opt)

		if nil != err {
			return
		}
	}

	return
}
