package module_test

import (
	"testing"

	"github.com/spycode-io/spycli/model"
	"github.com/spycode-io/spycli/module"
	"github.com/spycode-io/spycli/project"
)

func TestNewModule(t *testing.T) {

	base := model.NewScaffold("My Module", ".iac-test", "templates/mdl")

	_, err := module.NewModule(base, "my-stack", "my-lib", "v1", project.DefaultRegions)

	if nil != err {
		t.Error(err)
	}
}
