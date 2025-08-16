package test_template_exercise

import (
	"fmt"
	"os"
	"os/exec"
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
	if err != nil {
		fmt.Println(err)
		fmt.Println(string(out))
		os.Exit(1)
	}

	b, err := os.ReadFile(outfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	expectedPath := "exercise_expected.txt"
	expected, err := os.ReadFile(expectedPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	assert.Equal(t, string(expected), string(b))
}
