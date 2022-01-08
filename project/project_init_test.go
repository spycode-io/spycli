package project_test

import (
	"fmt"
	"testing"

	"github.com/spycode-io/spycli/blueprint"
	"github.com/spycode-io/spycli/model"
	"github.com/spycode-io/spycli/module"
	"github.com/spycode-io/spycli/project"
)

const projectFolder = ".iac-test"
const stackName = "integration-test-local"

func TestLocalFlow(t *testing.T) {

	NewProjectStructure(t)

	err := project.InitProject(projectFolder)
	if err != nil {
		t.Error(err)
	}
}

func NewProjectStructure(t *testing.T) {

	//Create a new blueprint
	bp, err := blueprint.NewBlueprint(
		model.NewScaffold(
			"BP Integration Test Local",
			projectFolder, "templates/bp"),
		"git@github.com:spycode-io/bp-integration-test-local.git",
		"v0.0.0",
		stackName,
		project.DefaultRegions)

	if nil != err || nil == bp {
		t.Error(err)
	}

	//Create a new project
	prj, err := project.NewProject(
		model.NewScaffold("Prj Integration Test Local", projectFolder, "templates/prj"),
		"aws",
		stackName,
		bp.BluePrintPath,
		"v0.0.0",
		"tf-modules",
		"v0.0.0",
		project.DefaultEnvironments, project.DefaultRegions)

	if nil != err || nil == prj {
		t.Error(err)
	}

	//Create a new module
	module, err := module.NewModule(
		model.NewScaffold(
			"Module 01",
			fmt.Sprintf(".iac-test/%s/%s/any",
				bp.SlugName, bp.Stack),
			"templates/mdl"),
	)

	if nil != err || nil == module {
		t.Error(err)
	}
}
