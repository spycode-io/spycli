package lib

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return
	}

	defer destfile.Close()
	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			_ = os.Chmod(dest, sourceinfo.Mode())
		}
	}
	return
}

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

	// directory, _ := os.Open(source)
	// objects, err := directory.Readdir(-1)
	// if err != nil {
	// 	return
	// }

	// for _, obj := range objects {
	// 	var sourcefilepointer, destinationfilepointer string

	// 	sourcefilepointer, err = filepath.Abs(source + "/" + obj.Name())
	// 	if err != nil {
	// 		return
	// 	}

	// 	destinationfilepointer, err = filepath.Abs(dest + "/" + obj.Name())

	// 	//Verify if folder is ignored
	// 	if obj.IsDir() && StringInSlice(obj.Name(), ignoreFolders) {
	// 		if verbose {
	// 			log.Print(fmt.Sprintf("ignoring %s", destinationfilepointer))
	// 		}
	// 		continue
	// 	}

	// 	if obj.IsDir() {
	// 		if verbose {
	// 			log.Printf("Linking folder %s -> %s", sourcefilepointer, destinationfilepointer)
	// 		}
	// 		err = os.Symlink(sourcefilepointer, destinationfilepointer)
	// 	}

	// 	if err != nil {
	// 		return
	// 	}
	// }
	// return
}
