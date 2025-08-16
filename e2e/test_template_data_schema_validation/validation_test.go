package test_template_exercise

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExercise(t *testing.T) {
	t.Parallel()
	outfile := "./exercise.txt"
	//nolint:errcheck
	defer os.Remove(outfile)

	out, err := exec.Command(
		"go", "run", "github.com/paivagustavo/mockery/v3",
		"--config", "./.mockery.yml").CombinedOutput()
	assert.Error(t, err)
	expectedString := "ERR (root): foo is required"
	assert.True(t, strings.Contains(string(out), expectedString), "expected string in stdout not found: \"%s\"", expectedString)
}
