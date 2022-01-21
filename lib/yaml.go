package lib

import (
	"io/ioutil"
	"os"

	"go.uber.org/config"
	"gopkg.in/yaml.v2"
)

func MergeYaml(src string, dst string) (err error) {

	var srcFile, dstFile *os.File

	if srcFile, err = os.Open(src); err != nil {
		return
	}

	if dstFile, err = os.Open(dst); err != nil {
		return
	}

	var dstYml *config.YAML
	if dstYml, err = config.NewYAML(config.Source(dstFile), config.Source(srcFile)); err != nil {
		return
	}

	dstFile.Close()
	srcFile.Close()

	var result map[string]interface{}
	if err = dstYml.Get(config.Root).Populate(&result); err != nil {
		return
	}

	err = WriteToYaml(dst, result)

	return
}

func WriteToYaml(path string, obj interface{}) (err error) {

	var data []byte
	if data, err = yaml.Marshal(&obj); nil != err {
		return
	}

	err = ioutil.WriteFile(path, data, os.ModeSticky|os.ModePerm)

	return
}
