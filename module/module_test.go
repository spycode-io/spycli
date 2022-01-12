package module_test

import (
	"testing"

	"github.com/spycode-io/spycli/model"
	"github.com/spycode-io/spycli/module"
)

func TestNewModule(t *testing.T) {

	base := model.NewScaffold("My VPC", ".iac-test", "templates/mdl")
	_, err := module.NewModule(base, "tf-modules/vpc", false)

	if nil != err {
		t.Error(err)
	}
}

func TestNewLocalModule(t *testing.T) {

	base := model.NewScaffold("My VPC Local", ".iac-test", "templates/mdl")
	_, err := module.NewModule(base, "tf-modules/vpc", false)

	if nil != err {
		t.Error(err)
	}
}
