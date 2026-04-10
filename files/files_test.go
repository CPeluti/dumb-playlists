package files

import (
	"dumbplaylists/configuration"
	"testing"
)

func TestFiles(t *testing.T) {
	//TODO: improve tests
	t.Run("write files to directory", func(t *testing.T) {
		config, _ := configuration.ReadYaml("./fixtures/config.yaml")
		processFiles("teste", config)
	})

}
