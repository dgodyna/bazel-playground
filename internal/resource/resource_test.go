package resource

import (
	"gotest.tools/assert"
	"io/ioutil"
	"path/filepath"
	"testing"
)

// some test which will perform reading of the file from resource directory
func TestReadFile(t *testing.T) {

	content, err := ioutil.ReadFile(filepath.Join("testdata", "test.yaml"))
	assert.NilError(t, err)
	assert.Equal(t, "data: test", string(content))

}
