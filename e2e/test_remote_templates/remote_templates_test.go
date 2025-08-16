package testremotetemplates

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var configTemplate = `
dir: %s
filename: %s
template: %s
formatter: noop
force-file-write: true
pkgname: test_pkgname
template-data:
  foo: foo
  bar: bar
packages:
  github.com/paivagustavo/mockery/v3/internal/fixtures/template_exercise:
    interfaces:
      Exercise:
`

func TestRemoteTemplates(t *testing.T) {
	// the temp dir needs to reside within the mockery project because mockery
	// requires a go.mod file to function correctly. Using t.TempDir() won't work
	// because of this.
	tmpDirBase := "./test"
	_ = os.RemoveAll(tmpDirBase)
	require.NoError(t, os.Mkdir(tmpDirBase, 0o755))

	//nolint:errcheck
	defer os.RemoveAll(tmpDirBase)

	type test struct {
		name             string
		schema           string
		expectMockeryErr bool
	}
	for _, tt := range []test{
		{
			name: "schema validation OK",
			schema: `{
	"$schema": "http://json-schema.org/draft-07/schema#",
	"title": "vektra/mockery matryer mock",
	"type": "object",
	"additionalProperties": false,
	"properties": {
		"foo": {
		"type": "string"
		},
		"bar": {
		"type": "string"
		}
	},
	"required": ["foo", "bar"]
}`,
			expectMockeryErr: false,
		},
		{
			name: "Required parameter doesn't exist",
			schema: `{
	"$schema": "http://json-schema.org/draft-07/schema#",
	"title": "vektra/mockery matryer mock",
	"type": "object",
	"additionalProperties": false,
	"properties": {
		"foo": {
			"type": "string"
		},
		"bar": {
			"type": "string"
		},
		"baz": {
			"type": "string"
		}
	},
	"required": ["foo", "bar", "baz"]
}`,
			expectMockeryErr: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			tmpdir := path.Join(tmpDirBase, t.Name())
			require.NoError(t, os.MkdirAll(tmpdir, 0o755))

			configFile := path.Join(tmpdir, ".mockery.yml")
			outFile := path.Join(tmpdir, "out.txt")

			templateName := "template.templ"
			mux := http.NewServeMux()
			mux.HandleFunc(fmt.Sprintf("/%s", templateName), func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "Hello, world!")
			})
			mux.HandleFunc(fmt.Sprintf("/%s.schema.json", templateName), func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, tt.schema)
			})

			ts := httptest.NewServer(mux)
			defer ts.Close()

			fullPath := fmt.Sprintf("%s/%s", ts.URL, templateName)

			parent, name := path.Split(outFile)
			configFileContents := fmt.Sprintf(
				configTemplate,
				parent,
				name,
				fullPath,
			)
			require.NoError(t, os.WriteFile(configFile, []byte(configFileContents), 0o600))

			//nolint: gosec
			out, err := exec.Command(
				"go", "run", "github.com/paivagustavo/mockery/v3",
				"--config", configFile).CombinedOutput()
			if tt.expectMockeryErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err, string(out))
				outFileBytes, err := os.ReadFile(outFile)
				require.NoError(t, err)
				assert.Equal(t, "Hello, world!", string(outFileBytes))
			}
		})
	}
}
