package parsePlaylists

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadPlaylist(path string) (string, error) {
	data, err := os.ReadFile(path)
	check(err)
	return string(data), nil
}
func DefineSongOrder(splittedLines [][]string) ([][]string, error) {
	digitCount := len(strconv.Itoa(len(splittedLines)))

	for index, line := range splittedLines {
		trackNameIndex := len(line) - 1
		trackName := line[trackNameIndex]

		line[trackNameIndex] = fmt.Sprintf(
			"%0*d - %s",
			digitCount,
			index,
			trackName)
	}
	return splittedLines, nil
}
func ParsePlaylist(path string) ([][]string, error) {
	data, _ := ReadPlaylist(path)
	if len(data) <= 0 {
		panic("playlist is empty")
	}
	lines := strings.Split(data, "\n")
	splittedLines := make([][]string, len(lines))
	for index, line := range lines {
		splittedLines[index] = strings.Split(line, "/")
	}
	splittedLines, _ = DefineSongOrder(splittedLines)
	return splittedLines, nil
}
