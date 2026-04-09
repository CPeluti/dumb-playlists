package files

import (
	"dumbplaylists/configuration"
	"dumbplaylists/parsePlaylists"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	check(err)
	return string(data), nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func prepareFiles(deviceName string, config configuration.Config) {
	for _, device := range config.Devices {
		if device.Name == deviceName {
			playlistFolderPath := config.PlaylistsDir
			if filepath.IsLocal(config.PlaylistsDir) {
				//Uses relative path based on the config file path
				playlistFolderPath = filepath.Join(config.ConfigPath, "..", config.PlaylistsDir)
			}
			info, err := os.Stat(playlistFolderPath)
			if err != nil || !info.IsDir() {
				fmt.Println("Error in the playlists path")
			}
			for _, playlist := range device.Allowlist {
				filetype := "m3u"
				if splitted := strings.Split(playlist, "."); len(splitted) > 1 {
					filetype = splitted[len(splitted)-1]
				}
				playlistFilename := playlist + "." + filetype

				playlists, _ := parsePlaylists.ParsePlaylist(filepath.Join(playlistFolderPath, playlistFilename))
				fmt.Println(playlists)
			}
		}
	}
}
