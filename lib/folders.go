package lib

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func CopyDir(source string, dest string, ignoreFolders []string, verbose bool) (err error) {

	// get properties of source dir
	var sourceinfo fs.FileInfo
	sourceinfo, err = os.Stat(source)
	if err != nil {
		return
	}

	// create dest dir
	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()
		destinationfilepointer := dest + "/" + obj.Name()

		//Verify if folder is ignored
		if obj.IsDir() && StringInSlice(obj.Name(), ignoreFolders) {
			if verbose {
				log.Print(fmt.Sprintf("ignoring %s", destinationfilepointer))
			}
			continue
		}

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer, ignoreFolders, verbose)
		} else {
			// perform copy
			err = CopyFile(sourcefilepointer, destinationfilepointer)
		}
	}
	return
}

func LinkChild(source string, dest string, ignoreFolders []string, verbose bool) (err error) {

	_, err = os.Stat(source)
	if err != nil {
		return
	}

	sourcePath, err := filepath.Abs(source)
	if err != nil {
		return
	}

	destinyPath, err := filepath.Abs(dest)
	if err != nil {
		return
	}

	err = os.Symlink(sourcePath, destinyPath)

	return
}
