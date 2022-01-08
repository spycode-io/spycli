package project

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func InitProject(basPath string) (err error) {

	log.Printf("initializing project on %s", basPath)
	configFiles, err := GetConfigFiles(basPath, "prj.hcl")

	if err != nil {
		return
	}

	fmt.Println(configFiles)
	return nil
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
