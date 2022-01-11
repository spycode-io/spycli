package assets_test

import (
	"io/ioutil"
	"testing"

	"github.com/spycode-io/spycli/assets"
)

type testObj struct {
	Test string
}

func TestWriteFile(t *testing.T) {

	obj := testObj{
		Test: "Test Write File",
	}

	fileSet := assets.NewFileSet("templates")
	err := fileSet.WriteObjToFile("test.tmpl", ".iac-test/test-file", obj)
	if nil != err {
		t.Error(err)
	}

	fileContent, err := ioutil.ReadFile(".iac-test/test-file")
	if nil != err {
		t.Error(err)
	}

	if string(fileContent) != obj.Test {
		t.Errorf("Invalid content: want be %s but %s found", fileContent, obj.Test)
	}
}
