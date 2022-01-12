package project_test

import (
	"testing"

	"github.com/spycode-io/spycli/model"
	"github.com/spycode-io/spycli/project"
)

func TestCreateAWSProject(t *testing.T) {

	baseScaffold := model.NewScaffold("My Project", ".iac-test", "templates/prj")

	_, err := project.NewProject(
		baseScaffold,
		"aws",
		"my-stack",
		"my-blueprint",
		false, "", "",
		project.DefaultEnvironments, project.DefaultRegions)
	if nil != err {
		t.Error(err)
	}
}
