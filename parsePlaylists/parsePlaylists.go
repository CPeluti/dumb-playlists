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

type Output struct {
	Source []string
	Target []string
}

func ReadPlaylist(path string) (string, error) {
	data, err := os.ReadFile(path)
	check(err)
	return string(data), nil
}
func DefineTrackOrder(lines []string) ([]string, error) {
	digitCount := len(strconv.Itoa(len(lines)))
	var list []string
	for index, line := range lines {
		splittedLine := strings.Split(line, "/")
		trackNameIndex := len(splittedLine) - 1
		trackName := splittedLine[trackNameIndex]

		newTrackname := fmt.Sprintf(
			"%0*d - %s",
			digitCount,
			index,
			trackName)
		list = append(list, newTrackname)

	}
	return list, nil
}
func ParsePlaylist(path string) (Output, error) {
	data, _ := ReadPlaylist(path)
	if len(data) <= 0 {
		panic("playlist is empty")
	}
	lines := strings.Split(data, "\n")
	list, _ := DefineTrackOrder(lines)
	out := Output{
		Source: lines,
		Target: list,
	}
	return out, nil
}
