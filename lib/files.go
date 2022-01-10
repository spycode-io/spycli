package lib

import (
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
