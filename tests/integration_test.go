package tests_test

import (
	"fmt"
	"testing"

	"github.com/spycode-io/spycli/blueprint"
	"github.com/spycode-io/spycli/model"
	"github.com/spycode-io/spycli/module"
	"github.com/spycode-io/spycli/project"
)

func TestLocalFlow(t *testing.T) {

	//Create a new blueprint
	bpScaffold := model.NewScaffold(
		"BP Test",
		".iac-test", "templates/bp")

	bp, err := blueprint.NewBlueprint(
		bpScaffold,
		"git@github.com:spycode-io/bp-test.git",
		"v0.0.0",
		"web-stack",
		[]string{})

	if nil != err || nil == bp {
		t.Error(err)
	}

	//Create a new project
	prj, err := project.NewProject(
		model.NewScaffold("Prj Test", ".iac-test", "templates/prj"),
		"aws",
		"web-stack",
		"git@github.com:spycode-io/bp-test.git",
		"v0.0.0",
		"../../../../tf-modules",
		"",
		project.DefaultEnvironments, project.DefaultRegions)

	if nil != err || nil == prj {
		t.Error(err)
	}

	//Create new modules
	anyRegionBasePath := fmt.Sprintf(".iac-test/%s/%s/_any", bp.SlugName, bp.Stack)

	vpc, err := createModule(anyRegionBasePath, "Test VPC", "aws/vpc")
	if nil != err || nil == vpc {
		t.Error(err)
	}

	vms, err := createModule(anyRegionBasePath, "Test VMS", "aws/vm")
	if nil != err || nil == vms {
		t.Error(err)
	}
}

func createModule(baseFolder string, name string, moduleName string) (*module.Module, error) {
	scaffold := model.NewScaffold(
		name,
		baseFolder,
		"templates/mdl")

	return module.NewModule(scaffold, moduleName)
}
