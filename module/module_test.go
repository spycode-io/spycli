package module_test

import (
	"testing"

	"github.com/spycode-io/spycli/model"
	"github.com/spycode-io/spycli/module"
)

func TestNewModule(t *testing.T) {

	base := model.NewScaffold("My Module", ".iac-test", "templates/mdl")

	_, err := module.NewModule(base)

	if nil != err {
		t.Error(err)
	}
}
