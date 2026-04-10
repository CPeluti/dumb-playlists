package files

import (
	"dumbplaylists/configuration"
	"dumbplaylists/parsePlaylists"
	"errors"
	"fmt"
	"io"
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
func checkDirectory(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func prepareOutputs(playlistFolderPath string, device configuration.Device) ([]parsePlaylists.Output, error) {
	var playlists []parsePlaylists.Output

	for _, playlist := range device.Allowlist {
		filetype := "m3u"
		if splitted := strings.Split(playlist, "."); len(splitted) > 1 {
			filetype = splitted[len(splitted)-1]
		}
		playlistFilename := playlist + "." + filetype

		outputs, _ := parsePlaylists.ParsePlaylist(filepath.Join(playlistFolderPath, playlistFilename))

		playlists = append(playlists, outputs)
	}
	return playlists, nil
}
func getPlaylistsPath(config configuration.Config) (string, error) {
	playlistFolderPath := config.PlaylistsDir
	if filepath.IsLocal(config.PlaylistsDir) {
		//Uses relative path based on the config file path
		playlistFolderPath = filepath.Join(config.ConfigPath, "..", config.PlaylistsDir)
	}

	if !checkDirectory(playlistFolderPath) {
		return "", errors.New("PlaylistsDir is not a path")
	}
	return playlistFolderPath, nil
}
func processFiles(deviceName string, config configuration.Config) {
	//TODO: add check to see if target folder exists
	playlistFolderPath, err := getPlaylistsPath(config)
	if err != nil {
		panic(err)
	}
	var selectedDevice configuration.Device
	for _, device := range config.Devices {
		if device.Name == deviceName {
			//check if source files and destination are valid
			selectedDevice = device
		}
	}
	outputs, _ := prepareOutputs(playlistFolderPath, selectedDevice)

	for _, output := range outputs {
		for i := 0; i < len(output.Source); i++ {
			targetPath := filepath.Join(selectedDevice.Path, output.Target[i])
			copy(output.Source[i], targetPath)
		}
	}
}
func checkPlaylistsOnDevice() {
	//TODO: check if playlist already exists, if so only copies what's missing and what's extra
}
func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
