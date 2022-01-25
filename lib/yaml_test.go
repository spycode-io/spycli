package lib_test

import (
	"os"
	"testing"

	"github.com/spycode-io/spycli/lib"
)

const (
	dstYml = `foo: foo
propertie2: 2
`
	srcYml = `foo: bar
propertie1: 1
`
)

func TestMergeYml(t *testing.T) {
	os.RemoveAll(".iac-test/yaml-test")

	srcFile, err := lib.CreateFile(".iac-test/yaml-test/src.yml")
	if nil != err {
		t.Error(err)
	}

	_, err = srcFile.Write([]byte(srcYml))
	if nil != err {
		t.Error(err)
	}

	srcFile.Close()

	dstFile, err := lib.CreateFile(".iac-test/yaml-test/dst.yml")
	if nil != err {
		t.Error(err)
	}

	_, err = dstFile.Write([]byte(dstYml))
	if nil != err {
		t.Error(err)
	}
	dstFile.Close()

	_, _, err = lib.MergeYaml(".iac-test/yaml-test/src.yml", ".iac-test/yaml-test/dst.yml")
	if nil != err {
		t.Error(err)
	}
}
