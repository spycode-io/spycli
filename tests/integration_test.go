package tests_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/spycode-io/spycli/blueprint"
	"github.com/spycode-io/spycli/model"
	"github.com/spycode-io/spycli/module"
	"github.com/spycode-io/spycli/project"
)

func TestIntegrationLocalFlow(t *testing.T) {

	//Cleam test folder
	os.RemoveAll(".iac-test")

	//Create a new blueprint
	bpScaffold := model.NewScaffold(
		"Bp Test",
		".iac-test", "templates/bp")

	bp, err := blueprint.NewBlueprint(
		bpScaffold,
		"bp-test",
		[]string{})

	if nil != err || nil == bp {
		t.Error(err)
	}

	//Create a new project
	prj, err := project.NewProject(
		model.NewScaffold(
			"My Project",
			".iac-test", "templates/prj"),
		"aws",
		"web-stack",
		"bp-test",
		false,
		"",
		"",
		project.DefaultEnvironments, project.DefaultRegions)

	if nil != err || nil == prj {
		t.Error(err)
	}

	//Create new modules
	anyRegionBasePath := fmt.Sprintf(".iac-test/%s/%s/_any", bp.SlugName, bp.Stack)

	vpc, err := createModule(anyRegionBasePath, "My VPC", "vpc")
	if nil != err || nil == vpc {
		t.Error(err)
	}

	vms, err := createModule(anyRegionBasePath, "My VMS", "vm")
	if nil != err || nil == vms {
		t.Error(err)
	}

	project.InitProject(prj.ProjectPath, true)

	if _, err := os.Stat(fmt.Sprintf("%s/dev/us-east-1/my-vpc", prj.ProjectPath)); errors.Is(err, os.ErrNotExist) {
		t.FailNow()
	}

	if _, err := os.Stat(fmt.Sprintf("%s/dev/us-east-1/my-vms", prj.ProjectPath)); errors.Is(err, os.ErrNotExist) {
		t.FailNow()
	}
}

func TestRemoteFlow(t *testing.T) {

	//Cleam test folder
	os.RemoveAll(".iac-test")

	//Create a new blueprint
	bpScaffold := model.NewScaffold(
		"Bp Test",
		".iac-test", "templates/bp")

	bp, err := blueprint.NewBlueprint(
		bpScaffold,
		"bp-test",
		[]string{})

	if nil != err || nil == bp {
		t.Error(err)
	}

	//Create a new project
	prj, err := project.NewProject(
		model.NewScaffold("My Project", ".iac-test", "templates/prj"),
		"aws",
		"web-stack",
		"bp-test",
		false, "", "",
		project.DefaultEnvironments, project.DefaultRegions)

	if nil != err || nil == prj {
		t.Error(err)
	}

	//Create new modules
	anyRegionBasePath := fmt.Sprintf(".iac-test/%s/%s/_any", bp.SlugName, bp.Stack)

	vpc, err := createModule(anyRegionBasePath, "My VPC", "vpc")
	if nil != err || nil == vpc {
		t.Error(err)
	}

	vms, err := createModule(anyRegionBasePath, "My VMS", "vm")
	if nil != err || nil == vms {
		t.Error(err)
	}

	err = project.InitProject(prj.ProjectPath, false)
	if nil != err {
		t.Error(err)
	}
}

func createModule(baseFolder string, name string, moduleUrl string) (*module.Module, error) {
	scaffold := model.NewScaffold(
		name,
		baseFolder,
		"templates/mdl")

	return module.NewModule(scaffold, moduleUrl, true)
}
