package lib

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func CreateFile(path string) (file *os.File, err error) {
	err = os.MkdirAll(filepath.Dir(path), os.ModeSticky|os.ModePerm)
	if err != nil {
		return
	}
	file, err = os.Create(path)
	return
}

func CopyFile(src string, dst string) (err error) {

	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()

	err = os.MkdirAll(filepath.Dir(dst), os.ModeSticky|os.ModePerm)
	if err != nil {
		return
	}

	// Create new file
	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}
	defer dstFile.Close()

	// Copy the bytes to destination from source
	bytesWritten, err := io.Copy(dstFile, srcFile)
	if err != nil {
		return
	}

	log.Printf("%d bytes writen from %s to %s", bytesWritten, src, dst)
	err = dstFile.Sync()
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
