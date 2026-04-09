package parsePlaylists

import (
	"fmt"
	"os"
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
	for index, line := range splittedLines {
		trackNameIndex := len(line) - 1
		trackName := line[trackNameIndex]
		line[trackNameIndex] = fmt.Sprintf("%d - %s", index, trackName)
	}
	return splittedLines, nil
}
func ParsePlaylist() {
	args := os.Args[1:]
	path := args[0]
	data, _ := ReadPlaylist(path)
	lines := strings.Split(data, "\n")
	splittedLines := make([][]string, len(lines))
	for index, line := range lines {
		splittedLines[index] = strings.Split(line, "/")
	}
	splittedLines, _ = DefineSongOrder(splittedLines)

}
