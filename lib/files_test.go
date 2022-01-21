package lib_test

import (
	"os"
	"testing"

	"github.com/spycode-io/spycli/lib"
)

func TestCopyFile(t *testing.T) {

	os.RemoveAll(".iac-test/copy-test")

	src := ".iac-test/copy-test/file"
	dst := ".iac-test/copy-test/dst/dst_file"

	f, err := lib.CreateFile(src)
	if nil != err {
		t.Error(err)
	}
	defer f.Close()

	if !lib.FileExists(src) {
		t.Fail()
	}

	err = lib.CopyFile(src, dst)
	if nil != err {
		t.Error(err)
	}

	if !lib.FileExists(dst) {
		t.Fail()
	}
}
