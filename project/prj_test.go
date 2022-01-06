package project_test

import (
	"testing"

	"github.com/spycode-io/spycli/assets"
	"github.com/spycode-io/spycli/project"
)

func TestCreateAWSProject(t *testing.T) {

	_, err := project.New(assets.TemplatesData, ".iac-test", "aws", "my-project", "my-stack", "blueprint", "v1", project.DefaultEnvironments, project.DefaultRegions)
	if nil != err {
		t.Error(err)
	}
}
