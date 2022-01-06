package model_test

import (
	"testing"

	"github.com/spycode-io/spycli/model"
)

func TestNewGenericScaffold(t *testing.T) {
	base := model.NewScaffold("My Scaffold", ".iac-test", "templates/generic")
	_, err := model.NewGenericScafold(base, "foo", "bar")

	if nil != err {
		t.Error(err)
	}
}
