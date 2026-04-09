package files

import (
	"dumbplaylists/configuration"
	"testing"
)

func TestFiles(t *testing.T) {
	t.Run("Find the list of files for each playlist", func(t *testing.T) {
		config, _ := configuration.ReadYaml("./fixtures/config.yaml")
		prepareFiles("teste", config)
	})

}
