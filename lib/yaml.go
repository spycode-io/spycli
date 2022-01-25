package lib

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"go.uber.org/config"
	"gopkg.in/yaml.v2"
)

func MergeYaml(src string, dst string) (
	srcYamlMap map[string]interface{},
	dstYamlMap map[string]interface{},
	err error) {

	var srcFile, dstFile *os.File

	if srcFile, err = os.Open(src); err != nil {
		return
	}

	if dstFile, err = os.Open(dst); err != nil {
		return
	}

	//var dstYamlOption, srcYamlOption config.Source
	srcOpt := config.Source(srcFile)
	dstOpt := config.Source(dstFile)

	var (
		srcYaml, dstYaml *config.YAML
	)

	if srcYaml, err = config.NewYAML(srcOpt); err != nil {
		return
	}

	if dstYaml, err = config.NewYAML(srcOpt, dstOpt); err != nil {
		return
	}

	dstFile.Close()
	srcFile.Close()

	//var result map[string]interface{}
	if err = srcYaml.Get(config.Root).Populate(&srcYamlMap); err != nil {
		return
	}

	if err = dstYaml.Get(config.Root).Populate(&dstYamlMap); err != nil {
		return
	}

	return
}

func WriteToYaml(path string, obj interface{}) (err error) {

	if err = os.MkdirAll(filepath.Dir(path), os.ModeSticky|os.ModePerm); err != nil {
		return
	}

	var data []byte
	if data, err = yaml.Marshal(&obj); nil != err {
		return
	}

	err = ioutil.WriteFile(path, data, os.ModeSticky|os.ModePerm)

	return
}
