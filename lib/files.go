package lib

import (
	"log"
	"os"
	"path/filepath"
)

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

func FileExists(path string) (ex bool) {
	_, err := os.Stat(path)
	ex = true
	if os.IsNotExist(err) {
		ex = false
		log.Printf("File %s does not exists", path)
		return
	}
	return
}
