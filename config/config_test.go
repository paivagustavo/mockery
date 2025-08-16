package config

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRootConfig(t *testing.T) {
	tests := []struct {
		name    string
		config  string
		wantErr error
	}{
		{
			name: "unrecognized parameter",
			config: `
packages:
  github.com/foo/bar:
    config:
      unknown: param
`,
			wantErr: fmt.Errorf("'packages[github.com/foo/bar].config' has invalid keys: unknown"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configFile := path.Join(t.TempDir(), "config.yaml")
			require.NoError(t, os.WriteFile(configFile, []byte(tt.config), 0o600))

			flags := pflag.NewFlagSet("test", pflag.ExitOnError)
			flags.String("config", "", "")

			require.NoError(t, flags.Parse([]string{"--config", configFile}))

			_, _, err := NewRootConfig(context.Background(), flags)
			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else {
				var original error
				cursor := err
				for cursor != nil {
					original = cursor
					cursor = errors.Unwrap(cursor)
				}
				assert.Equal(t, tt.wantErr.Error(), original.Error())
			}
		})
	}
}

func TestNewRootConfigUnknownEnvVar(t *testing.T) {
	t.Setenv("MOCKERY_UNKNOWN", "foo")
	configFile := path.Join(t.TempDir(), "config.yaml")
	require.NoError(t, os.WriteFile(configFile, []byte(`
packages:
  github.com/paivagustavo/mockery/v3:
`), 0o600))

	flags := pflag.NewFlagSet("test", pflag.ExitOnError)
	flags.String("config", "", "")

	require.NoError(t, flags.Parse([]string{"--config", configFile}))
	_, _, err := NewRootConfig(context.Background(), flags)
	assert.NoError(t, err)
}
