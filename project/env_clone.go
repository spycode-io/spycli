package project

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/gosimple/slug"
	"github.com/spycode-io/spycli/lib"
)

func CloneEnv(basePath string, name string, src string) (err error) {

	fullBasePath, err := filepath.Abs(basePath)
	name = slug.Make(name)

	if err != nil {
		return
	}

	log.Printf("Cloning environment from %s to %s", src, name)

	prjConfigFiles, err := GetConfigFiles(fullBasePath, "prj.yml", DefaultIgnoredFolders)

	if err != nil {
		return
	}

	for _, prj := range prjConfigFiles {

		var prjConfig *ProjectConfig
		prjConfig, err = GetProjectConfig(prj)
		if err != nil {
			return
		}

		var isValid = false
		isValid, err = ValidateProject(prjConfig)
		if !isValid {
			log.Printf("Could not initialize: %s", err.Error())
		}

		srcPath := fmt.Sprintf("%s/%s", prjConfig.ProjectPath, src)
		if !lib.FileExists(srcPath) {
			err = errors.New("source environment not exists in project")
			return
		}

		//Create base file
		toPath := fmt.Sprintf("%s/%s", prjConfig.ProjectPath, name)
		toEnvFile := fmt.Sprintf("%s/env.yml", toPath)

		envConfig := &EnvConfig{Environment: name}
		if err = lib.WriteToYaml(toEnvFile, envConfig); nil != err {
			return
		}

		//Merge base file and source file
		srcEnvFile := fmt.Sprintf("%s/env.yml", srcPath)

		var dstYaml, srcYaml map[string]interface{}
		srcYaml, dstYaml, err = lib.MergeYaml(srcEnvFile, toEnvFile)

		if err != nil {
			return
		}

		dstYaml["ignore"] = srcYaml["ignore"]
		lib.WriteToYaml(toEnvFile, dstYaml)

		var files []string
		files, err = filepath.Glob(fmt.Sprintf("%s/*/region.yml", srcPath))
		if err != nil {
			return
		}

		for _, f := range files {
			toRegionFile := strings.Replace(f, src, name, 1)
			log.Printf("copying %s to %s", f, toRegionFile)
			err = lib.CopyFile(f, toRegionFile)
			if err != nil {
				return
			}
		}
	}
	return
}
