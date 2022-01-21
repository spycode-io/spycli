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

var DefaultIgnoredFiles []string = []string{"prj.yml", "env.yml", "region.yml", ".git"}

func InitProject(basePath string, linkInit bool) (err error) {

	fullBasePath, err := filepath.Abs(basePath)
	if err != nil {
		return
	}

	log.Printf("Initializing project on %s", fullBasePath)
	prjConfigFiles, err := GetConfigFiles(fullBasePath, "prj.yml")

	if err != nil {
		return
	}

	for _, prj := range prjConfigFiles {
		var prjConfig *ProjectConfig
		prjConfig, err = GetProjectConfig(prj)
		if err != nil {
			return
		}

		isValid, err := ValidateProject(prjConfig)
		if !isValid {
			log.Printf("Could not initialize: %s", err.Error())
		}

		DoInit(prjConfig, linkInit)
	}
	return
}

func DoInit(prjConfig *ProjectConfig, linkInit bool) (err error) {

	regionConfigFiles, err := GetConfigFiles(prjConfig.ProjectPath, "region.yml")

	if err != nil {
		return
	}

	for _, reg := range regionConfigFiles {
		var regConfig *RegionConfig
		regConfig, err = GetRegionConfig(reg)
		if err != nil {
			return
		}

		regFolder := filepath.Base(regConfig.BasePath)
		source := fmt.Sprintf("%s/%s/%s", prjConfig.BluePrintPath, prjConfig.Stack, "_any")
		SyncBlueprintFolders(source, regConfig.BasePath, prjConfig.Ignore, linkInit)

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

	prjConfig.ProjectPath, err = filepath.Abs(filepath.Dir(filePath))
	if err != nil {
		return
	}

	prjConfig.BasePath, err = filepath.Abs(filepath.Dir(filepath.Dir(filePath)))
	if err != nil {
		return
	}

	//for local blueprint
	prjConfig.BluePrintPath, err = filepath.Abs(fmt.Sprintf("%s/%s", prjConfig.BasePath, prjConfig.BluePrint))
	prjConfig.Ignore = append(prjConfig.Ignore, DefaultIgnoredFiles...)

	return
}

func GetEnvConfig(filePath string) (envConfig *EnvConfig, err error) {
	var data []byte

	data, err = os.ReadFile(filePath)
	if err != nil {
		return
	}
	err = yaml.Unmarshal([]byte(data), &envConfig)
	if err != nil {
		return
	}

	return
}

func ValidateProject(prjConfig *ProjectConfig) (isValid bool, err error) {
	isValid = lib.FileExists(prjConfig.ProjectPath) && lib.FileExists(prjConfig.BluePrintPath)
	if isValid {
		return
	}
	err = fmt.Errorf("the %s is valid ", prjConfig.Name)
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

	folders, err := ioutil.ReadDir(workingFolder)

	for _, f := range folders {

		if skip, _ := skipFile(f.Name(), ignoreFolders); !skip {

			source := fmt.Sprintf("%s/%s", workingFolder, f.Name())
			dest := fmt.Sprintf("%s/%s", destinyFolder, f.Name())

			log.Printf("Linking blueprint folder %s -> %s", source, dest)

			CleanStackFolder(dest, ignoreFolders)
			err = lib.LinkChild(source, dest, ignoreFolders, verbose)
		}
	}

	return
}

func CleanStackFolder(stackFolder string, ignoreFolders []string) (err error) {
	if skip, _ := skipFile(stackFolder, ignoreFolders); !skip {
		log.Println("Cleaning folder ", stackFolder)
		os.RemoveAll(stackFolder)
	}

	return
}

func CopyBlueprintFolders(workingFolder string, destinyFolder string, ignoreFolders []string, verbose bool) (err error) {

	folders, err := ioutil.ReadDir(workingFolder)

	for _, f := range folders {

		if skip, _ := skipFile(f.Name(), ignoreFolders); !skip {
			source := fmt.Sprintf("%s/%s", workingFolder, f.Name())
			dest := fmt.Sprintf("%s/%s", destinyFolder, f.Name())

			log.Printf("Copying blueprint folder %s -> %s", source, dest)
			CleanStackFolder(dest, ignoreFolders)

			err = cp.Copy(source, dest)
			if nil != err {
				return
			}
		}
	}

	return
}

func skipFile(path string, ignoreFolders []string) (skip bool, err error) {

	// if !lib.FileExists(path) {
	// 	skip = true
	// 	err = errors.New("file not found")
	// 	return
	// }

	f, err := os.Stat(path)

	if nil == err {
		skip = lib.StringInSlice(f.Name(), ignoreFolders)
	} else {
		skip = lib.StringInSlice(path, ignoreFolders)
	}

	if skip {
		log.Printf("Skiping file %s", path)
	}

	return
}
