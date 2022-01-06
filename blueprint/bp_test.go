package blueprint_test

import (
	"testing"

	"github.com/spycode-io/spycli.git/assets"
	"github.com/spycode-io/spycli.git/blueprint"
)

func TestAbs(t *testing.T) {

	_, err := blueprint.NewBlueprint(assets.TemplatesData, "iac-test", "bp-test", "blueprint", "v1", "stack-x", []string{"eastus", "westus", "centralus"})

	if nil != err {
		t.Error(err)
	}
}
