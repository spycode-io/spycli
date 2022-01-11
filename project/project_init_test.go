package project_test

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

func TestLinkingFiles(t *testing.T) {
	os.RemoveAll(".iac-test/source")
	os.RemoveAll(".iac/destiny")

	os.MkdirAll(".iac-test/source/m1", os.ModePerm)
	os.MkdirAll(".iac-test/source/m2", os.ModePerm)

	f1 := []byte("f1")
	os.WriteFile(".iac-test/source/m1/f1", f1, 0644)

	f2 := []byte("f2")
	os.WriteFile(".iac-test/source/m2/f2", f2, 0644)

	os.MkdirAll(".iac-test/destiny", os.ModePerm)

	project.LinkBlueprintFolders(".iac-test/source", ".iac-test/destiny", []string{}, true)

	if _, err := os.Stat(".iac-test/destiny/m1/f1"); errors.Is(err, os.ErrNotExist) {
		t.Fail()
	}

	if _, err := os.Stat(".iac-test/destiny/m1/f1"); errors.Is(err, os.ErrNotExist) {
		t.Fail()
	}
}

func TestCopyingFiles(t *testing.T) {
	os.RemoveAll(".iac-test/source")
	os.RemoveAll(".iac/destiny")

	os.MkdirAll(".iac-test/source/m1", os.ModePerm)
	os.MkdirAll(".iac-test/source/m2", os.ModePerm)

	f1 := []byte("f1")
	os.WriteFile(".iac-test/source/m1/f1", f1, 0644)

	f2 := []byte("f2")
	os.WriteFile(".iac-test/source/m2/f2", f2, 0644)

	os.MkdirAll(".iac-test/destiny", os.ModePerm)

	project.CopyBlueprintFolders(".iac-test/source", ".iac-test/destiny", []string{}, true)

	if _, err := os.Stat(".iac-test/destiny/m1/f1"); errors.Is(err, os.ErrNotExist) {
		t.Fail()
	}

	if _, err := os.Stat(".iac-test/destiny/m1/f1"); errors.Is(err, os.ErrNotExist) {
		t.Fail()
	}
}

func TestLocalFlow(t *testing.T) {

	prj, err := NewProjectStructure(t)
	if err != nil {
		t.Error(err)
	}

	_, err = NewBlueprint(t)
	if err != nil {
		t.Error(err)
	}

	err = project.InitProject(prj.ProjectPath, false)
	if err != nil {
		t.Error(err)
	}
}

func NewProjectStructure(t *testing.T) (*project.ProjectScaffold, error) {

	//Create a new project
	os.RemoveAll(".iac-test/my-project")
	return project.NewProject(
		model.NewScaffold("My Project", ".iac-test", "templates/prj"),
		"aws",
		"web-stack",
		".iac-test/bp-test",
		"v0.0.0",
		"../../../../tf-aws-modules",
		"",
		"",
		project.DefaultEnvironments, project.DefaultRegions)
}

func NewBlueprint(t *testing.T) (bp *blueprint.BlueprintScaffold, err error) {
	//Create a new blueprint
	os.RemoveAll(".iac-test/bp-test")
	bpScaffold := model.NewScaffold(
		"BP Test",
		".iac-test", "templates/bp")

	bp, err = blueprint.NewBlueprint(
		bpScaffold,
		"git@github.com:spycode-io/bp-test.git",
		"v0.0.0",
		"web-stack",
		[]string{})

	if nil != err {
		return
	}

	//Create new modules
	anyRegionBasePath := fmt.Sprintf(".iac-test/%s/%s/_any", bp.SlugName, bp.Stack)

	vpc, err := CreateModule(anyRegionBasePath, "VPC Test", "vpc")
	if nil != err || nil == vpc {
		t.Error(err)
	}

	vms, err := CreateModule(anyRegionBasePath, "VMS Test", "vm")
	if nil != err || nil == vms {
		t.Error(err)
	}

	return
}

func CreateModule(baseFolder string, name string, moduleName string) (*module.Module, error) {
	scaffold := model.NewScaffold(
		name,
		baseFolder,
		"templates/mdl")

	return module.NewModule(scaffold, moduleName)
}
