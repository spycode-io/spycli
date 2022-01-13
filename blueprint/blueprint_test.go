package blueprint_test

import (
	"testing"

	"github.com/spycode-io/spycli/blueprint"
	"github.com/spycode-io/spycli/model"
	"github.com/spycode-io/spycli/project"
)

func TestNewBlueprint(t *testing.T) {

	baseScaffold := model.NewScaffold("My Blueprint", ".iac-test", "templates/bp")

	_, err := blueprint.NewBlueprint(
		baseScaffold,
		"git@github.com:spycode-io/tf-modules.git",
		project.DefaultRegions)

	if nil != err {
		t.Error(err)
	}
}
