package module_test

import (
	"testing"

	"github.com/spycode-io/spycli/model"
	"github.com/spycode-io/spycli/module"
)

func TestNewModule(t *testing.T) {

	base := model.NewScaffold("My VPC", ".iac-test", "templates/mdl")

	_, err := module.NewModule(base, "vpc")

	if nil != err {
		t.Error(err)
	}
}
