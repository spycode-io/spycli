package blueprint_test

import (
	"testing"

	"github.com/spycode-io/spycli/assets"
	"github.com/spycode-io/spycli/blueprint"
)

func TestAbs(t *testing.T) {

	_, err := blueprint.NewBlueprint(assets.TemplatesData, ".iac-test", "bp-test", "blueprint", "v1", "stack-x", []string{"us-east-1", "us-east-2", "us-west-1"})

	if nil != err {
		t.Error(err)
	}
}
