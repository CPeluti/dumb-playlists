package configuration

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestConfiguration(t *testing.T) {
	t.Run("read the yaml file", func(t *testing.T) {
		path := "./fixtures/config.yaml"
		absPath, _ := filepath.Abs(path)
		got, err := ReadYaml(path)
		want := Config{
			ConfigPath:   absPath,
			PlaylistsDir: "/mnt/playlists",
			Devices: []Device{
				{
					Name:      "teste",
					Path:      "/mnt/teste",
					Allowlist: []string{"playlist1", "playlist2"},
				},
				{
					Name:      "teste2",
					Path:      "/mnt/teste2",
					Allowlist: []string{"playlist3", "playlist4"},
				},
			},
		}
		if err != nil {
			t.Errorf("Got error %q", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %q but got %q", want, got)
		}
	})
	t.Run("read the yaml file", func(t *testing.T) {
		got, err := ReadYaml("./fixtures/invalid_config.yaml")
		want := "Key: 'Config.PlaylistsDir' Error:Field validation for 'PlaylistsDir' failed on the 'required' tag"
		if err == nil {
			t.Errorf("Expected error but got none.")
		}
		if err.Error() != want {
			t.Errorf("expected %q but got %q", want, got)
		}
	})
}
