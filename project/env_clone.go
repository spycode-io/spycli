package project

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gosimple/slug"
	"github.com/spycode-io/spycli/lib"
	yaml "gopkg.in/yaml.v2"
)

func CloneEnv(basePath string, name string, from string) (err error) {

	fullBasePath, err := filepath.Abs(basePath)
	name = slug.Make(name)
	from = slug.Make(from)

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

		var isValid = false
		isValid, err = ValidateProject(prjConfig)
		if !isValid {
			log.Printf("Could not initialize: %s", err.Error())
		}

		fromPath := fmt.Sprintf("%s/%s", prjConfig.ProjectPath, from)
		if !lib.FileExists(fromPath) {
			err = errors.New("from environment not exists in project")
			return
		}

		toPath := fmt.Sprintf("%s/%s", prjConfig.ProjectPath, name)
		os.MkdirAll(toPath, 0755)

		//copy env config file
		fromEnvFile := fmt.Sprintf("%s/env.yml", fromPath)
		toEnvFile := fmt.Sprintf("%s/env.yml", toPath)

		err = lib.CopyFile(fromEnvFile, toEnvFile)
		if err != nil {
			return
		}

		var envConfig *EnvConfig
		envConfig, err = GetEnvConfig(toEnvFile)
		if err != nil {
			return
		}

		envConfig.Environment = name
		var data []byte
		data, err = yaml.Marshal(envConfig)

		if err != nil {
			return
		}

		err = ioutil.WriteFile(toEnvFile, data, 0666)
		if err != nil {
			return
		}

		var files []string
		files, err = filepath.Glob(fmt.Sprintf("%s/*/region.yml", fromPath))
		if err != nil {
			return
		}

		for _, f := range files {
			toRegionFile := strings.Replace(f, from, name, 1)
			log.Printf("copying %s to %s", f, toRegionFile)
			err = lib.CopyFile(f, toRegionFile)
			if err != nil {
				return
			}
		}

		//lib.CopyDir(fromPath, toPath, []string{".git", ".terragrunt-cache", ".terraform.lock.hcl"}, true)

	}
	return
}
