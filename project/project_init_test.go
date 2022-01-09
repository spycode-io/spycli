package project_test

import (
	"testing"

	"github.com/spycode-io/spycli/model"
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

	//Create a new project
	prj, err := project.NewProject(
		model.NewScaffold("Prj Integration Test Local", projectFolder, "templates/prj"),
		"aws",
		stackName,
		"bp-integration-test",
		"v0.0.0",
		"tf-modules",
		"v0.0.0",
		project.DefaultEnvironments, project.DefaultRegions)

	if nil != err || nil == prj {
		t.Error(err)
	}
}
