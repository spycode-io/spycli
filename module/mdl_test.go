package module_test

import (
	"testing"

	"github.com/spycode-io/spycli/module"
)

func TestNewModule(t *testing.T) {
	_, err := module.NewModule(
		"my-vpc",
		".iac-test",
		"git@github.com:spycode-io/tf-modules.git",
		"master",
		"aws",
		"vpc")

	if nil != err {
		t.Error(err)
	}
}
