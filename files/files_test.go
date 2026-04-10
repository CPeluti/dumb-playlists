package files

import (
	"dumbplaylists/configuration"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestProcessFiles(t *testing.T) {
	//TODO: improve tests
	t.Run("Should create the target folder if it doesn't exists", func(t *testing.T) {
		cleanEnv(t)
		config, _ := configuration.ReadYaml("./fixtures/config.yaml")

		processFiles("teste", config)

		if !checkDirectory("./fixtures/device") {
			t.Errorf("Directory not created")
		} else {
			os.RemoveAll("./fixtures/device")
		}
	})

	t.Run("Copied all the target files with the right name", func(t *testing.T) {
		cleanEnv(t)
		config, _ := configuration.ReadYaml("./fixtures/config.yaml")
		processFiles("teste", config)
		expected := []string{
			"0 - a.flac",
			"1 - b.flac",
		}
		infos, _ := os.ReadDir("./fixtures/device")

		var names []string
		for _, info := range infos {
			names = append(names, info.Name())
		}
		fmt.Println(names)
		fmt.Println()
		if !reflect.DeepEqual(names, expected) {
			t.Errorf("Expected %q and got %q", expected, names)
		}
	})
}
func cleanEnv(t *testing.T) {
	t.Helper()
	t.Cleanup(func() {
		os.RemoveAll("./fixtures/device")
	})
}
